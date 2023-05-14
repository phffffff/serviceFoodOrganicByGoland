package cartBusiness

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/cart/model"
)

type ListCartStore interface {
	ListDataWithCondition(c context.Context, filter *cartModel.Filter, paging *common.Paging) ([]cartModel.CartLst, error)
}

type listCartBiz struct {
	store ListCartStore
}

func NewListCartBiz(store ListCartStore) *listCartBiz {
	return &listCartBiz{store: store}
}

func (biz *listCartBiz) ListCart(c context.Context, filter *cartModel.Filter, paging *common.Paging) ([]cartModel.CartLst, error) {
	list, err := biz.store.ListDataWithCondition(c, filter, paging)
	if err != nil || list == nil {
		return nil, common.ErrInternal(err)
	}
	return list, nil
}
