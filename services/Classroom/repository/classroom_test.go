package repository_test

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/models"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("cant load env var : ", err.Error())
	}
}

func TestCreateClassroom(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	newClass := models.Classroom{
		Name:        "Metematika Diskrit 1",
		TeacherId:   primitive.NewObjectID(),
		Description: "Dosen Pak Andy Hepy, Semester 1",
	}
	newClass.Pre()

	res, err := classroomRepo.CreateClassroom(&newClass)
	assert.NoError(t, err)
	assert.Equal(t, newClass.Id.Hex(), res.Id.Hex())
	assert.Equal(t, newClass.Name, res.Name)
	assert.Equal(t, newClass.TeacherId.Hex(), res.TeacherId.Hex())
	assert.Equal(t, newClass.Description, res.Description)
}
func TestGetClassroomById(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	newClass := models.Classroom{
		Name:        "Struktur data & algoritma",
		TeacherId:   primitive.NewObjectID(),
		Description: "Dosen Pak Rehmawan Bagus, Semester 3",
	}
	newClass.Pre()

	_, err = classroomRepo.CreateClassroom(&newClass)
	assert.NoError(t, err)

	res, err := classroomRepo.GetClassroomById(newClass.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, newClass.Id.Hex(), res.Id.Hex())
	assert.Equal(t, newClass.Name, res.Name)
	assert.Equal(t, newClass.TeacherId.Hex(), res.TeacherId.Hex())
	assert.Equal(t, newClass.Description, res.Description)
}
func TestGetClassroomByName(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	newClass := models.Classroom{
		Name:        "Basis Data",
		TeacherId:   primitive.NewObjectID(),
		Description: "Dosen Pak Eko Supriyadi, Semester 3",
	}
	newClass.Pre()

	_, err = classroomRepo.CreateClassroom(&newClass)
	assert.NoError(t, err)

	res, err := classroomRepo.GetClassroomByName(newClass.Name)
	assert.NoError(t, err)
	assert.Equal(t, newClass.Id.Hex(), res.Id.Hex())
	assert.Equal(t, newClass.Name, res.Name)
	assert.Equal(t, newClass.TeacherId.Hex(), res.TeacherId.Hex())
	assert.Equal(t, newClass.Description, res.Description)
}
func TestGetAllClassroom(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	_, err = classroomRepo.GetAllClassroom()
	assert.NoError(t, err)
}
func TestUpdateClassroom(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	class := models.Classroom{
		Name:        "Logic Math",
		TeacherId:   primitive.NewObjectID(),
		Description: "Dosen Pak Andy Hepi, Semester 1",
	}
	class.Pre()

	_, err = classroomRepo.CreateClassroom(&class)
	assert.NoError(t, err)

	class.Name = "Logika Matematika"
	class.Description = "Dosen Pak Andi Hepi, Semester 1"

	res, err := classroomRepo.UpdateClassroom(&class)
	assert.NoError(t, err)
	assert.Equal(t, class.Id.Hex(), res.Id.Hex())
	assert.Equal(t, class.Name, res.Name)
	assert.Equal(t, class.TeacherId.Hex(), res.TeacherId.Hex())
	assert.Equal(t, class.Description, res.Description)
}
func TestDeleteClassroom(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	newClass := models.Classroom{
		Name:        "Pendidika Pancasila",
		TeacherId:   primitive.NewObjectID(),
		Description: "Dosen Pak Kusnan, Mata Kuliah Umum",
	}
	newClass.Pre()

	_, err = classroomRepo.CreateClassroom(&newClass)
	assert.NoError(t, err)

	res, err := classroomRepo.DeleteClassroom(&newClass)
	assert.NoError(t, err)
	assert.Equal(t, newClass.Id.Hex(), res.Id.Hex())
	assert.Equal(t, newClass.Name, res.Name)
	assert.Equal(t, newClass.TeacherId.Hex(), res.TeacherId.Hex())
	assert.Equal(t, newClass.Description, res.Description)
}
func TestDeleteStudentInClass(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	reqJoin := models.StudentClass{
		ClassroomId: primitive.NewObjectID(),
		StudentId:   primitive.NewObjectID(),
	}
	reqJoin.Pre()

	res, err := classroomRepo.AddStudent(&reqJoin)
	assert.NoError(t, err)
	assert.Equal(t, reqJoin.Id.Hex(), res.Id.Hex())
	assert.Equal(t, reqJoin.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, reqJoin.StudentId.Hex(), res.StudentId.Hex())
	assert.Equal(t, reqJoin.Status, res.Status)

	delRes, err := classroomRepo.DeleteStudent(res)
	assert.NoError(t, err)
	assert.Equal(t, delRes.Id.Hex(), res.Id.Hex())
	assert.Equal(t, delRes.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, delRes.StudentId.Hex(), res.StudentId.Hex())
	assert.Equal(t, delRes.Status, res.Status)
}
func TestAddStudent(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	studentClass := models.StudentClass{
		ClassroomId: primitive.NewObjectID(),
		StudentId:   primitive.NewObjectID(),
	}
	studentClass.Pre()

	res, err := classroomRepo.AddStudent(&studentClass)
	assert.NoError(t, err)
	assert.Equal(t, studentClass.Id.Hex(), res.Id.Hex())
	assert.Equal(t, studentClass.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, studentClass.StudentId.Hex(), res.StudentId.Hex())
}
func TestStudents(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	classroom_id := primitive.NewObjectID()

	studentAddeds := []models.StudentClass{}

	for i := 0; i < 5; i++ {
		studentClass := models.StudentClass{
			ClassroomId: classroom_id,
			StudentId:   primitive.NewObjectID(),
			Status:      true,
		}
		studentClass.Pre()

		add_res, err := classroomRepo.AddStudent(&studentClass)
		assert.NoError(t, err)
		assert.Equal(t, studentClass.Id.Hex(), add_res.Id.Hex())
		assert.Equal(t, studentClass.ClassroomId.Hex(), add_res.ClassroomId.Hex())
		assert.Equal(t, studentClass.StudentId.Hex(), add_res.StudentId.Hex())
		studentAddeds = append(studentAddeds, studentClass)
	}

	res, err := classroomRepo.GetStudents(classroom_id.Hex())
	assert.NoError(t, err)
	for _, sc := range res {
		for _, sa := range studentAddeds {
			if sc.Hex() == sa.StudentId.Hex() {
				assert.Equal(t, sc.Hex(), sa.StudentId.Hex())
				continue
			}
		}
		assert.Error(t, errors.New("not matching"))
	}
}
func TestJoinRequest(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	classroom_id := primitive.NewObjectID()

	studentAddeds := []models.StudentClass{}

	for i := 0; i < 5; i++ {
		studentClass := models.StudentClass{
			ClassroomId: classroom_id,
			StudentId:   primitive.NewObjectID(),
			Status:      false,
		}
		studentClass.Pre()

		add_res, err := classroomRepo.AddStudent(&studentClass)
		assert.NoError(t, err)
		assert.Equal(t, studentClass.Id.Hex(), add_res.Id.Hex())
		assert.Equal(t, studentClass.ClassroomId.Hex(), add_res.ClassroomId.Hex())
		assert.Equal(t, studentClass.StudentId.Hex(), add_res.StudentId.Hex())
		assert.Equal(t, studentClass.Status, add_res.Status)
		studentAddeds = append(studentAddeds, studentClass)
	}

	res, err := classroomRepo.GetJoinRequests(classroom_id.Hex())
	assert.NoError(t, err)
	for _, sc := range res {
		for _, sa := range studentAddeds {
			if sc.Id.Hex() == sa.StudentId.Hex() {
				assert.Equal(t, sc.Id.Hex(), sa.StudentId.Hex())
				continue
			}
		}
		assert.Error(t, errors.New("not matching"))
	}
}
func TestGetJoinRequest(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	classroom_id := primitive.NewObjectID()

	studentClass := models.StudentClass{
		ClassroomId: classroom_id,
		StudentId:   primitive.NewObjectID(),
		Status:      false,
	}
	studentClass.Pre()

	add_res, err := classroomRepo.AddStudent(&studentClass)
	assert.NoError(t, err)
	assert.Equal(t, studentClass.Id.Hex(), add_res.Id.Hex())
	assert.Equal(t, studentClass.ClassroomId.Hex(), add_res.ClassroomId.Hex())
	assert.Equal(t, studentClass.StudentId.Hex(), add_res.StudentId.Hex())
	assert.Equal(t, studentClass.Status, add_res.Status)
}
func TestUpdateJoinRequest(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	reqJoin := models.StudentClass{
		ClassroomId: primitive.NewObjectID(),
		StudentId:   primitive.NewObjectID(),
	}
	reqJoin.Pre()

	res, err := classroomRepo.AddStudent(&reqJoin)
	assert.NoError(t, err)
	assert.Equal(t, reqJoin.Id.Hex(), res.Id.Hex())
	assert.Equal(t, reqJoin.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, reqJoin.StudentId.Hex(), res.StudentId.Hex())
	assert.Equal(t, reqJoin.Status, res.Status)

	reqJoin.Status = true
	updateRes, err := classroomRepo.UpdateStudent(&reqJoin)
	assert.NoError(t, err)
	assert.Equal(t, updateRes.Id.Hex(), res.Id.Hex())
	assert.Equal(t, updateRes.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, updateRes.StudentId.Hex(), res.StudentId.Hex())
	assert.Equal(t, updateRes.Status, res.Status)
}
func TestDeleteJoinRequest(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	reqJoin := models.StudentClass{
		ClassroomId: primitive.NewObjectID(),
		StudentId:   primitive.NewObjectID(),
	}
	reqJoin.Pre()

	res, err := classroomRepo.AddStudent(&reqJoin)
	assert.NoError(t, err)
	assert.Equal(t, reqJoin.Id.Hex(), res.Id.Hex())
	assert.Equal(t, reqJoin.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, reqJoin.StudentId.Hex(), res.StudentId.Hex())
	assert.Equal(t, reqJoin.Status, res.Status)

	updateRes, err := classroomRepo.DeleteStudent(res)
	assert.NoError(t, err)
	assert.Equal(t, updateRes.Id.Hex(), res.Id.Hex())
	assert.Equal(t, updateRes.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, updateRes.StudentId.Hex(), res.StudentId.Hex())
	assert.Equal(t, updateRes.Status, res.Status)
}
func TestGetStudent(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)

	reqJoin := models.StudentClass{
		ClassroomId: primitive.NewObjectID(),
		StudentId:   primitive.NewObjectID(),
	}
	reqJoin.Pre()

	res, err := classroomRepo.AddStudent(&reqJoin)
	assert.NoError(t, err)
	assert.Equal(t, reqJoin.Id.Hex(), res.Id.Hex())
	assert.Equal(t, reqJoin.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, reqJoin.StudentId.Hex(), res.StudentId.Hex())
	assert.Equal(t, reqJoin.Status, res.Status)

	res, err = classroomRepo.GetStudent(res)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, reqJoin.Status)
}
func TestGetClassroomByTeacherId(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)
	teacherId := primitive.NewObjectID()

	listClass := []*primitive.ObjectID{}
	for i := 0; i < 5; i++ {
		class := models.Classroom{
			Id:          primitive.NewObjectID(),
			Name:        fmt.Sprintf("Test Class Get Teacher Classroom %d", i),
			Description: "alsnaklsa aklsdjasjd",
			TeacherId:   teacherId,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		res, err := classroomRepo.CreateClassroom(&class)
		assert.NoError(t, err)
		listClass = append(listClass, &res.Id)
	}

	myClass, err := classroomRepo.GetAllClassroomByTeacherId(teacherId.Hex())
	assert.NoError(t, err)

	for _, class := range myClass {
		for _, listId := range listClass {
			if class.Id.Hex() == listId.Hex() {
				assert.Equal(t, class.Id.Hex(), listId.Hex())
				continue
			}
		}
		assert.Error(t, errors.New("not match list id"))
	}
}
func TestGetClassroomByStudentId(t *testing.T) {
	dbcfg := db.NewConfig()
	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepo := repository.NewClassroomRepository(conn)
	studentId := primitive.NewObjectID()

	listClass := []*primitive.ObjectID{}
	for i := 0; i < 5; i++ {
		join_class := models.StudentClass{
			Id:          primitive.NewObjectID(),
			ClassroomId: primitive.NewObjectID(),
			StudentId:   studentId,
			Status:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		res, err := classroomRepo.AddStudent(&join_class)
		assert.NoError(t, err)
		listClass = append(listClass, &res.ClassroomId)
	}

	myClass, err := classroomRepo.GetAllClassroomByTeacherId(studentId.Hex())
	assert.NoError(t, err)

	for _, class := range myClass {
		for _, listId := range listClass {
			if class.Id.Hex() == listId.Hex() {
				assert.Equal(t, class.Id.Hex(), listId.Hex())
				continue
			}
		}
		assert.Error(t, errors.New("not match list id"))
	}
}
