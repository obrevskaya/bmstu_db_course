package context

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"go.uber.org/zap"
)

const (
	key    = "user"
	keyLog = "logger"
)

func UserToContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (*models.User, error) {
	user := ctx.Value(key)
	if user == nil {
		return nil, fmt.Errorf("user error: %w", errors.ErrGet)
	}

	return user.(*models.User), nil
}

func LoggerToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, keyLog, logger)
}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	v := ctx.Value(keyLog)
	if v == nil {
		return zap.NewNop().Sugar()
	}
	return v.(*zap.SugaredLogger)
}
