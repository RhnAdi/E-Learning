package repository

import (
	"context"
	"fmt"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassroomRepository interface {
	CreateClassroom(req *models.Classroom) (*models.Classroom, error)
	GetClassroomById(id string) (*models.Classroom, error)
	GetClassroomByName(name string) (*models.Classroom, error)
	GetAllClassroom() ([]*models.Classroom, error)
	UpdateClassroom(req *models.Classroom) (*models.Classroom, error)
	DeleteClassroom(req *models.Classroom) (*models.Classroom, error)
	JoinClass(req *models.StudentClass) (*models.StudentClass, error)
	AddStudent(req *models.StudentClass) (*models.StudentClass, error)
	DeleteStudentInClass(req *models.StudentClass) (*models.StudentClass, error)
	Students(classId string) (students []*primitive.ObjectID, err error)
	JoinRequest(classId string) (students []*models.StudentClass, err error)
	UpdateJoinRequest(req *models.StudentClass) (students *models.StudentClass, err error)
	DeleteJoinRequest(req *models.StudentClass) (students *models.StudentClass, err error)
}

type classroomRepository struct {
	c  *mongo.Collection
	sc *mongo.Collection
}

const ClassroomCollection = "classroom"
const StudentClass = "student_class"

func NewClassroomRepository(conn db.Connection) ClassroomRepository {
	return &classroomRepository{
		c:  conn.DB().Collection(ClassroomCollection),
		sc: conn.DB().Collection(StudentClass),
	}
}

func (r *classroomRepository) CreateClassroom(req *models.Classroom) (*models.Classroom, error) {
	res, err := r.c.InsertOne(context.Background(), req)
	req.Id = res.InsertedID.(primitive.ObjectID)
	return req, err
}

func (r *classroomRepository) GetClassroomById(id string) (class *models.Classroom, err error) {
	classId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	err = r.c.FindOne(context.Background(), bson.M{"_id": classId}).Decode(&class)
	return
}

func (r *classroomRepository) GetClassroomByName(name string) (class *models.Classroom, err error) {
	err = r.c.FindOne(context.Background(), bson.M{"name": name}).Decode(&class)
	return
}

func (r *classroomRepository) GetAllClassroom() (class []*models.Classroom, err error) {
	cur, err := r.c.Find(context.Background(), bson.M{})
	for cur.Next(context.Background()) {
		var temp_class models.Classroom
		err = cur.Decode(&temp_class)
		class = append(class, &temp_class)
	}
	return
}

func (r *classroomRepository) UpdateClassroom(req *models.Classroom) (*models.Classroom, error) {
	_, err := r.c.UpdateByID(context.Background(), req.Id, bson.M{"$set": bson.M{
		"name":        req.Name,
		"description": req.Description,
	}})
	return req, err
}

func (r *classroomRepository) DeleteClassroom(req *models.Classroom) (res *models.Classroom, err error) {
	err = r.c.FindOneAndDelete(context.Background(), bson.M{"_id": req.Id}).Decode(&res)
	return
}

func (r *classroomRepository) JoinClass(req *models.StudentClass) (*models.StudentClass, error) {
	insertRes, err := r.sc.InsertOne(context.Background(), req)
	req.Id = insertRes.InsertedID.(primitive.ObjectID)
	return req, err
}

func (r *classroomRepository) AddStudent(req *models.StudentClass) (*models.StudentClass, error) {
	inserted, err := r.sc.InsertOne(context.Background(), req)
	req.Id = inserted.InsertedID.(primitive.ObjectID)
	return req, err
}

func (r *classroomRepository) Students(classId string) (students []*primitive.ObjectID, err error) {
	class_id, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return
	}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "classroom_id", Value: class_id}, {Key: "status", Value: true}}}},
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$student_id"}}}},
	}

	cur, err := r.sc.Aggregate(context.Background(), pipeline)
	if err != nil {
		return students, err
	}
	type Temp struct {
		Id primitive.ObjectID `bson:"_id"`
	}
	for cur.Next(context.Background()) {
		var temp Temp
		err = cur.Decode(&temp)
		students = append(students, &temp.Id)
	}
	return
}

func (r *classroomRepository) JoinRequest(classId string) (students []*models.StudentClass, err error) {
	class_id, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return
	}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "classroom_id", Value: class_id}, {Key: "status", Value: false}}}},
		// {{Key: "$group", Value: bson.D{{Key: "_id", Value: "$student_id"}}}},
	}

	cur, err := r.sc.Aggregate(context.Background(), pipeline)
	if err != nil {
		return students, err
	}

	for cur.Next(context.Background()) {
		var temp models.StudentClass
		err = cur.Decode(&temp)
		students = append(students, &temp)
	}
	return
}

func (r *classroomRepository) UpdateJoinRequest(req *models.StudentClass) (students *models.StudentClass, err error) {
	res := r.sc.FindOneAndUpdate(context.Background(), bson.M{"_id": req.Id}, bson.M{"$set": req})
	return req, res.Err()
}

func (r *classroomRepository) DeleteJoinRequest(req *models.StudentClass) (students *models.StudentClass, err error) {
	res := r.sc.FindOneAndDelete(context.Background(), bson.M{"_id": req.Id})
	return req, res.Err()
}

func (r *classroomRepository) DeleteStudentInClass(req *models.StudentClass) (*models.StudentClass, error) {
	err := r.sc.FindOneAndDelete(context.Background(), req).Decode(&req)
	fmt.Println(err)
	return req, err
}
