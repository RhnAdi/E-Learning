package repository

import (
	"context"
	"errors"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassroomRepository interface {
	CreateClassroom(req *models.Classroom) (*models.Classroom, error)
	GetClassroomById(id string) (models.Classroom, error)
	GetClassroomByName(name string) (*models.Classroom, error)
	GetAllClassroom() ([]*models.Classroom, error)
	UpdateClassroom(req *models.Classroom) (*models.Classroom, error)
	DeleteClassroom(req *models.Classroom) (*models.Classroom, error)
	AddStudent(req *models.StudentClass) (*models.StudentClass, error)
	GetStudent(req *models.StudentClass) (*models.StudentClass, error)
	GetStudents(classId string) (students []*primitive.ObjectID, err error)
	UpdateStudent(req *models.StudentClass) (students *models.StudentClass, err error)
	DeleteStudent(req *models.StudentClass) (*models.StudentClass, error)
	GetJoinRequests(classId string) (students []*models.StudentClass, err error)
	StudentStatus(req *models.StudentClass) (bool, error)
	GetAllClassroomByTeacherId(string) ([]*models.Classroom, error)
	GetAllClassroomByStudentId(string) ([]*models.Classroom, error)
}

type classroomRepository struct {
	c  *mongo.Collection
	sc *mongo.Collection
}

const ClassroomCollection = "classroom"
const StudentClass = "student"

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

func (r *classroomRepository) GetClassroomById(id string) (class models.Classroom, err error) {
	classId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return class, err
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

func (r *classroomRepository) AddStudent(req *models.StudentClass) (*models.StudentClass, error) {
	inserted, err := r.sc.InsertOne(context.Background(), req)
	req.Id = inserted.InsertedID.(primitive.ObjectID)
	return req, err
}

func (r *classroomRepository) GetStudent(req *models.StudentClass) (*models.StudentClass, error) {
	res := r.sc.FindOne(context.Background(), bson.M{"_id": req.Id})
	if err := res.Err(); err != nil {
		return req, err
	}
	err := res.Decode(&req)
	return req, err
}

func (r *classroomRepository) GetStudents(classId string) (students []*primitive.ObjectID, err error) {
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

func (r *classroomRepository) GetJoinRequests(classId string) (students []*models.StudentClass, err error) {
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

func (r *classroomRepository) StudentStatus(req *models.StudentClass) (bool, error) {
	if req.StudentId.IsZero() && req.ClassroomId.IsZero() {
		return false, errors.New("student_id and classroom_id can't empty")
	}
	err := r.sc.FindOne(context.Background(), bson.M{"classroom_id": req.ClassroomId, "student_id": req.StudentId}).Decode(&req)
	return req.Status, err
}

func (r *classroomRepository) UpdateStudent(req *models.StudentClass) (students *models.StudentClass, err error) {
	res := r.sc.FindOneAndUpdate(context.Background(), bson.M{"_id": req.Id}, bson.M{"$set": req})
	return req, res.Err()
}

func (r *classroomRepository) DeleteStudent(req *models.StudentClass) (*models.StudentClass, error) {
	err := r.sc.FindOneAndDelete(context.Background(), bson.M{"_id": req.Id}).Decode(&req)
	return req, err
}

func (r *classroomRepository) GetAllClassroomByTeacherId(id string) (res []*models.Classroom, err error) {
	teacherId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	cursor, err := r.c.Find(context.Background(), bson.M{"teacher_id": teacherId})
	if err != nil {
		return
	}
	for cursor.Next(context.Background()) {
		tmp := new(models.Classroom)
		cursor.Decode(tmp)
		res = append(res, tmp)
	}
	return
}
func (r *classroomRepository) GetAllClassroomByStudentId(id string) ([]*models.Classroom, error) {
	res := []*models.Classroom{}
	studentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}
	pipe := mongo.Pipeline{
		{
			{
				Key: "$lookup", Value: bson.M{
					"from":         "classroom",
					"localField":   "classroom_id",
					"foreignField": "_id",
					"as":           "myclass",
				},
			},
		},
		{
			{
				Key: "$match", Value: bson.M{
					"student_id": studentId,
					"status":     true,
				},
			},
		},
		{
			{
				Key: "$project", Value: bson.M{
					"myclass": 1,
					"_id":     0,
				},
			},
		},
		{
			{
				Key: "$unwind", Value: "$myclass",
			},
		},
		{
			{
				Key: "$replaceRoot", Value: bson.M{
					"newRoot": "$myclass",
				},
			},
		},
	}
	cursor, err := r.sc.Aggregate(context.Background(), pipe)
	if err != nil {
		return res, err
	}
	for cursor.Next(context.Background()) {
		tmp := new(models.Classroom)
		err := cursor.Decode(tmp)
		if err != nil {
			return res, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
