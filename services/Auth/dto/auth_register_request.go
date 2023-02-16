package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type AuthRegisterRequest struct {
	Username string
	Email    string
	Password string
	Role     string
}

func (d *AuthRegisterRequest) FromProtoBuffer(pb *pb.RegisterRequest) {
	d.Username = pb.Username
	d.Email = pb.Email
	d.Password = pb.Password
	d.Role = pb.Role.String()
}

func (d *AuthRegisterRequest) ToProtoBuffer() *pb.RegisterRequest {
	return &pb.RegisterRequest{
		Username: d.Username,
		Email:    d.Email,
		Password: d.Password,
		Role:     *pb.Role(pb.Role_value[d.Role]).Enum(),
	}
}
