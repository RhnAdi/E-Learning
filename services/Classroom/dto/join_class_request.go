package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type JoinClassRequest struct {
	StudentId   primitive.ObjectID
	ClassroomId primitive.ObjectID
}
