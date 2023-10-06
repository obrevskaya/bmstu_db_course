package controllers

import (
	"context"
	"errors"
	"fmt"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/google/uuid"
)

type OrderController struct {
	orderRep interfaces.IOrderRepository
	elemRep  interfaces.IOrderElementRepository
	userRep  interfaces.IUserRepository
	billRep  interfaces.IBillRepository
	wineRep  interfaces.IWineRepository
}

func NewOrderController(b interfaces.IBillRepository, el interfaces.IOrderElementRepository,
	u interfaces.IUserRepository, o interfaces.IOrderRepository, w interfaces.IWineRepository) *OrderController {
	return &OrderController{
		billRep:  b,
		elemRep:  el,
		userRep:  u,
		orderRep: o,
		wineRep:  w,
	}
}

func (o *OrderController) Create(ctx context.Context, order *models.Order) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to create order", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start create order")
	_, err = o.orderRep.GetByUserInProcess(ctx, user.ID)
	if errors.Is(err, myErrors.ErrNotFound) {
		err = o.orderRep.Insert(ctx, order)
		if err != nil {
			myContext.LoggerFromContext(ctx).Errorw("cannot create order", "error", err)
			return fmt.Errorf("create order: %w", err)
		}
	} else {
		myContext.LoggerFromContext(ctx).Warnw("error in get order by user")
		return fmt.Errorf("error in get new order by user")
	}
	myContext.LoggerFromContext(ctx).Infow("successfully create order")

	return nil
}

func (o *OrderController) GetByID(ctx context.Context, ID uuid.UUID) (*models.Order, error) {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to get by id order", "error", err)
		return nil, fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start get by id order")
	order, err := o.orderRep.Get(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get by id order", "error", err)
		return nil, fmt.Errorf("get order: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get by id order")

	return order, err
}

func (o *OrderController) GetByUserInProcess(ctx context.Context, ID uuid.UUID) (*models.Order, error) {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to get in process order", "error", err)
		return nil, fmt.Errorf("get user: %w", err)
	}

	myContext.LoggerFromContext(ctx).Infow("start get in process order")

	order, err := o.orderRep.GetByUserInProcess(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get in process order", "error", err)
		return nil, fmt.Errorf("get order by user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get in process order")
	return order, err
}

func (o *OrderController) Delete(ctx context.Context, ID uuid.UUID) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to delete order", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	if user.Status < models.Admin {
		myContext.LoggerFromContext(ctx).Warnw("don't access rights to get in process order")
		return fmt.Errorf("access rights: %w", myErrors.ErrAccess)
	}
	myContext.LoggerFromContext(ctx).Infow("start delete order")

	err = o.orderRep.Delete(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot delete order", "error", err)
		return fmt.Errorf("delete order: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully delete order")
	return nil
}

func (o *OrderController) Update(ctx context.Context, order *models.Order) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to update order", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	if user.Status != models.Admin && user.ID != order.IDUser {
		myContext.LoggerFromContext(ctx).Warnw("don't access rights")
		return fmt.Errorf("access rights: %w", myErrors.ErrAccess)
	}
	myContext.LoggerFromContext(ctx).Infow("start update order")

	orderOld, err := o.orderRep.Get(ctx, order.ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get order to update order", "error", err)
		return fmt.Errorf("get old order: %w", err)
	}

	if orderOld.Status == models.PlacedOrder {
		myContext.LoggerFromContext(ctx).Warnw("order already placed")
		return fmt.Errorf("order already placed")
	}

	user, err = o.userRep.Get(ctx, order.IDUser)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to update order", "error", err)
		return fmt.Errorf("error user get: %w", err)
	}

	if order.IsPoints && user.Points < order.TotalPrice {
		myContext.LoggerFromContext(ctx).Warnw("error balance to update order")
		return fmt.Errorf("error balance: %w", myErrors.ErrPoints)
	}

	err = o.orderRep.Update(ctx, order)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot update order", "error", err)
		return fmt.Errorf("update order: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully update order")
	return nil
}
