package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type ListClassroomResponse struct {
	AllClassroom []*ClassroomResponse
}

func (d *ListClassroomResponse) FromProtoBuffer(pb *pb.ListClassroomResponse) {
	for _, user := range pb.AllClassroom {
		data := new(ClassroomResponse)
		data.FromProtoBuffer(user)
		d.AllClassroom = append(d.AllClassroom, data)
	}
}

func (d *ListClassroomResponse) ToProtoBuffer() (pb pb.ListClassroomResponse) {
	for _, user := range d.AllClassroom {
		temp := user.ToProtoBuffer()
		pb.AllClassroom = append(pb.AllClassroom, &temp)
	}
	return
}
