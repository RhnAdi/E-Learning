package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type UpdateUserRequest struct {
	Username string
}

func (d *UpdateUserRequest) FromProtoBuffer(pb *pb.UpdateUserRequest) {
	d.Username = pb.Username
}

func (d *UpdateUserRequest) ToProtoBuffer() *pb.UpdateUserRequest {
	return &pb.UpdateUserRequest{
		Username: d.Username,
	}
}
