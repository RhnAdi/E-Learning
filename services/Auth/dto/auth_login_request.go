package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type AuthLoginRequest struct {
	Email    string
	Password string
}

func (d *AuthLoginRequest) FromProtoBuffer(pb *pb.LoginRequest) {
	d.Email = pb.Email
	d.Password = pb.Password
}

func (d *AuthLoginRequest) ToProtoBuffer() *pb.LoginRequest {
	return &pb.LoginRequest{
		Email:    d.Email,
		Password: d.Password,
	}
}
