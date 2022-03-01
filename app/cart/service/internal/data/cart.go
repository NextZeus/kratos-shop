package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/biz"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ biz.CartRepo = (*cartRepo)(nil)

type cartRepo struct {
	data     *Data
	cartColl *mongo.Collection
	log      *log.Helper
}

type ItemSchema struct {
	ItemId   int64 `bson:"itemId"`
	Quantity int64 `bson:"quantity"`
}

// CartSchema mongodb 数据库结构定义
type CartSchema struct {
	UserId int64        `bson:"uid"`
	Items  []ItemSchema `bson:"items"`
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data:     data,
		cartColl: data.db.Collection("cart"),
		log:      log.NewHelper(log.With(logger, "module", "repo/cart")),
	}
}

func (cr cartRepo) GetCart(ctx context.Context, uid int64) (*biz.Cart, error) {
	result := &CartSchema{}
	if err := cr.cartColl.FindOne(ctx, bson.M{"uid": uid}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return &biz.Cart{UserId: result.UserId}, nil
		}
		return nil, err
	}
	items := make([]biz.Item, 0)
	for _, item := range result.Items {
		items = append(items, biz.Item{
			Id:       item.ItemId,
			Quantity: item.Quantity,
		})
	}
	return &biz.Cart{
		UserId: result.UserId,
		Items:  items,
	}, nil
}

func (cr cartRepo) SaveCart(ctx context.Context, c *biz.Cart) error {
	items := bson.A{}
	for _, item := range c.Items {
		items = append(items, bson.M{"itemId": item.Id, "quantity": item.Quantity})
	}

	result := cr.cartColl.FindOneAndUpdate(
		ctx,
		bson.M{"uid": c.UserId},
		bson.D{
			{"uid", c.UserId},
			{"items", items},
		},
	)

	return result.Err()
}

func (cr cartRepo) DeleteCart(ctx context.Context, uid int64) error {
	_, err := cr.cartColl.DeleteOne(ctx, bson.M{"uid": uid})
	return err
}
