package controllers

import (
	"context"
	"reflect"
	"testing"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/containers"
	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres"
	"github.com/google/uuid"
)

func TestOrderController_Get(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	orderRep := postgres.NewOR(db)

	o := &models.Order{
		ID:         uuid.MustParse("b8b87f28-3cbb-425c-ac4d-52015710d61b"),
		IDUser:     uuid.MustParse("06100e99-e631-474f-85e4-e5e794925f67"),
		TotalPrice: 28519,
		IsPoints:   false,
		Status:     "placed",
	}

	type fields struct {
		orderRep interfaces.IOrderRepository
		userRep  interfaces.IUserRepository
		billRep  interfaces.IBillRepository
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Order
		wantErr bool
	}{
		{
			name: "successful get order",
			fields: fields{
				orderRep: orderRep,
				billRep:  nil,
				userRep:  nil,
			},
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
				ID: uuid.MustParse("b8b87f28-3cbb-425c-ac4d-52015710d61b"),
			},
			want:    o,
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderRep: orderRep,
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.MustParse("b8b87f28-3cbb-425c-ac4d-52015710d61b"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error get order",
			fields: fields{
				orderRep: orderRep,
				billRep:  nil,
				userRep:  nil,
			},
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
				ID: uuid.MustParse("b8b87f28-3cbb-425c-ac4d-52015990d61b"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderController{
				orderRep: tt.fields.orderRep,
				userRep:  tt.fields.userRep,
				billRep:  tt.fields.billRep,
			}
			got, err := o.GetByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
