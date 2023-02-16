package dto

import (
	"time"

	"github.com/RhnAdi/elearning-microservice/pb"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentClassResponse struct {
	Id          primitive.ObjectID
	ClassroomId primitive.ObjectID
	StudentId   primitive.ObjectID
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d *StudentClassResponse) FromProtoBuffer(pb *pb.StudentClassResponse) (err error) {
	d.Id, err = primitive.ObjectIDFromHex(pb.Id)
	if err != nil {
		return
	}
	d.ClassroomId, err = primitive.ObjectIDFromHex(pb.ClassroomId)
	if err != nil {
		return
	}
	d.StudentId, err = primitive.ObjectIDFromHex(pb.StudentId)
	if err != nil {
		return
	}
	d.Status = pb.Status
	d.CreatedAt = time.Unix(pb.CreatedAt, 0)
	d.UpdatedAt = time.Unix(pb.UpdatedAt, 0)
	return nil
}

func (d *StudentClassResponse) ToProtoBuffer() (pb pb.StudentClassResponse) {
	pb.Id = d.Id.Hex()
	pb.ClassroomId = d.ClassroomId.Hex()
	pb.StudentId = d.StudentId.Hex()
	pb.Status = d.Status
	pb.CreatedAt = d.CreatedAt.Unix()
	pb.UpdatedAt = d.UpdatedAt.Unix()
	return
}

func (d *StudentClassResponse) FromStudentClassModel(req *models.StudentClass) {
	d.Id = req.Id
	d.ClassroomId = req.ClassroomId
	d.StudentId = req.StudentId
	d.Status = req.Status
	d.CreatedAt = req.CreatedAt
	d.UpdatedAt = req.UpdatedAt
}
