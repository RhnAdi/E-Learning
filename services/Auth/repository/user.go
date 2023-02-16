package repository

import (
	"context"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/services/Auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(user *models.User) (*models.User, error)
	GetById(id string) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	GetAll() (users []models.User, err error)
	Update(user models.User) (models.User, error)
	Delete(user models.User) (models.User, error)
}

type userRepository struct {
	c *mongo.Collection
}

const UserCollection = "users"

func NewUserRepository(conn db.Connection) UserRepository {
	return &userRepository{c: conn.DB().Collection(UserCollection)}
}

func (r *userRepository) Save(user *models.User) (*models.User, error) {
	res, err := r.c.InsertOne(context.Background(), user)
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, err
}

func (r *userRepository) GetById(id string) (user models.User, err error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	err = r.c.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&user)
	return
}

func (r *userRepository) GetByEmail(email string) (user models.User, err error) {
	err = r.c.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return
}

func (r *userRepository) GetAll() (users []models.User, err error) {
	cur, err := r.c.Find(context.Background(), bson.M{})
	for cur.Next(context.Background()) {
		var user models.User
		err = cur.Decode(&user)
		users = append(users, user)
	}
	return
}

func (r *userRepository) Update(user models.User) (models.User, error) {
	_, err := r.c.UpdateOne(context.Background(), bson.M{"_id": user.Id}, bson.M{"$set": user})
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) Delete(user models.User) (models.User, error) {
	_, err := r.c.DeleteOne(context.Background(), user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
