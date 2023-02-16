package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
}

func (d *AuthResponse) FromProtoBuffer(pb *pb.AuthResponse) {
	d.AccessToken = pb.AccessToken
	d.RefreshToken = pb.RefreshToken
}

func (d *AuthResponse) ToProtoBuffer() *pb.AuthResponse {
	return &pb.AuthResponse{
		AccessToken:  d.AccessToken,
		RefreshToken: d.RefreshToken,
	}
}
