package models

import (
	"time"

	"github.com/RhnAdi/elearning-microservice/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id, omitempty"`
	Username  string             `bson:"username, omitempty"`
	Email     string             `bson:"email, omitempty"`
	Role      string             `bson:"role, omitempty"`
	Password  string             `bson:"password, omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (u *User) PreSave() (err error) {
	u.Id = primitive.NewObjectID()
	hash_password, err := security.EcryptPassword(u.Password)
	u.Password = hash_password
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return err
}

func (u *User) VerifyPassword(password string) error {
	return security.VerifyPassword(u.Password, password)
}
