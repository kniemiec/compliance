package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ComplianceRepository interface {
	initializeRepository()
	findByNameAndLastName(lastName string, name string) []SanctionedUser
}

type ComplianceRepositoryImpl struct {
}

func (repository ComplianceRepositoryImpl) findByNameAndLastName(lastName string, name string) []SanctionedUser {
	var response []SanctionedUser
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic(err)
	} else {
		collection := client.Database("compliancedb").Collection("blacklistedUsers")

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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic(err)
	} else {
		collection := client.Database("compliancedb").Collection("blacklistedUsers")
		fmt.Println(collection)

		user := bson.D{{"lastName", "Andrzej"}, {"name", "Nowak"}}
		result, err := collection.InsertOne(context.TODO(), user)

		if err != nil {
			panic(err)
		}
		// display the id of the newly inserted object
		fmt.Println(result.InsertedID)
	}
}
