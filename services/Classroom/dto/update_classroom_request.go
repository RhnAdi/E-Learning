package dto

import (
	"github.com/RhnAdi/elearning-microservice/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateClassroomRequest struct {
	Id          primitive.ObjectID
	Name        string
	Description string
}

func (d *UpdateClassroomRequest) FromProtoBuffer(req *pb.UpdateClassroomRequest) (err error) {
	d.Id, err = primitive.ObjectIDFromHex(req.Id)
	d.Name = req.Name
	d.Description = req.Description
	return
}

func (d *UpdateClassroomRequest) ToProtoBuffer() (res pb.UpdateClassroomRequest) {
	res.Name = d.Name
	res.Description = d.Description
	return
}
