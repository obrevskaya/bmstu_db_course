package controllers

import (
	"context"
	"fmt"
	"net/mail"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/google/uuid"
)

const MinLengthPassword = 6

type UserController struct {
	userRep interfaces.IUserRepository
}

func NewUserController(i interfaces.IUserRepository) *UserController {
	return &UserController{
		userRep: i,
	}
}

func (u *UserController) Create(ctx context.Context, user *models.User) error {
	myContext.LoggerFromContext(ctx).Infow("start create user", "user", user.Login)
	userContext, err := myContext.UserFromContext(ctx)
	if userContext != nil {
		myContext.LoggerFromContext(ctx).Warnw("user already exists")
		return fmt.Errorf("user: %w", myErrors.ErrAlreadyExist)
	}

	if len(user.Password) < MinLengthPassword {
		myContext.LoggerFromContext(ctx).Warnw("incorrect len password")
		return fmt.Errorf("password: %w", myErrors.ErrLenPassword)
	}

	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("incorrect email")
		return fmt.Errorf("email: %w", myErrors.ErrMail)
	}

	userRegistered, _ := u.Authorize(ctx, user.Login, user.Password)
	if userRegistered != nil {
		myContext.LoggerFromContext(ctx).Warnw("user already exist", "userID", user.ID, "userName", user.Login)
		return fmt.Errorf("user: %w", myErrors.ErrAlreadyExist)
	}

	err = u.userRep.Insert(ctx, user)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("can't create user", "userID", user.ID, "userName", user.Login, "error", err)
		return fmt.Errorf("create user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully create user", "userID", user.ID, "userName", user.Login)
	return nil
}

func (u *UserController) Authorize(ctx context.Context, login string, password string) (*models.User, error) {
	myContext.LoggerFromContext(ctx).Infow("start authorize user", "login", login)
	user, err := u.userRep.Authorize(ctx, login, password)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("can't authorize user", "login", login, "error", err)
		return nil, fmt.Errorf("authorization user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully authorize user", "userID", user.ID, "login", user.Login)

	return user, nil
}

func (u *UserController) Get(ctx context.Context, ID uuid.UUID) (*models.User, error) {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to get user", "error", err)
		return nil, fmt.Errorf("get context: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("start get user", "userID", ID)
	user, err = u.userRep.Get(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("error get user", "error", err)
		return nil, fmt.Errorf("get user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully get user", "userID", user.ID, "userName", user.Login)
	return user, nil
}

func (u *UserController) UpdateUserPoints(ctx context.Context, ID uuid.UUID, points int) error {
	user, err := myContext.UserFromContext(ctx)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user to update points", "error", err)
		return fmt.Errorf("get context: %w", err)
	}

	if user.Status != models.Admin {
		myContext.LoggerFromContext(ctx).Warnw("doesn't access")
		return fmt.Errorf("access rights: %w", myErrors.ErrAccess)
	}
	myContext.LoggerFromContext(ctx).Infow("start update points")

	user, err = u.userRep.Get(ctx, ID)
	if err != nil {
		myContext.LoggerFromContext(ctx).Warnw("cannot get user", "userId", ID, "error", err)
		return fmt.Errorf("get user: %w", err)
	}

	if user.Points+points < 0 {
		myContext.LoggerFromContext(ctx).Warnw("not enough points")
		return fmt.Errorf("balance: %w", myErrors.ErrPoints)
	}

	err = u.userRep.UpdateUserPoints(ctx, ID, user.Points+points)
	if err != nil {
		myContext.LoggerFromContext(ctx).Errorw("cannot update user points", "userID", ID, "error", err)
		return fmt.Errorf("update user: %w", err)
	}
	myContext.LoggerFromContext(ctx).Infow("successfully update points")
	return nil
}
