package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Classroom struct {
	Id          primitive.ObjectID `bson:"_id, omitempty"`
	TeacherId   primitive.ObjectID `bson:"teacher_id, omitempty"`
	Name        string             `bson:"name, omitempty"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (s *Classroom) Pre() {
	s.Id = primitive.NewObjectID()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

type StudentClass struct {
	Id          primitive.ObjectID `bson:"_id, omitempty"`
	ClassroomId primitive.ObjectID `bson:"classroom_id, omitempty"`
	StudentId   primitive.ObjectID `bson:"student_id, omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Status      bool               `bson:"status"`
}

func (cs *StudentClass) Pre() {
	cs.Id = primitive.NewObjectID()
	cs.CreatedAt = time.Now()
	cs.UpdatedAt = time.Now()
}
