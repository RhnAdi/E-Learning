package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentClassRequest struct {
	ClassroomId primitive.ObjectID
	StudentId   primitive.ObjectID
}
