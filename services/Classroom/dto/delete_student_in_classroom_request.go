package dto

import (
	"github.com/RhnAdi/elearning-microservice/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteStudentInClassroom struct {
	ClassroomId primitive.ObjectID
	StudentId   primitive.ObjectID
}

func (d *DeleteStudentInClassroom) FromProtoBuffer(pb *pb.DeleteStudentInClassRequest) (err error) {
	d.ClassroomId, err = primitive.ObjectIDFromHex(pb.ClassroomId)
	if err != nil {
		return
	}
	d.StudentId, err = primitive.ObjectIDFromHex(pb.StudentId)
	if err != nil {
		return
	}
	return nil
}

func (d *DeleteStudentInClassroom) ToProtoBuffer() (pb pb.DeleteStudentInClassRequest) {
	pb.ClassroomId = d.ClassroomId.Hex()
	pb.StudentId = d.StudentId.Hex()
	return
}
