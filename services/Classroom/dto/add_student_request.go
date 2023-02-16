package dto

import (
	"github.com/RhnAdi/elearning-microservice/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddStudentRequest struct {
	ClassroomId primitive.ObjectID
	StudentId   primitive.ObjectID
}

func (d *AddStudentRequest) FromProtoBuffer(pb *pb.AddStudentRequest) (err error) {
	d.ClassroomId, err = primitive.ObjectIDFromHex(pb.ClassroomId)
	if err != nil {
		return
	}
	d.StudentId, err = primitive.ObjectIDFromHex(pb.StudentId)
	if err != nil {
		return
	}
	return
}

func (d *AddStudentRequest) ToProtoBuffer() (pb pb.AddStudentRequest) {
	pb.ClassroomId = d.ClassroomId.Hex()
	pb.StudentId = d.StudentId.Hex()
	return
}
