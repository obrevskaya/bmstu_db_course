package controllers

import (
	"context"
	"fmt"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/google/uuid"
)

const Percent = 0.1

type BillController struct {
	billRep  interfaces.IBillRepository
	userRep  interfaces.IUserRepository
	orderRep interfaces.IOrderRepository
}

func NewBillController(b interfaces.IBillRepository, u interfaces.IUserRepository,
	o interfaces.IOrderRepository) *BillController {
	return &BillController{
		billRep:  b,
		userRep:  u,
		orderRep: o,
	}
}

func (b *BillController) Create(ctx context.Context, bill *models.Bill) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to create bill", "billID", bill.ID, "err", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start create bill")
	order, err := b.orderRep.Get(ctx, bill.IDOrder)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get order to create bill", "billID", bill.ID, "err", err)
		return fmt.Errorf("get order: %w", err)
	}

	if user.Status < models.Admin && user.ID != order.IDUser {
		myContext.LoggerFromContext(ctx).Warnw("doesn't access", "billID", bill.ID)
		return fmt.Errorf("access rights: %w", myErrors.ErrAccess)
	}

	err = b.billRep.Insert(ctx, bill)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot create bill", "billID", bill.ID, "err", err)
		return fmt.Errorf("create bill: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully create bill", "billID", bill.ID)
	return nil
}

func (b *BillController) Get(ctx context.Context, ID uuid.UUID) (*models.Bill, error) {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to get bill", "billID", ID, "err", err)
		return nil, fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start get bill", "billID", ID)
	bill, err := b.billRep.Get(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get bill", "billID", ID, "err", err)
		return nil, fmt.Errorf("get bill: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get bill", "billID", ID)
	return bill, nil
}

func (b *BillController) UpdateBillStatus(ctx context.Context, ID uuid.UUID, status string) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to update bill status", "billID", ID, "err", err)
		return fmt.Errorf("get user: %w", err)
	}
	if user.Status < models.Admin {
		myContext.LoggerFromContext(ctx).Warnw("don't access to update bill status", "billID", ID)
		return fmt.Errorf("access rights: %w", myErrors.ErrAccess)
	}

	myContext.LoggerFromContext(ctx).Infow("start update bill status", "billID", ID)

	bill, err := b.billRep.Get(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get bill to update bill status", "billID", ID, "err", err)
		return fmt.Errorf("get bill: %w", err)
	}
	if bill.Status == models.PaidBill {
		myContext.LoggerFromContext(ctx).Warnw("bill also paid", "billID")
		return fmt.Errorf("bill also paid")
	}

	order, err := b.orderRep.Get(ctx, bill.IDOrder)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get order to update bill status", "billID", ID, "err", err)
		return fmt.Errorf("get order: %w", err)
	}

	user, err = b.userRep.Get(ctx, order.IDUser)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to update bill status", "billID", ID, "err", err)
		return fmt.Errorf("get user: %w", err)
	}

	points := user.Points
	if order.IsPoints == false {
		points += int(float64(bill.Price) * Percent)
	}

	err = b.billRep.UpdateBillStatus(ctx, ID, status, points)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot update bill status", "billID", ID, "err", err)
		return fmt.Errorf("update bill: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully update bill status", "billID", ID)

	return nil
}
