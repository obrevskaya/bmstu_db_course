package http

import (
	"net/http"

	mycontext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	"go.uber.org/zap"
)

func middleware(u interfaces.IUserController, logger *zap.SugaredLogger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := u.Authorize(ctx, r.Header.Get("Login"), r.Header.Get("Password"))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx = mycontext.UserToContext(ctx, user)
		ctx = mycontext.LoggerToContext(ctx, logger)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
