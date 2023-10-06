package controllers

import (
	"context"
	"reflect"
	"testing"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestUser_Authorize(t *testing.T) {
	mc := minimock.NewController(t)
	u := &models.User{
		ID:       uuid.New(),
		Login:    "obrevskaya",
		Password: "qwerty1234",
		Fio:      "Obrevskaya Veronika",
		Email:    "obrevskaya.vera@mail.ru",
		Points:   0,
		Status:   models.Customer,
	}

	type fields struct {
		userRep interfaces.IUserRepository
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:   "successful Authorize",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).AuthorizeMock.Return(u, nil)},
			args: args{
				ctx:      context.Background(),
				login:    "obrevskaya",
				password: "qwerty1234",
			},
			want:    u,
			wantErr: false,
		},
		{
			name:   "error Authorize in data",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).AuthorizeMock.Return(u, myErrors.ErrInsert)},
			args: args{
				ctx:      context.Background(),
				login:    "obrevskaya",
				password: "qwerty1234",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				userRep: tt.fields.userRep,
			}
			got, err := u.Authorize(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authorize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Create(t *testing.T) {
	mc := minimock.NewController(t)

	u := &models.User{
		ID:       uuid.New(),
		Login:    "Kir",
		Password: "qwertyuiop[",
		Email:    "1@mail.ru",
		Points:   1,
	}

	type fields struct {
		userRep interfaces.IUserRepository
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "successful create user",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).InsertMock.Return(nil).AuthorizeMock.Return(nil, myErrors.ErrGet)},
			args: args{
				ctx: myContext.UserToContext(context.Background(), nil),
				user: &models.User{
					ID:       uuid.New(),
					Login:    "ovanov",
					Password: "qwerty1235",
					Fio:      "Ovanov Vova",
					Email:    "ivanov.vera@mail.ru",
					Points:   0,
					Status:   models.Customer},
			},
			wantErr: false,
		},
		{
			name:   "user already exists",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).InsertMock.Return(nil)},
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
				user: &models.User{
					ID:       uuid.New(),
					Login:    "ovanov",
					Password: "qwerty1235",
					Fio:      "Ovanov Vova",
					Email:    "ivanov.vera@mail.ru",
					Points:   0,
					Status:   models.Customer},
			},
			wantErr: true,
		},
		{
			name:   "password is too short",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).InsertMock.Return(nil)},
			args: args{
				ctx: myContext.UserToContext(context.Background(), nil),
				user: &models.User{
					ID:       uuid.New(),
					Login:    "ovanov",
					Password: "q",
					Fio:      "Ovanov Vova",
					Email:    "ivanov.vera@mail.ru",
					Points:   0,
					Status:   models.Customer},
			},
			wantErr: true,
		},
		{
			name:   "email is wrong",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).InsertMock.Return(nil)},
			args: args{
				ctx: myContext.UserToContext(context.Background(), nil),
				user: &models.User{
					ID:       uuid.New(),
					Login:    "ovanov",
					Password: "qweerty56468",
					Fio:      "Ovanov Vova",
					Email:    "ivanov.verl.ru",
					Points:   0,
					Status:   models.Customer},
			},
			wantErr: true,
		},
		{
			name:   "error user already authorize",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).InsertMock.Return(nil).AuthorizeMock.Return(u, nil)},
			args: args{
				ctx: myContext.UserToContext(context.Background(), nil),
				user: &models.User{
					ID:       uuid.New(),
					Login:    "ovanov",
					Password: "qwerty12390",
					Fio:      "Ovanov Vova",
					Email:    "ivanov.vera@mail.ru",
					Points:   0,
					Status:   models.Customer},
			},
			wantErr: true,
		},
		{
			name:   "error create in user data",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).InsertMock.Return(myErrors.ErrInsert).AuthorizeMock.Return(nil, myErrors.ErrGet)},
			args: args{
				ctx: myContext.UserToContext(context.Background(), nil),
				user: &models.User{
					ID:       uuid.New(),
					Login:    "ovanov",
					Password: "qwerty12390",
					Fio:      "Ovanov Vova",
					Email:    "ivanov.vera@mail.ru",
					Points:   0,
					Status:   models.Customer},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				userRep: tt.fields.userRep,
			}
			if err := u.Create(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Get(t *testing.T) {
	mc := minimock.NewController(t)

	u := &models.User{
		ID:       uuid.New(),
		Login:    "ovanov",
		Password: "qwerty1235",
		Fio:      "Ovanov vova",
		Email:    "obr.vera@mail.ru",
		Points:   0,
		Status:   models.Customer,
	}

	type fields struct {
		userRep interfaces.IUserRepository
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:   "successful get user",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).GetMock.Return(u, nil)},
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
				ID: uuid.New(),
			},
			want:    u,
			wantErr: false,
		},
		{
			name:   "no context in get user",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).GetMock.Return(u, nil)},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "error in get user data",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).GetMock.Return(u, myErrors.ErrGetDB)},
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
				ID: uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				userRep: tt.fields.userRep,
			}
			got, err := u.Get(tt.args.ctx, tt.args.ID)
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

func TestUser_UpdateUserPoints(t *testing.T) {
	mc := minimock.NewController(t)
	u := &models.User{
		ID:       uuid.New(),
		Login:    "ovanov",
		Password: "qwerty1235",
		Fio:      "Ovanov vova",
		Email:    "obr.vera@mail.ru",
		Points:   10,
		Status:   models.Customer,
	}
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
	}{
		{
			name:   "successful update user",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).UpdateUserPointsMock.Return(nil).GetMock.Return(u, nil)},
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
				ID:     uuid.New(),
				points: 100,
			},
			wantErr: false,
		},
		{
			name:   "no context in update user",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).UpdateUserPointsMock.Return(nil).GetMock.Return(u, nil)},
			args: args{
				ctx:    context.Background(),
				ID:     uuid.New(),
				points: 100,
			},
			wantErr: true,
		},
		{
			name:   "no access rights in update user",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).UpdateUserPointsMock.Return(nil).GetMock.Return(u, nil)},
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
				ID:     uuid.New(),
				points: 100,
			},
			wantErr: true,
		},
		{
			name:   "get user error",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).UpdateUserPointsMock.Return(nil).GetMock.Return(u, myErrors.ErrGetDB)},
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
				ID:     uuid.New(),
				points: 100,
			},
			wantErr: true,
		},
		{
			name:   "balance error",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).UpdateUserPointsMock.Return(nil).GetMock.Return(u, nil)},
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
				ID:     uuid.New(),
				points: -100,
			},
			wantErr: true,
		},
		{
			name:   "update user error in data",
			fields: fields{userRep: mocks.NewIUserRepositoryMock(mc).UpdateUserPointsMock.Return(myErrors.ErrUpdate).GetMock.Return(u, nil)},
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
				ID:     uuid.New(),
				points: 100,
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
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
