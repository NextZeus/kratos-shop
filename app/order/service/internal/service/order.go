package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nextzeus/kratos-shop/app/order/service/internal/biz"

	pb "github.com/nextzeus/kratos-shop/api/order/service/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer

	oc  *biz.OrderUseCase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderUseCase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "server/order")),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	x, err := s.oc.Create(ctx, &biz.Order{})
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderReply{
		Id: x.Id,
	}, nil
}
func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderReply, error) {
	x, err := s.oc.Update(ctx, &biz.Order{})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateOrderReply{
		Id: x.Id,
	}, nil
}
func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	x, err := s.oc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderReply{
		Id: x.Id,
	}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	rv, err := s.oc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	rs := make([]*pb.ListOrderReply_Order, 0)
	for _, x := range rv {
		rs = append(rs, &pb.ListOrderReply_Order{
			Id: x.Id,
		})
	}
	return &pb.ListOrderReply{
		Orders: rs,
	}, nil
}
