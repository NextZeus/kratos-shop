package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/nextzeus/kratos-shop/api/user/service/v1"
	"math/rand"
)

var (
	ErrUserNotFound = errors.New("server not found")
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "useCase/server"))}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) Save(ctx context.Context, in *v1.SaveUserRequest) (*v1.SaveUserReply, error) {
	user := &User{
		Id:       rand.Int63(),
		Username: in.Name,
		Password: in.Password,
	}
	_, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.SaveUserReply{
		Id: user.Id,
	}, nil
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) GetByUserName(ctx context.Context, in *v1.GetUserByUserNameRequest) (*v1.GetUserByUserNameReply, error) {
	user, err := uc.repo.FindByUsername(ctx, in.Username)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserByUserNameReply{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}
