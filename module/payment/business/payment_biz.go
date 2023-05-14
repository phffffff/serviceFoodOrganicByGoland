package paymentBusiness

import (
	"context"
	"go_service_food_organic/common"
)

type PaymentRepositoy interface {
	PaymentRepo(c context.Context) error
}

type paymentBiz struct {
	repo PaymentRepositoy
}

func NewPaymentBiz(repo PaymentRepositoy) *paymentBiz {
	return &paymentBiz{repo: repo}
}

func (biz *paymentBiz) Payment(c context.Context) error {
	if err := biz.repo.PaymentRepo(c); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
