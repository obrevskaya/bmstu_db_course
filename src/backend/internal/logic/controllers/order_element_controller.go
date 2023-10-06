package controllers

import (
	"context"
	"fmt"
	"time"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	interfaces2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/google/uuid"
)

type OrderElemController struct {
	orderElemRep interfaces2.IOrderElementRepository
	orderRep     interfaces2.IOrderRepository
	wineRep      interfaces2.IWineRepository

	orderLogic interfaces2.IOrderController
}

func NewElemController(el interfaces2.IOrderElementRepository, o interfaces2.IOrderRepository,
	w interfaces2.IWineRepository, oC interfaces2.IOrderController) *OrderElemController {
	return &OrderElemController{
		orderElemRep: el,
		orderRep:     o,
		wineRep:      w,
		orderLogic:   oC,
	}
}

func (el *OrderElemController) Create(ctx context.Context, elem *models.OrderElement) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to create elem", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start create elem", "elemID", elem.ID)
	wine, err := el.wineRep.GetWine(ctx, elem.IDWine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get wine to create elem", "elemID", elem.ID, "error", err)
		return fmt.Errorf("get wine: %w", err)
	}
	if wine.Count < elem.Count {
		myContext.LoggerFromContext(ctx).Warnw("doesn't cnt wine to create elem", "elemID", elem.ID)
		return fmt.Errorf("cnt wine: %w", errors.ErrCntWine)
	}

	order, err := el.orderLogic.GetByUserInProcess(ctx, user.ID)
	if order == nil {
		order = &models.Order{
			IDUser:     user.ID,
			TotalPrice: 0,
			IsPoints:   false,
			Status:     models.ProcessOrder,
		}
		err = el.orderRep.Insert(ctx, order)
		if err != nil {
			myContext.LoggerFromContext(ctx).Warnw("cannot insert order to create elem", "elemID", elem.ID, "error", err)
			return fmt.Errorf("insert order: %w", err)
		}
	}

	elem.IDOrder = order.ID
	if el.orderElemRep.WineInOrder(ctx, wine.ID, order.ID) == nil {
		myContext.LoggerFromContext(ctx).Warnw("wine in order", "Wine id", wine.ID)
		return fmt.Errorf("wine also in order")
	}

	order.TotalPrice += wine.Price * elem.Count

	err = el.orderElemRep.Insert(ctx, elem, order)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot insert elem", "elemID", elem.ID, "error", err)
		return fmt.Errorf("create elemet order: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully create elem", "elemID", elem.ID, "error", err)
	return nil
}

func (el *OrderElemController) GetByID(ctx context.Context, ID uuid.UUID) (*models.OrderElement, error) {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get wine to get by id elem", "elemID", ID, "error", err)
		return nil, fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start get elem by id", "elemID", ID)
	element, err := el.orderElemRep.GetByID(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get elem by id", "elemID", ID, "error", err)
		return nil, fmt.Errorf("get element order: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get elem by id", "elemID", ID)
	return element, nil
}

func (el *OrderElemController) GetByOrder(ctx context.Context, IDOrder uuid.UUID) ([]*models.OrderElement, error) {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get wine to get by order elem", "orderID", IDOrder, "error", err)
		return nil, fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start get elem by order", "orderID", IDOrder)
	elements, err := el.orderElemRep.GetByOrder(ctx, IDOrder)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get elem by order", "orderID", IDOrder, "error", err)
		return nil, fmt.Errorf("get all elements order: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get elem by id", "orderID", IDOrder)
	return elements, nil
}

func (el *OrderElemController) Add(ctx context.Context, ID uuid.UUID) error {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to add elem", "elemID", ID, "error", err)
		return fmt.Errorf("get user: %w", err)
	}

	myContext.LoggerFromContext(ctx).Infow("start add elem")

	elem, err := el.orderElemRep.GetByID(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get elem", "elemID", ID, "error", err)
		return fmt.Errorf("get element order: %w", err)
	}

	wine, err := el.wineRep.GetWine(ctx, elem.IDWine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get wine to add elem", "elemID", ID, "error", err)
		return fmt.Errorf("get wine: %w", err)
	}

	if elem.Count+1 > wine.Count {
		myContext.LoggerFromContext(ctx).Warnw("not enough cnt wines to add elem", "elemID", ID, "error", err)
		return fmt.Errorf("cnt wine: %w", errors.ErrCntWine)
	}

	order, err := el.orderRep.Get(ctx, elem.IDOrder)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get order to add elem", "elemID", ID, "error", err)
		return fmt.Errorf("get order: %w", err)
	}

	order.TotalPrice += wine.Price

	err = el.orderElemRep.Add(ctx, elem, order)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot add elem", "elemID", ID, "error", err)
		return fmt.Errorf("add element order: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully add elem", "elemID", ID)
	return nil
}

func (el *OrderElemController) Decrease(ctx context.Context, ID uuid.UUID) error {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to decrease elem", "elemID", ID, "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start decrease elem")

	elem, err := el.orderElemRep.GetByID(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get elem to decrease elem", "elemID", ID, "error", err)
		return fmt.Errorf("get element order: %w", err)
	}

	wine, err := el.wineRep.GetWine(ctx, elem.IDWine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get wine to decrease elem", "elemID", ID, "error", err)
		return fmt.Errorf("get wine: %w", err)
	}

	order, err := el.orderRep.Get(ctx, elem.IDOrder)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get order to decrease elem", "elemID", ID, "error", err)
		return fmt.Errorf("get wine: %w", err)
	}

	if order.TotalPrice < wine.Price {
		myContext.LoggerFromContext(ctx).Warnw("total price less wine price")
		return fmt.Errorf("Price: %w", errors.ErrPrice)
	}
	order.TotalPrice -= wine.Price

	if elem.Count != 1 {
		err = el.orderElemRep.Decrease(ctx, elem, order)
		if err != nil {
			myContext.LoggerFromContext(ctx).Errorw("cannot decrease elem", "elemID", ID, "error", err)
			return fmt.Errorf("decrease element order: %w", err)
		}
	} else {
		err = el.orderElemRep.Delete(ctx, ID, order)
		if err != nil {
			myContext.LoggerFromContext(ctx).Errorw("cannot delete elem", "elemID", ID, "error", err)
			return fmt.Errorf("delete in decrease element order: %w", err)
		}
	}
	myContext.LoggerFromContext(ctx).Infow("successfully decrease elem")
	return nil
}

func (el *OrderElemController) Delete(ctx context.Context, ID uuid.UUID) error {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to delete elem", "elemID", ID, "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start delete elem")
	elem, err := el.orderElemRep.GetByID(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get elem to delete elem", "elemID", ID, "error", err)
		return fmt.Errorf("get element order: %w", err)
	}

	wine, err := el.wineRep.GetWine(ctx, elem.IDWine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get wine to delete elem", "elemID", ID, "error", err)
		return fmt.Errorf("get wine: %w", err)
	}

	order, err := el.orderRep.Get(ctx, elem.IDOrder)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get order to delete elem", "elemID", ID, "error", err)
		return fmt.Errorf("get order: %w", err)
	}

	if order.TotalPrice < wine.Price*elem.Count {
		myContext.LoggerFromContext(ctx).Warnw("total price less wines price")
		return fmt.Errorf("total price: %w", errors.ErrPrice)
	}
	order.TotalPrice -= wine.Price * elem.Count

	err = el.orderElemRep.Delete(ctx, ID, order)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot delete order", "elemID", ID, "error", err)
		return fmt.Errorf("delete element order: %w", err)
	}

	myContext.LoggerFromContext(ctx).Infow("successfully delete elem")
	return nil
}

func (el *OrderElemController) ResearchCreate(ctx context.Context, elem *models.OrderElement) (error, time.Duration) {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		return fmt.Errorf("get user: %w", err), 0
	}
	start := time.Now()
	wine, err := el.wineRep.GetWine(ctx, elem.IDWine)
	if err != nil {
		return fmt.Errorf("get wine: %w", err), 0
	}
	if wine.Count < elem.Count {
		return fmt.Errorf("cnt wine: %w", errors.ErrCntWine), 0
	}

	order, err := el.orderLogic.GetByUserInProcess(ctx, user.ID)
	if order == nil {
		order = &models.Order{
			IDUser:     user.ID,
			TotalPrice: 0,
			IsPoints:   false,
			Status:     models.ProcessOrder,
		}
		err = el.orderRep.Insert(ctx, order)
		if err != nil {
			return fmt.Errorf("insert order: %w", err), 0
		}
	}

	elem.IDOrder = order.ID
	if el.orderElemRep.WineInOrder(ctx, wine.ID, order.ID) == nil {
		return fmt.Errorf("wine also in order"), 0
	}

	order.TotalPrice += wine.Price * elem.Count

	err = el.orderElemRep.Insert(ctx, elem, order)
	if err != nil {
		return fmt.Errorf("create elemet order: %w", err), 0
	}
	duration := time.Since(start)
	return nil, duration
}
func (el *OrderElemController) ResearchCreateTr(ctx context.Context, elem *models.OrderElement) (error, time.Duration) {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		return fmt.Errorf("get user: %w", err), 0
	}
	start := time.Now()

	wine, err := el.wineRep.GetWine(ctx, elem.IDWine)
	if err != nil {
		return fmt.Errorf("get wine: %w", err), 0
	}

	order, err := el.orderLogic.GetByUserInProcess(ctx, user.ID)
	if order == nil {
		order = &models.Order{
			IDUser:     user.ID,
			TotalPrice: 0,
			IsPoints:   false,
			Status:     models.ProcessOrder,
		}
		err = el.orderRep.Insert(ctx, order)
		if err != nil {
			return fmt.Errorf("insert order: %w", err), 0
		}
	}

	elem.IDOrder = order.ID
	order.TotalPrice += wine.Price * elem.Count

	err = el.orderElemRep.Insert(ctx, elem, order)
	if err != nil {
		return fmt.Errorf("create elemet order: %w", err), 0
	}
	duration := time.Since(start)
	return nil, duration
}
