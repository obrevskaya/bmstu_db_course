package controllers

import (
	"context"
	"errors"
	"testing"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/containers"
	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres"
	"github.com/google/uuid"
)

func TestUserController_UpdateUserPoints(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	userRepo := postgres.NewUR(db)

	type fields struct {
		userRep interfaces.IUserRepository
	}
	type args struct {
		ctx    context.Context
		ID     uuid.UUID
		points int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(u *UserController) error
	}{
		{
			name:   "successful update user",
			fields: fields{userRep: userRepo},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   50,
					Status:   models.Admin,
				}),
				ID:     uuid.MustParse("4ec2b1fb-1102-4ce1-b0d5-2b7da62d0023"),
				points: 200,
			},
			wantErr: false,
			check: func(u *UserController) error {
				user, err := u.Get(myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}), uuid.MustParse("4ec2b1fb-1102-4ce1-b0d5-2b7da62d0023"))
				if err != nil {
					return err
				}
				if user.Points != 300 {
					return errors.New("not update points user")
				}
				return nil
			},
		},
		{
			name:   "no context in update user",
			fields: fields{userRep: userRepo},
			args: args{
				ctx:    context.Background(),
				ID:     uuid.New(),
				points: 100,
			},
			wantErr: true,
		},
		{
			name:   "no access rights in update user",
			fields: fields{userRep: userRepo},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Customer,
				}),
				ID:     uuid.MustParse("4ec2b1fb-1102-4ce1-b0d5-2b7da62d0023"),
				points: 100,
			},
			wantErr: true,
		},
		{
			name:   "balance error",
			fields: fields{userRep: userRepo},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}),
				ID:     uuid.MustParse("4ec2b1fb-1102-4ce1-b0d5-2b7da62d0023"),
				points: -3000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				userRep: tt.fields.userRep,
			}
			if err := u.UpdateUserPoints(tt.args.ctx, tt.args.ID, tt.args.points); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserPoints() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
