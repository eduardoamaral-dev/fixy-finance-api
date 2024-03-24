package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type MongoRepository struct {
	client   *mongo.Client
	database *mongo.Database
}

func InitializeMongoDbConnection() (*MongoRepository, error) {
	return NewMongoDBRepository(os.Getenv("MONGO_CONNECTION_STRING"), "fixy")
}

func NewMongoDBRepository(connectionString, dbName string) (*MongoRepository, error) {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	database := client.Database(dbName)

	return &MongoRepository{
		client:   client,
		database: database,
	}, nil
}

func (r *MongoRepository) GetClient() *mongo.Client {
	return r.client
}

func (r *MongoRepository) GetDatabase() *mongo.Database {
	return r.database
}

func (r *MongoRepository) Close() error {
	return r.client.Disconnect(context.Background())
}
