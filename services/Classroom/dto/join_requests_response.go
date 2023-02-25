package dto

import "github.com/RhnAdi/elearning-microservice/pb"

type JoinRequestResponse struct {
	Requests []*StudentClassResponse
}

func (d *JoinRequestResponse) FromProtoBuffer(pb *pb.JoinClassRequestsResponse) {
	for _, temp_pb := range pb.Requests {
		temp := new(StudentClassResponse)
		temp.FromProtoBuffer(temp_pb)
		d.Requests = append(d.Requests, temp)
	}
}

func (d *JoinRequestResponse) ToProtoBuffer() *pb.JoinClassRequestsResponse {
	pb := new(pb.JoinClassRequestsResponse)
	for _, temp_d := range d.Requests {
		temp := temp_d.ToProtoBuffer()
		pb.Requests = append(pb.Requests, temp)
	}
	return pb
}
