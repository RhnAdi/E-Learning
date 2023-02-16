package dto

import (
	"github.com/RhnAdi/elearning-microservice/pb"
)

type ListUserResponse struct {
	Users []*UserResponse
}

func (d *ListUserResponse) FromProtoBuffer(pb *pb.ListUserResponse) {
	for _, user := range pb.Users {
		data := new(UserResponse)
		data.FromProtoBuffer(user)
		d.Users = append(d.Users, data)
	}
}

func (d *ListUserResponse) ToProtoBuffer() (res pb.ListUserResponse) {
	for _, user := range d.Users {
		tmp := user.ToProtoBuffer()
		res.Users = append(res.Users, &tmp)
	}
	return
}
