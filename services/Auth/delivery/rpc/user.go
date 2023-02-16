package rpc

import (
	"context"

	"github.com/RhnAdi/elearning-microservice/pb"
	"github.com/RhnAdi/elearning-microservice/services/Auth/dto"
	"github.com/RhnAdi/elearning-microservice/services/Auth/usecase"
)

type userRpc struct {
	userUsecase usecase.UserUsecase
	*pb.UnimplementedUserServiceServer
}

func NewUserRpc(userUsecase usecase.UserUsecase) pb.UserServiceServer {
	return &userRpc{
		userUsecase: userUsecase,
	}
}

func (d *userRpc) GetUserById(ctx context.Context, req *pb.Id) (*pb.UserResponse, error) {
	res, err := d.userUsecase.GetUserById(req.Id)
	resPb := res.ToProtoBuffer()
	return &resPb, err
}

func (d *userRpc) GetUserByEmail(ctx context.Context, req *pb.Email) (*pb.UserResponse, error) {
	res, err := d.userUsecase.GetUserByEmail(req.Email)

	resPb := res.ToProtoBuffer()
	return &resPb, err
}

func (d *userRpc) GetAllUser(ctx context.Context, req *pb.Empty) (*pb.ListUserResponse, error) {
	res, err := d.userUsecase.GetAllUser()
	resPb := res.ToProtoBuffer()
	return &resPb, err
}

func (d *userRpc) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	convertReq := new(dto.UpdateUserRequest)
	convertReq.FromProtoBuffer(req)

	res, err := d.userUsecase.UpdateUser(req.Id, convertReq)

	resPb := res.ToProtoBuffer()
	return &resPb, err
}

func (d *userRpc) DeleteUser(ctx context.Context, req *pb.Id) (*pb.Id, error) {
	res, err := d.userUsecase.DeleteUser(req.Id)

	return &pb.Id{Id: res.Id.Hex()}, err
}
