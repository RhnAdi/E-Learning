package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type connection struct {
	session  *mongo.Session
	database *mongo.Database
}

func NewConnection(cfg DBConfig) (Connection, error) {
	// log.Println("connect to database : ", cfg.Dsn())
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Dsn()))
	if err != nil {
		return &connection{}, err
	}
	session, err := client.StartSession()
	if err != nil {
		return &connection{}, err
	}
	return &connection{session: &session, database: client.Database(cfg.DbName())}, nil
}

func (c *connection) Close() {
	c.database.Client().Disconnect(context.Background())
}

func (c *connection) DB() *mongo.Database {
	return c.database
}
