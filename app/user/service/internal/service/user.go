package service

import (
	"context"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/biz"

	pb "github.com/nextzeus/kratos-shop/api/user/service/v1"
)

func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordRequest) (*pb.VerifyPasswordReply, error) {
	rv, err := s.uc.VerifyPassword(ctx, &biz.User{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &pb.VerifyPasswordReply{
		Ok: rv,
	}, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}
func (s *UserService) Save(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserReply, error) {
	return s.uc.Save(ctx, req)
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, nil
}
func (s *UserService) GetUserByUserName(ctx context.Context, req *pb.GetUserByUserNameRequest) (*pb.GetUserByUserNameReply, error) {
	return s.uc.GetByUserName(ctx, req)
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *UserService) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.CreateAddressReply, error) {
	return &pb.CreateAddressReply{}, nil
}
func (s *UserService) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressReply, error) {
	return &pb.GetAddressReply{}, nil
}
func (s *UserService) ListAddresses(ctx context.Context, req *pb.ListAddressesRequest) (*pb.ListAddressesReply, error) {
	return &pb.ListAddressesReply{}, nil
}
func (s *UserService) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.CreateCardReply, error) {
	return &pb.CreateCardReply{}, nil
}
func (s *UserService) GetCard(ctx context.Context, req *pb.GetCardRequest) (*pb.GetCardReply, error) {
	return &pb.GetCardReply{}, nil
}
func (s *UserService) ListCard(ctx context.Context, req *pb.ListCardRequest) (*pb.ListCardReply, error) {
	return &pb.ListCardReply{}, nil
}
func (s *UserService) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteCardReply, error) {
	return &pb.DeleteCardReply{}, nil
}
