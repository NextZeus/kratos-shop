package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Card 结构体定义
type Card struct {
	Id      int64
	CardNo  string
	CCV     string
	Expires string
}

// CardRepo 接口定义
type CardRepo interface {
	CreateCard(ctx context.Context, c *Card) (*Card, error)
	GetCard(ctx context.Context, id int64) (*Card, error)
	ListCard(ctx context.Context, uid int64) ([]*Card, error)
}

type CardUseCase struct {
	repo CardRepo
	log  *log.Helper
}

func NewCardUseCase(repo CardRepo, logger log.Logger) *CardUseCase {
	return &CardUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/card"))}
}

func (uc *CardUseCase) Create(ctx context.Context, c *Card) (*Card, error) {
	return uc.repo.CreateCard(ctx, c)
}

func (uc *CardUseCase) Get(ctx context.Context, id int64) (*Card, error) {
	return uc.repo.GetCard(ctx, id)
}

func (uc *CardUseCase) ListCard(ctx context.Context, uid int64) ([]*Card, error) {
	return uc.repo.ListCard(ctx, uid)
}
