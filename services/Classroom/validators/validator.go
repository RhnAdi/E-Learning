package validators

import (
	"errors"

	"github.com/RhnAdi/elearning-microservice/services/Classroom/dto"
)

var (
	ErrInvalidId          = errors.New("invalid id")
	ErrEmtyId             = errors.New("id can't be empty")
	ErrInvalidClassroomId = errors.New("invalid classroom id")
	ErrInvalidStudentId   = errors.New("invalid student id")
	ErrInvalidTeacherId   = errors.New("invalid teacher id")
	ErrEmptyName          = errors.New("name can't be empty")
	ErrClassroomNotFound  = errors.New("classroom not found")
	ErrAlreadyJoinClass   = errors.New("you already joined class")
)

func ValidateCreateClassroom(req *dto.CreateClassroomRequest) (err error) {
	if req.Name == "" {
		return ErrEmptyName
	}

	if req.TeacherId.IsZero() {
		return ErrInvalidTeacherId
	}

	return nil
}

func ValidateUpdateClassroom(req *dto.UpdateClassroomRequest) (err error) {
	if req.Name == "" {
		return ErrEmptyName
	}

	return nil
}

func ValidateJoinClass(req *dto.JoinClassRequest) (err error) {
	if req.ClassroomId.IsZero() {
		return ErrInvalidClassroomId
	}

	if req.StudentId.IsZero() {
		return ErrInvalidStudentId
	}

	return nil
}

func ValidateAddStudentRequest(req *dto.AddStudentRequest) (err error) {
	if req.ClassroomId.IsZero() {
		return ErrInvalidClassroomId
	}

	if req.StudentId.IsZero() {
		return ErrInvalidStudentId
	}

	return nil
}

func ValidateJoinRequest(id, teacherId string) bool {
	if id == "" {
		return false
	}
	if teacherId == "" {
		return false
	}
	return true
}
