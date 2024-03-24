package services

import (
	"context"
	"fixy-finance-api/external/mongodb"
	. "fixy-finance-api/internal/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ExpenseService struct{}

func (service ExpenseService) SaveNewExpense(newExpense NewExpenseRequest) (*mongo.InsertOneResult, error) {
	dbConnection, connectionError := mongodb.InitializeMongoDbConnection()

	if connectionError != nil {
		fmt.Println("MongoDB Connection Error", connectionError)
	}

	entry := newExpense
	entry.CreatedAt = time.Now()

	result, insertionError := dbConnection.GetDatabase().Collection("expenses").InsertOne(context.TODO(), entry)
	if insertionError != nil {
		fmt.Println("MongoDB Insertion Error:", insertionError)
	}

	return result, nil
}

func (service ExpenseService) GetAllExpenses() ([]bson.M, error) {
	dbConnection, connectionError := mongodb.InitializeMongoDbConnection()
	if connectionError != nil {
		fmt.Println("MongoDB Connection Error", connectionError)
	}
	filter := bson.M{}

	cursor, readingError := dbConnection.GetDatabase().Collection("expenses").Find(context.Background(), filter)
	if readingError != nil {
		println(readingError)
	}

	defer cursor.Close(context.Background())

	var results []bson.M
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
