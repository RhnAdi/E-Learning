package dto

import (
	"time"

	"github.com/RhnAdi/elearning-microservice/pb"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClassroomResponse struct {
	Id          primitive.ObjectID
	TeacherId   primitive.ObjectID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d *ClassroomResponse) FromProtoBuffer(req *pb.ClassroomResponse) (err error) {
	d.Id, err = primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return
	}
	d.TeacherId, err = primitive.ObjectIDFromHex(req.TeacherId)
	if err != nil {
		return
	}
	d.Name = req.Name
	d.Description = req.Description
	d.CreatedAt = time.Unix(req.CreatedAt, 0)
	d.UpdatedAt = time.Unix(req.UpdatedAt, 0)
	return nil
}

func (d *ClassroomResponse) ToProtoBuffer() *pb.ClassroomResponse {
	res := new(pb.ClassroomResponse)
	res.Id = d.Id.Hex()
	res.TeacherId = d.TeacherId.Hex()
	res.Name = d.Name
	res.Description = d.Description
	res.CreatedAt = d.CreatedAt.Unix()
	res.UpdatedAt = d.UpdatedAt.Unix()
	return res
}

func (d *ClassroomResponse) FromModelClassroom(model *models.Classroom) {
	d.Id = model.Id
	d.TeacherId = model.TeacherId
	d.Name = model.Name
	d.Description = model.Description
	d.CreatedAt = model.CreatedAt
	d.UpdatedAt = model.UpdatedAt
}
