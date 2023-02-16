package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type CreateUserRequest struct {
	Username string
	Email    string
	Password string
	Role     string
}

func (d *CreateUserRequest) FromProtoBuffer(pb *pb.CreateUserRequest) {
	d.Username = pb.Username
	d.Email = pb.Email
	d.Password = pb.Password
	d.Role = pb.Role
}

func (d *CreateUserRequest) ToProtoBuffer() *pb.CreateUserRequest {
	return &pb.CreateUserRequest{
		Username: d.Username,
		Email:    d.Email,
		Password: d.Password,
		Role:     d.Role,
	}
}
