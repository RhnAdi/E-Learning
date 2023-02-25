package rpc

import (
	"context"

	"github.com/RhnAdi/elearning-microservice/config/jwt"
	"github.com/RhnAdi/elearning-microservice/pb"
	"github.com/RhnAdi/elearning-microservice/pkg/interceptor"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/dto"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type classroomRpc struct {
	classroomUsecase usecase.ClassroomUsecase
	pb.UnimplementedClassroomServiceServer
}

func NewClassroomRpc(classroomUsecase usecase.ClassroomUsecase) pb.ClassroomServiceServer {
	return &classroomRpc{
		classroomUsecase: classroomUsecase,
	}
}

func (d *classroomRpc) CreateClassroom(ctx context.Context, req *pb.CreateClassroomRequest) (*pb.ClassroomResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	request := new(dto.CreateClassroomRequest)
	request.FromProtoBuffer(req)
	teacher_id, err := primitive.ObjectIDFromHex(account_info.Id)
	request.TeacherId = teacher_id
	res, err := d.classroomUsecase.CreateClassroom(request)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) GetClassroomById(ctx context.Context, req *pb.Id) (*pb.ClassroomResponse, error) {
	res, err := d.classroomUsecase.GetClassroomById(req.Id)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) GetClassroomByName(ctx context.Context, req *pb.Name) (*pb.ClassroomResponse, error) {
	res, err := d.classroomUsecase.GetClassroomById(req.Name)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) UpdateClassroom(ctx context.Context, req *pb.UpdateClassroomRequest) (*pb.ClassroomResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	request := new(dto.UpdateClassroomRequest)
	request.FromProtoBuffer(req)
	if err := d.isTeacherInClass(account_info.Id); err != nil {
		return &pb.ClassroomResponse{}, err
	}
	res, err := d.classroomUsecase.UpdateClassroom(request)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) DeleteClassroom(ctx context.Context, req *pb.Id) (*pb.ClassroomResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	if err := d.isTeacherInClass(account_info.Id); err != nil {
		return &pb.ClassroomResponse{}, err
	}
	res, err := d.classroomUsecase.DeleteClassroom(req.Id)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) GetAllClassroom(ctx context.Context, req *pb.Empty) (*pb.ListClassroomResponse, error) {
	res, err := d.classroomUsecase.GetAllClassroom()
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) JoinClass(ctx context.Context, req *pb.Id) (*pb.StudentClassResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	request := new(dto.JoinClassRequest)
	classroom_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return &pb.StudentClassResponse{}, err
	}
	student_id, err := primitive.ObjectIDFromHex(account_info.Id)
	request.ClassroomId = classroom_id
	request.StudentId = student_id
	res, err := d.classroomUsecase.JoinClass(request)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) GetAllJoinRequest(ctx context.Context, req *pb.Id) (*pb.JoinClassRequestsResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	if err := d.isTeacherInClass(account_info.Id); err != nil {
		return &pb.JoinClassRequestsResponse{}, err
	}
	res, err := d.classroomUsecase.GetAllJoinRequests(req.Id)
	if err != nil {
		return &pb.JoinClassRequestsResponse{}, err
	}
	return res.ToProtoBuffer(), nil
}
func (d *classroomRpc) AcceptJoinRequest(ctx context.Context, req *pb.Id) (*pb.StudentClassResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	if err := d.isTeacherInClass(account_info.Id); err != nil {
		return &pb.StudentClassResponse{}, err
	}
	res, err := d.classroomUsecase.AcceptJoinRequest(req.Id, account_info.Id)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) RejectJoinRequest(ctx context.Context, req *pb.Id) (*pb.StudentClassResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	if err := d.isTeacherInClass(account_info.Id); err != nil {
		return &pb.StudentClassResponse{}, err
	}
	res, err := d.classroomUsecase.RejectJoinRequest(req.Id, account_info.Id)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) GetStudentInfo(ctx context.Context, req *pb.Id) (*pb.StudentClassResponse, error) {
	res, err := d.classroomUsecase.GetStudent(req.Id)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) MyClass(ctx context.Context, req *pb.Empty) (*pb.ListClassroomResponse, error) {
	jwtclaims := ctx.Value(interceptor.CtxKey("claim"))
	account_info := jwtclaims.(jwt.JWTClaims)
	res, err := d.classroomUsecase.MyClass(account_info.Id, account_info.Role)
	return res.ToProtoBuffer(), err
}
func (d *classroomRpc) GetAllStudents(ctx context.Context, req *pb.Id) (*pb.StudentsResponse, error) {
	res, err := d.classroomUsecase.GetAllStudents(req.Id)
	return res.ToProtoBuffer(), err
}

func (d *classroomRpc) isTeacherInClass(id string) error {
	class, err := d.classroomUsecase.GetClassroomById(id)
	if err != nil {
		return err
	}
	if class.TeacherId.Hex() != id {
		return status.Errorf(codes.PermissionDenied, "you are not teacher")
	}
	return nil
}
