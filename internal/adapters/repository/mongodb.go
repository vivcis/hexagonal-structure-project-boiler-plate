package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
	"vivicis/github.com/notification-service/internal/ports"
)

type mongoDB struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewUser(mongo *mongo.Client) ports.UserRepository {
	return &Db{Mongo: mongo}
}

// NewMongoDB creates an instance of mongoDB type
func NewMongoDB(client *mongo.Client, database string, timeout time.Duration) *mongoDB {
	return &mongoDB{
		client:   client,
		database: database,
		timeout:  timeout,
	}
}

// ConnectMongoDB connects the mongodb and returns Db type and assigns the mongodb client to the Db client field
func ConnectMongoDB(c *databaseConfig) (*Db, error) {
	logrus.Info("Connecting to Mongo DB ")
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxTimeout)
	defer cancel()
	mongoUri := "mongodb://localhost:27017/notification_service"
	//mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%v/%s", c.user, c.password, c.host, c.port, c.database)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	log.Println("here", mongoUri)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	mgo := NewMongoDB(client, c.database, c.ctxTimeout)
	return &Db{Mongo: mgo.client}, nil

}
