package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_DB_URI = os.Getenv("MONGO_URL") // "mongodb://root:example@localhost:27017"
const MONGO_DB_NAME = "compliancedb"
const BLACKLISTED_USERS_COLLECTION = "blaclistedUsers"

type ComplianceRepository interface {
	initializeRepository()
	findByNameAndLastName(lastName string, name string) []SanctionedUser
	insert(user UserData) OperationResult
	findAll() []UserData
}

type ComplianceRepositoryImpl struct {
}

func (repository ComplianceRepositoryImpl) findByNameAndLastName(lastName string, name string) []SanctionedUser {
	var response []SanctionedUser
	client := connectToDb()
	if client != nil {
		collection := client.Collection(BLACKLISTED_USERS_COLLECTION)

		filter := bson.D{{"lastName", lastName}, {"name", name}}

		results, err := collection.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}

		for results.Next(context.TODO()) {
			var user SanctionedUser
			if err := results.Decode(&user); err != nil {
				fmt.Println(err)
			}
			response = append(response, user)

		}

		// display the id of the newly inserted object
	}
	return response
}

func (repository ComplianceRepositoryImpl) initializeRepository() {

	client := connectToDb()
	if client != nil {
		collection := client.Collection(BLACKLISTED_USERS_COLLECTION)

		user := bson.D{{"lastName", "Andrzej"}, {"name", "Nowak"}}
		result, err := collection.InsertOne(context.TODO(), user)

		if err != nil {
			panic(err)
		}
		// display the id of the newly inserted object
		fmt.Println(result.InsertedID)
	}
}

func (repository ComplianceRepositoryImpl) insert(user UserData) OperationResult {
	client := connectToDb()
	result := OperationResult{}
	if client != nil {
		collection := client.Collection(BLACKLISTED_USERS_COLLECTION)
		if collection == nil {
			result.status = -1
			result.description = "Unable to retrieve collection: " + BLACKLISTED_USERS_COLLECTION
		} else {
			_, err := collection.InsertOne(context.TODO(), user)
			if err != nil {
				panic(err)
				result.status = -1
				result.description = "Unable to insert element to collection: " + BLACKLISTED_USERS_COLLECTION
			}
			result.status = 0
		}
	} else {
		result.status = -1
		result.description = "Unable to open database: " + MONGO_DB_URI + " " + MONGO_DB_NAME
	}
	return result

}

func (repository ComplianceRepositoryImpl) findAll() []UserData {
	client := connectToDb()
	var users []UserData
	if client != nil {
		collection := client.Collection(BLACKLISTED_USERS_COLLECTION)
		if collection == nil {
			panic("unable to retrieve collection: " + BLACKLISTED_USERS_COLLECTION)
		} else {
			cursor, err := collection.Find(context.TODO(), bson.D{})
			if err != nil {
				panic("error retrieving data")
			}
			cursor.All(context.TODO(), &users)
			return users
		}
	} else {
		panic("unable to retrieve users")
		return nil
	}
}

func connectToDb() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_DB_URI))
	if err != nil {
		panic(err)
		return nil
	} else {
		return client.Database(MONGO_DB_NAME)
	}
}
