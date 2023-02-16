package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type UpdateClassroomRequest struct {
	Name        string
	Description string
}

func (d *UpdateClassroomRequest) FromProtoBuffer(req *pb.UpdateClassroomRequest) {
	d.Name = req.Name
	d.Description = req.Description
}

func (d *UpdateClassroomRequest) ToProtoBuffer() (res pb.UpdateClassroomRequest) {
	res.Name = d.Name
	res.Description = d.Description
	return
}
