package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/biz"

	pb "github.com/nextzeus/kratos-shop/api/cart/service/v1"
)

type CartService struct {
	pb.UnimplementedCartServer

	cc  *biz.CartUseCase
	log *log.Helper
}

func NewCartService(cc *biz.CartUseCase, logger log.Logger) *CartService {
	return &CartService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/cart")),
	}
}

func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (reply *pb.GetCartReply, err error) {
	reply = &pb.GetCartReply{
		Items: make([]*pb.Item, 0),
	}
	c, err := s.cc.GetCart(ctx, req.Uid)
	if err != nil {
		return reply, err
	}
	for _, item := range c.Items {
		reply.Items = append(reply.Items, &pb.Item{
			ItemId:   item.Id,
			Quantity: item.Quantity,
		})
	}
	return
}
func (s *CartService) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartReply, error) {
	return &pb.DeleteCartReply{}, nil
}
func (s *CartService) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemReply, error) {
	return &pb.AddItemReply{}, nil
}
func (s *CartService) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemReply, error) {
	return &pb.UpdateItemReply{}, nil
}
func (s *CartService) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemReply, error) {
	return &pb.DeleteItemReply{}, nil
}
