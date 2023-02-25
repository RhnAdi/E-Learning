package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/RhnAdi/elearning-microservice/services/Classroom/dto"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/models"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/repository"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/validators"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClassroomUsecase interface {
	CreateClassroom(req *dto.CreateClassroomRequest) (*dto.ClassroomResponse, error)
	GetClassroomById(id string) (*dto.ClassroomResponse, error)
	GetClassroomByName(id string) (*dto.ClassroomResponse, error)
	UpdateClassroom(req *dto.UpdateClassroomRequest) (*dto.ClassroomResponse, error)
	DeleteClassroom(id string) (*dto.ClassroomResponse, error)
	GetAllClassroom() (res *dto.ListClassroomResponse, err error)
	JoinClass(req *dto.JoinClassRequest) (*dto.StudentClassResponse, error)
	GetAllJoinRequests(id string) (res *dto.JoinRequestResponse, err error)
	AddStudents(req *dto.AddStudentRequest) (*dto.StudentClassResponse, error)
	GetStudent(id string) (*dto.StudentClassResponse, error)
	GetAllStudents(classroomId string) (*dto.StudentsResponse, error)
	AcceptJoinRequest(id string, teacherId string) (res *dto.StudentClassResponse, err error)
	RejectJoinRequest(id string, teacherId string) (res *dto.StudentClassResponse, err error)
	MyClass(id, role string) (*dto.ListClassroomResponse, error)
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

	if err := validators.ValidateCreateClassroom(req); err != nil {
		return res, err
	}

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
	if id == "" {
		return res, validators.ErrInvalidId
	}
	result, err := u.classroomRepo.GetClassroomById(id)
	res.FromModelClassroom(&result)
	return res, err
}
func (u *classroomUsecase) GetClassroomByName(name string) (*dto.ClassroomResponse, error) {
	res := new(dto.ClassroomResponse)
	result, err := u.classroomRepo.GetClassroomByName(name)
	res.FromModelClassroom(result)
	return res, err
}
func (u *classroomUsecase) GetAllClassroom() (res *dto.ListClassroomResponse, err error) {
	result, err := u.classroomRepo.GetAllClassroom()
	for _, class := range result {
		tmp := new(dto.ClassroomResponse)
		tmp.FromModelClassroom(class)
		res.AllClassroom = append(res.AllClassroom, tmp)
	}
	return
}
func (u *classroomUsecase) UpdateClassroom(req *dto.UpdateClassroomRequest) (*dto.ClassroomResponse, error) {
	res := new(dto.ClassroomResponse)
	if err := validators.ValidateUpdateClassroom(req); err != nil {
		return res, err
	}
	result, err := u.classroomRepo.UpdateClassroom(&models.Classroom{
		Id:          req.Id,
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
	if err := validators.ValidateJoinClass(req); err != nil {
		return res, err
	}
	_, err := u.classroomRepo.GetClassroomById(req.ClassroomId.Hex())
	if err != nil {
		return res, validators.ErrClassroomNotFound
	}

	_, err = u.classroomRepo.StudentStatus(&models.StudentClass{
		ClassroomId: req.ClassroomId,
		StudentId:   req.StudentId,
	})
	if err == nil {
		return res, validators.ErrAlreadyJoinClass
	}

	result, err := u.classroomRepo.AddStudent(&models.StudentClass{
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
func (u *classroomUsecase) AddStudents(req *dto.AddStudentRequest) (*dto.StudentClassResponse, error) {
	res := new(dto.StudentClassResponse)
	if err := validators.ValidateAddStudentRequest(req); err != nil {
		return res, err
	}
	_, err := u.classroomRepo.GetClassroomById(req.ClassroomId.Hex())
	if err != nil {
		return res, validators.ErrClassroomNotFound
	}

	result, err := u.classroomRepo.AddStudent(&models.StudentClass{
		Id:          primitive.NewObjectID(),
		ClassroomId: req.ClassroomId,
		StudentId:   req.StudentId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      true,
	})

	res.FromStudentClassModel(result)
	return res, err
}
func (u *classroomUsecase) GetStudent(id string) (res *dto.StudentClassResponse, err error) {
	if id == "" {
		err = validators.ErrEmtyId
		return
	}
	student_class_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}

	result, err := u.classroomRepo.GetStudent(&models.StudentClass{Id: student_class_id})
	res.FromStudentClassModel(result)

	return res, err
}
func (u *classroomUsecase) GetAllStudents(id string) (res *dto.StudentsResponse, err error) {
	if id == "" {
		err = validators.ErrEmtyId
		return
	}
	res.Students, err = u.classroomRepo.GetStudents(id)
	return
}
func (u *classroomUsecase) GetAllJoinRequests(id string) (res *dto.JoinRequestResponse, err error) {
	result, err := u.classroomRepo.GetJoinRequests(id)
	for _, student := range result {
		tmp := new(dto.StudentClassResponse)
		tmp.FromStudentClassModel(student)
		res.Requests = append(res.Requests, tmp)
	}
	return
}
func (u *classroomUsecase) AcceptJoinRequest(id string, teacherId string) (*dto.StudentClassResponse, error) {
	res := new(dto.StudentClassResponse)
	if ok := validators.ValidateJoinRequest(id, teacherId); !ok {
		return res, errors.New("request invalid")
	}
	studentClassId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}
	studentClass, err := u.classroomRepo.GetStudent(&models.StudentClass{
		Id: studentClassId,
	})
	if err != nil {
		return res, fmt.Errorf("id : %v : %v", studentClassId, err)
	}
	studentClass.Status = true
	studentClass.UpdatedAt = time.Now()
	classroomInfo, err := u.classroomRepo.GetClassroomById(studentClass.ClassroomId.Hex())
	if err != nil {
		return res, validators.ErrClassroomNotFound
	}
	if classroomInfo.TeacherId.Hex() != teacherId {
		return res, errors.New("you are not teacher")
	}
	result, err := u.classroomRepo.UpdateStudent(studentClass)
	res.FromStudentClassModel(result)
	return res, err
}
func (u *classroomUsecase) RejectJoinRequest(id string, teacherId string) (*dto.StudentClassResponse, error) {
	res := new(dto.StudentClassResponse)
	if ok := validators.ValidateJoinRequest(id, teacherId); !ok {
		return res, errors.New("request invalid")
	}
	studentClassId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}
	studentClass, err := u.classroomRepo.GetStudent(&models.StudentClass{
		Id: studentClassId,
	})
	if err != nil {
		return res, fmt.Errorf("request not found")
	}
	studentClass.Status = false
	studentClass.UpdatedAt = time.Now()
	classroomInfo, err := u.classroomRepo.GetClassroomById(studentClass.ClassroomId.Hex())
	if err != nil {
		return res, err
	}
	if classroomInfo.TeacherId.Hex() != teacherId {
		return res, errors.New("you are not teacher")
	}
	result, err := u.classroomRepo.DeleteStudent(studentClass)
	fmt.Println(result)
	if err != nil {
		return res, err
	}
	res.FromStudentClassModel(result)
	return res, nil
}
func (u *classroomUsecase) MyClass(id, role string) (*dto.ListClassroomResponse, error) {
	res := new(dto.ListClassroomResponse)
	if role == "student" {
		result, err := u.classroomRepo.GetAllClassroomByStudentId(id)
		if err != nil {
			return res, err
		}
		for _, class := range result {
			tmp := new(dto.ClassroomResponse)
			tmp.FromModelClassroom(class)
			res.AllClassroom = append(res.AllClassroom, tmp)
		}
	}
	if role == "teacher" {
		result, err := u.classroomRepo.GetAllClassroomByTeacherId(id)
		if err != nil {
			return res, err
		}
		for _, class := range result {
			tmp := new(dto.ClassroomResponse)
			tmp.FromModelClassroom(class)
			res.AllClassroom = append(res.AllClassroom, tmp)
		}
	}
	return res, nil
}
