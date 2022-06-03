package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDatabase() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic(err)
	} else {
		collection := client.Database("compliancedb").Collection("blacklistedUsers")
		fmt.Println(collection)

		user := bson.D{{"fullName", "User 1"}, {"age", 30}}

		result, err := collection.InsertOne(context.TODO(), user)

		if err != nil {
			panic(err)
		}
		// display the id of the newly inserted object
		fmt.Println(result.InsertedID)
	}
}
