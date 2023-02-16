package usecase

import (
	"errors"
	"time"

	"github.com/RhnAdi/elearning-microservice/services/Classroom/dto"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/models"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClassroomUsecase interface {
	CreateClassroom(req *dto.CreateClassroomRequest) (*dto.ClassroomResponse, error)
	GetClassroomById(id string) (*dto.ClassroomResponse, error)
	GetClassroomByName(id string) (*dto.ClassroomResponse, error)
	UpdateClassroom(id string, req *dto.UpdateClassroomRequest) (*dto.ClassroomResponse, error)
	DeleteClassroom(id string) (*dto.ClassroomResponse, error)
	JoinClass(req *dto.JoinClassRequest) (*dto.StudentClassResponse, error)
}

type classroomUsecase struct {
	classroomRepo repository.ClassroomRepository
}

func NewClassroomUsecase(classroomRepo repository.ClassroomRepository) ClassroomUsecase {
	return &classroomUsecase{
		classroomRepo: classroomRepo,
	}
}

func (u *classroomUsecase) CreateClassroom(req *dto.CreateClassroomRequest) (*dto.ClassroomResponse, error) {
	res := new(dto.ClassroomResponse)
	createRes, err := u.classroomRepo.CreateClassroom(&models.Classroom{
		Id:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		TeacherId:   req.TeacherId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return res, err
	}

	res.FromModelClassroom(createRes)
	return res, nil
}

func (u *classroomUsecase) GetClassroomById(id string) (*dto.ClassroomResponse, error) {
	res := new(dto.ClassroomResponse)
	result, err := u.classroomRepo.GetClassroomById(id)
	res.FromModelClassroom(result)
	return res, err
}

func (u *classroomUsecase) GetClassroomByName(name string) (*dto.ClassroomResponse, error) {
	res := new(dto.ClassroomResponse)
	result, err := u.classroomRepo.GetClassroomByName(name)
	res.FromModelClassroom(result)
	return res, err
}

func (u *classroomUsecase) UpdateClassroom(id string, req *dto.UpdateClassroomRequest) (*dto.ClassroomResponse, error) {
	res := new(dto.ClassroomResponse)
	classId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}
	result, err := u.classroomRepo.UpdateClassroom(&models.Classroom{
		Id:          classId,
		Name:        req.Name,
		Description: req.Description,
	})

	res.FromModelClassroom(result)
	return res, err
}

func (u *classroomUsecase) DeleteClassroom(id string) (*dto.ClassroomResponse, error) {
	res := new(dto.ClassroomResponse)
	classId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}

	result, err := u.classroomRepo.DeleteClassroom(&models.Classroom{Id: classId})
	res.FromModelClassroom(result)
	return res, err
}

func (u *classroomUsecase) JoinClass(req *dto.JoinClassRequest) (*dto.StudentClassResponse, error) {
	res := new(dto.StudentClassResponse)
	_, err := u.classroomRepo.GetClassroomById(req.ClassroomId.Hex())
	if err != nil {
		return res, errors.New("classroom not found")
	}

	result, err := u.classroomRepo.JoinClass(&models.StudentClass{
		Id:          primitive.NewObjectID(),
		ClassroomId: req.ClassroomId,
		StudentId:   req.StudentId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      false,
	})

	res.FromStudentClassModel(result)
	return res, err
}
