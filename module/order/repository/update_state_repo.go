package orderRepo

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

type UpdateOrderStateStore interface {
	FindDataWithCondition(c context.Context, cond map[string]interface{}, moreKeys ...string) (*orderModel.Order, error)
	UpdateState(c context.Context, id int, state string) error
	BeginTransaction() error
	RollbackTransaction() error
	CommitTransaction() error
}

type UpdateFoodRepo interface {
	UpdateCountFoodRepo(c context.Context, count, id int, typeOf string) error
}

type updateOrderStateRepo struct {
	store    UpdateOrderStateStore
	foodRepo UpdateFoodRepo
	req      common.Requester
}

func NewUpdateOrderStateRepo(store UpdateOrderStateStore, foodRepo UpdateFoodRepo, req common.Requester) *updateOrderStateRepo {
	return &updateOrderStateRepo{store: store, foodRepo: foodRepo, req: req}
}

func (repo *updateOrderStateRepo) UpdateOrderStateRepo(c context.Context, id int, state string) error {
	if repo.req.GetRole() != common.Admin {
		return common.ErrorNoPermission(nil)
	}

	order, err := repo.store.FindDataWithCondition(c, map[string]interface{}{"id": id}, "OrderDetails")
	if err != nil {
		return common.ErrRecordNotFound(orderModel.EntityName, err)
	}
	if order == nil {
		return common.ErrEntityNotExists(orderModel.EntityName, err)
	}

	if err := repo.store.BeginTransaction(); err != nil {
		return common.ErrInternal(err)
	}

	if state == orderModel.StateCancel && ((*order).Status != 0 && (*order).State != orderModel.StateCancel) {
		for _, item := range order.OrderDetails {
			if err := repo.foodRepo.UpdateCountFoodRepo(c, item.Quantity, item.FoodId, common.Increase); err != nil {
				if err := repo.store.RollbackTransaction(); err != nil {
					return common.ErrInternal(err)
				}
				return common.ErrInternal(err)
			}
		}

	}
	if state != orderModel.StateCancel && ((*order).Status == 0 && (*order).State == orderModel.StateCancel) {
		//phục hồi đơn
		for _, item := range order.OrderDetails {
			if err := repo.foodRepo.UpdateCountFoodRepo(c, item.Quantity, item.FoodId, common.Decrease); err != nil {
				if err := repo.store.RollbackTransaction(); err != nil {
					return common.ErrInternal(err)
				}
				return common.ErrInternal(err)
			}
		}
	}

	if err := repo.store.UpdateState(c, order.Id, state); err != nil {
		if err := repo.store.RollbackTransaction(); err != nil {
			return common.ErrInternal(err)
		}
		return common.ErrCannotCRUDEntity(orderModel.EntityName, common.Update, err)
	}

	if err := repo.store.CommitTransaction(); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
