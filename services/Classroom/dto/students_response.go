package dto

import (
	"github.com/RhnAdi/elearning-microservice/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentsResponse struct {
	Students []*primitive.ObjectID
}

func (d *StudentsResponse) FromProtoBuffer(pb *pb.StudentsResponse) (err error) {
	for _, tmp_pb := range pb.Students {
		studentId, err := primitive.ObjectIDFromHex(tmp_pb)
		if err != nil {
			return err
		}
		d.Students = append(d.Students, &studentId)
	}
	return nil
}

func (d *StudentsResponse) ToProtoBuffer() (pb pb.StudentsResponse) {
	for _, tmp_d := range d.Students {
		pb.Students = append(pb.Students, tmp_d.Hex())
	}
	return
}
