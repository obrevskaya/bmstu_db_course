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

type WineController struct {
	wineRep interfaces.IWineRepository
}

func NewWineController(i interfaces.IWineRepository) *WineController {
	return &WineController{
		wineRep: i,
	}
}

func (w *WineController) Create(ctx context.Context, wine *models.Wine) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to create wine", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	if user.Status < models.Admin {
		myContext.LoggerFromContext(ctx).Warnw("don't access rights")
		return fmt.Errorf("access rights: %w", myErrors.ErrAccess)
	}

	myContext.LoggerFromContext(ctx).Infow("start create wine")

	err = w.wineRep.Insert(ctx, wine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot create wine", "error", err)
		return fmt.Errorf("create wine: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully create wine")
	return nil
}

func (w *WineController) GetWine(ctx context.Context, ID uuid.UUID) (*models.Wine, error) {
	myContext.LoggerFromContext(ctx).Infow("start get wine")

	wine, err := w.wineRep.GetWine(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get wine", "error", err)
		return nil, fmt.Errorf("get wine: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get wine")
	return wine, nil
}

func (w *WineController) GetWines(ctx context.Context, limit int, skip int) ([]*models.Wine, error) {
	myContext.LoggerFromContext(ctx).Infow("start get wines")

	wines, err := w.wineRep.GetWines(ctx, limit, skip)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get wines", "error", err)
		return nil, fmt.Errorf("get wines: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get wines")
	return wines, nil
}

func (w *WineController) Delete(ctx context.Context, ID uuid.UUID) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to delete wine", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	if user.Status < models.Admin {
		myContext.LoggerFromContext(ctx).Warnw("don't access rights")
		return fmt.Errorf("access rights: %w", myErrors.ErrAccess)
	}
	myContext.LoggerFromContext(ctx).Infow("start delete wine")

	err = w.wineRep.Delete(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot delete wine", "error", err)
		return fmt.Errorf("delete wine: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully wine")
	return nil
}

func (w *WineController) Update(ctx context.Context, wine *models.Wine) error {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to update wine", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start update wine")

	err = w.wineRep.Update(ctx, wine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot update wine", "error", err)
		return fmt.Errorf("update wine: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully update wine")
	return nil
}
