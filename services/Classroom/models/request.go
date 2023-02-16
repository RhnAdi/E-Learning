package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JoinRequest struct {
	Id          primitive.ObjectID `bson:"_id, omitempty"`
	StudentId   primitive.ObjectID `bson:"student_id, omitempty"`
	ClassroomId primitive.ObjectID `bson:"classroom_id"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (jr *JoinRequest) Pre() {
	jr.Id = primitive.NewObjectID()
	jr.Status = "pending"
	jr.CreatedAt = time.Now()
	jr.UpdatedAt = time.Now()
}
