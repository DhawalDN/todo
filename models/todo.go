package models

import (
	"GOLANG/todo/db"
	"GOLANG/todo/forms"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" binding:"required"`
	// CreatedAt time.Time          `json:"created_at"`
	// UpdatedAt time.Time          `json:"updated_at"`
	Completed bool `json:"completed"`
}

type TaskModel struct{}

var server = "127.0.0.1"

var client = db.GetClient()

func (m *TaskModel) Create(data forms.CreateTaskCommand) error {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	insertResult, err := collection.InsertOne(context.TODO(), bson.M{"title": data.Title, "completed": data.Completed})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Println("Inserted a single document:", insertResult)
	return err
}

func (m *TaskModel) Get() (list []*Task, err error) {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	cur, err := collection.Find(context.TODO(), bson.M{"completed": false})
	for cur.Next(context.TODO()) {
		var li Task
		err = cur.Decode(&li)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		list = append(list, &li)
	}
	defer client.Disconnect(ctx)

	return list, err
}

func (m *TaskModel) Completed() (list []*Task, err error) {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	cur, err := collection.Find(context.TODO(), bson.M{"completed": true})
	for cur.Next(context.TODO()) {
		var li Task
		err = cur.Decode(&li)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		list = append(list, &li)
	}
	defer client.Disconnect(ctx)

	return list, err
}

func (m *TaskModel) Update(data forms.UpdateTaskCommand) (err error) {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	oid, _ := primitive.ObjectIDFromHex(data.ID)
	filter := bson.M{"_id": oid}
	fmt.Println("IN model: ")
	update := bson.M{"$set": bson.M{"title": data.Title, "completed": data.Completed}}
	// update := bson.M{"name": data.Name, "email": data.Email, "gender": data.Gender}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return err
}

func (m *TaskModel) Delete(data forms.UpdateTaskCommand) (err error) {

	var client = db.GetClient()
	var ctx = context.Background()
	collection := client.Database("todo").Collection("tasks")
	oid, _ := primitive.ObjectIDFromHex(data.ID)
	filter := bson.M{"_id": oid}
	//update := bson.M{"$set": bson.M{"isDelete": true}}
	fmt.Println(filter)
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Printf("Deleted Count: %v ", deleteResult.DeletedCount)

	return err
}
