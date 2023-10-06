package controllers

import (
	"context"
	"fmt"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/google/uuid"
)

type UserWineController struct {
	userWineRep interfaces.IUserWinesRepository
}

func NewUserWineController(uw interfaces.IUserWinesRepository) *UserWineController {
	return &UserWineController{
		userWineRep: uw,
	}
}

func (uw *UserWineController) Create(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to create user wine", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start create user wine")
	err = uw.userWineRep.Insert(ctx, IDUser, IDWine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot create user wine", "error", err)
		return fmt.Errorf("create user wine: %w", err)
	}

	myContext.LoggerFromContext(ctx).Infow("successfully create user wine")

	return nil
}

func (uw *UserWineController) GetByUser(ctx context.Context, IDUser uuid.UUID) ([]*models.UserWine, error) {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to get user wine", "error", err)
		return nil, fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start get user wine")
	userWines, err := uw.userWineRep.GetByUser(ctx, IDUser)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot get user wine", "error", err)
		return nil, fmt.Errorf("get user wine: %w", err)
	}

	myContext.LoggerFromContext(ctx).Infow("successfully get user wine")

	return userWines, nil
}

func (uw *UserWineController) DeleteWine(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error {
	_, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to delete user wine", "error", err)
		return fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start delete user wine")
	err = uw.userWineRep.DeleteWine(ctx, IDUser, IDWine)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot delete user wine", "error", err)
		return fmt.Errorf("delete user wine: %w", err)
	}

	myContext.LoggerFromContext(ctx).Infow("successfully delete user wine")

	return nil
}
