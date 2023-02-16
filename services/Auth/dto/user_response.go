package dto

import (
	"time"

	"github.com/RhnAdi/elearning-microservice/pb"
	"github.com/RhnAdi/elearning-microservice/services/Auth/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Id        primitive.ObjectID `json:"_id"`
	Role      string             `json:"role"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func (u *UserResponse) ToProtoBuffer() pb.UserResponse {
	return pb.UserResponse{
		Id:        u.Id.Hex(),
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}
}

func (u *UserResponse) FromProtoBuffer(user *pb.UserResponse) (err error) {
	u.Id, err = primitive.ObjectIDFromHex(user.GetId())
	if err != nil {
		return
	}
	u.Username = user.GetUsername()
	u.Email = user.GetEmail()
	u.Role = user.GetRole()
	u.CreatedAt = time.Unix(user.GetCreatedAt(), 0)
	u.UpdatedAt = time.Unix(user.GetUpdatedAt(), 0)
	return
}

func (u *UserResponse) FromUserModel(user models.User) {
	u.Id = user.Id
	u.Email = user.Email
	u.Username = user.Username
	u.Role = user.Role
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}
