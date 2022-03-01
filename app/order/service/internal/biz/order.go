package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	Id     int64
	UserId int64
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, c *Order) (*Order, error)
	GetOrder(ctx context.Context, id int64) (*Order, error)
	UpdateOrder(context.Context, *Order) (*Order, error)
	ListOrder(ctx context.Context, pageNum int64, pageSize int64) ([]*Order, error)
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/order"))}
}

func (uc *OrderUseCase) Create(ctx context.Context, o *Order) (*Order, error) {
	return uc.repo.CreateOrder(ctx, o)
}

func (uc *OrderUseCase) Update(ctx context.Context, o *Order) (*Order, error) {
	return uc.repo.UpdateOrder(ctx, o)
}

func (uc *OrderUseCase) Get(ctx context.Context, id int64) (*Order, error) {
	return uc.repo.GetOrder(ctx, id)
}

func (uc *OrderUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Order, error) {
	return uc.repo.ListOrder(ctx, pageNum, pageSize)
}
