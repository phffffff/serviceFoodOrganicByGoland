package paymentRepo

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/carts/model"
	foodModel "go_service_food_organic/module/food/model"
	orderModel "go_service_food_organic/module/order/model"
	orderRepo "go_service_food_organic/module/order/repository"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
	orderDetailRepo "go_service_food_organic/module/order_detail/repository"
	userModel "go_service_food_organic/module/user/model"
)

type ListCartStore interface {
	ListDataWithCondition(c context.Context, filter *cartModel.Filter, paging *common.Paging) ([]cartModel.Cart, error)
}

type FindFoodStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}) (*foodModel.Food, error)
}

type UpdateOrderPriceRepo interface {
	UpdateOrderPriceRepo(c context.Context, id int, data *orderModel.OrderUpdate) error
}

type UpdateFoodCountRepo interface {
	UpdateCountFoodRepo(c context.Context, count, id int) error
}

type DeleteCartWhenPayment interface {
	DeleteCartRepo(c context.Context) error
}
type paymentRepo struct {
	storeCart             ListCartStore
	storeFood             FindFoodStore
	storeOrder            orderRepo.CreateOrderStore
	storeOrderDetail      orderDetailRepo.CreateOrderDetailStore
	updatePriceRepo       UpdateOrderPriceRepo
	updateFoodCountRepo   UpdateFoodCountRepo
	deleteCartWhenPayment DeleteCartWhenPayment
	req                   common.Requester
}

func NewPaymentRepo(
	storeCart ListCartStore,
	storeFood FindFoodStore,
	storeOrder orderRepo.CreateOrderStore,
	storeOrderDetail orderDetailRepo.CreateOrderDetailStore,
	updatePriceRepo UpdateOrderPriceRepo,
	updateFoodCountRepo UpdateFoodCountRepo,
	deleteCartWhenPayment DeleteCartWhenPayment,
	req common.Requester) *paymentRepo {
	return &paymentRepo{
		storeCart:             storeCart,
		storeFood:             storeFood,
		storeOrder:            storeOrder,
		storeOrderDetail:      storeOrderDetail,
		updatePriceRepo:       updatePriceRepo,
		updateFoodCountRepo:   updateFoodCountRepo,
		deleteCartWhenPayment: deleteCartWhenPayment,
		req:                   req,
	}
}

func (repo *paymentRepo) PaymentRepo(c context.Context) error {

	if repo.req.GetUserId() == 0 {
		return userModel.ErrorUserNoLogin(nil)
	}
	userId := repo.req.GetUserId()

	filter := cartModel.Filter{UserId: userId}
	paging := common.Paging{}
	paging.FullFill()

	carts, err := repo.storeCart.ListDataWithCondition(c, &filter, &paging)

	if err != nil {
		return common.ErrEntityNotExists(cartModel.EntityName, err)
	}
	if carts == nil {
		return common.ErrRecordNotFound(cartModel.EntityName, err)
	}

	if err := repo.storeOrder.BeginTransaction(); err != nil {
		return common.ErrInternal(err)
	}
	if err := repo.storeOrderDetail.BeginTransaction(); err != nil {
		return common.ErrInternal(err)
	}

	var order orderModel.OrderCreate
	order.UserId = userId
	order.TotalPrice = 0
	if err := repo.storeOrder.Create(c, &order); err != nil {
		if err := repo.storeOrder.RollbackTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		if err := repo.storeOrderDetail.CommitTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		return common.ErrInternal(err)
	}

	var totalPrice float32

	for _, item := range carts {
		food, err := repo.storeFood.FindDataWithCondition(c, map[string]interface{}{"id": item.FoodId})
		if err != nil {
			if err := repo.storeOrder.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			if err := repo.storeOrderDetail.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			return common.ErrRecordNotFound(foodModel.EntityName, err)
		}
		if food == nil {
			if err := repo.storeOrder.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			if err := repo.storeOrderDetail.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			return common.ErrEntityNotExists(foodModel.EntityName, err)
		}
		if food.Status == 0 {
			if err := repo.storeOrder.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			if err := repo.storeOrderDetail.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			return common.ErrEntityDeleted(foodModel.EntityName, nil)
		}

		if food.Count < item.Quantity {
			if err := repo.storeOrder.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			if err := repo.storeOrderDetail.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			return orderDetailModel.ErrorQuantityInvalid(nil)
		}

		var orderDetail orderDetailModel.OrderDetailCreate

		orderDetail.OrderId = order.Id
		orderDetail.FoodId = item.FoodId
		orderDetail.Quantity = item.Quantity
		orderDetail.Price = item.Price

		if err := repo.storeOrderDetail.Create(c, &orderDetail); err != nil {
			if err := repo.storeOrderDetail.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			if err := repo.storeOrder.RollbackTransaction(); err != nil {
				return common.ErrInternal(err)
			}
			return common.ErrCannotCRUDEntity(orderDetailModel.EntityName, common.Create, err)
		}

		if err := repo.updateFoodCountRepo.UpdateCountFoodRepo(c, item.Quantity, item.FoodId); err != nil {
			return common.ErrInternal(err)
		}

		totalPrice += orderDetail.Price
	}

	var orderUpdate orderModel.OrderUpdate

	orderUpdate.TotalPrice = totalPrice

	if err := repo.storeOrder.CommitTransaction(); err != nil {
		return common.ErrInternal(err)
	}
	if err := repo.storeOrderDetail.CommitTransaction(); err != nil {
		return common.ErrInternal(err)
	}
	if err := repo.updatePriceRepo.UpdateOrderPriceRepo(c, order.Id, &orderUpdate); err != nil {
		return common.ErrInternal(err)
	}
	if err := repo.deleteCartWhenPayment.DeleteCartRepo(c); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
