package dto

import (
	"github.com/RhnAdi/elearning-microservice/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateClassroomRequest struct {
	Name        string
	Description string
	TeacherId   primitive.ObjectID
}

func (d *CreateClassroomRequest) FromProtoBuffer(req *pb.CreateClassroomRequest) {
	d.Name = req.Name
	d.Description = req.Description
}

func (d *CreateClassroomRequest) ToProtoBuffer() (res pb.CreateClassroomRequest) {
	res.Name = d.Name
	res.Description = d.Description
	return
}
