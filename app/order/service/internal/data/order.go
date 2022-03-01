package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nextzeus/kratos-shop/app/order/service/internal/biz"
	"github.com/nextzeus/kratos-shop/pkg/util/pagination"
	"gorm.io/gorm"
	"time"
)

var _ biz.OrderRepo = (*orderRepo)(nil)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

type Order struct {
	gorm.Model
	Id        int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/order")),
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, o *biz.Order) (*biz.Order, error) {
	or := Order{Id: o.Id, UserId: o.UserId}
	result := r.data.db.WithContext(ctx).Create(&or)
	return &biz.Order{
		Id: o.Id,
	}, result.Error
}

func (r *orderRepo) UpdateOrder(ctx context.Context, o *biz.Order) (*biz.Order, error) {
	or := Order{}
	result := r.data.db.WithContext(ctx).First(&or, o.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	or.UserId = o.UserId
	result = r.data.db.WithContext(ctx).Save(&or)
	return &biz.Order{
		Id: or.Id,
	}, result.Error
}

func (r *orderRepo) GetOrder(ctx context.Context, id int64) (*biz.Order, error) {
	o := Order{}
	result := r.data.db.WithContext(ctx).First(&o, id)
	return &biz.Order{
		Id: o.Id,
	}, result.Error
}

func (r *orderRepo) ListOrder(ctx context.Context, pageNum, pageSize int64) ([]*biz.Order, error) {
	var os []Order
	result := r.data.db.WithContext(ctx).Limit(int(pageSize)).
		Offset(int(pagination.GetPagOffset(pageNum, pageSize))).
		Find(&os)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.Order, 0)
	for _, o := range os {
		rv = append(rv, &biz.Order{
			Id: o.Id,
		})
	}
	return rv, nil
}
