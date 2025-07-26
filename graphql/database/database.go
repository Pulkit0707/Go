package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Pulkit0707/graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb://localhost:27017/graphql"

type DB struct {
	Client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	// Optional: Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	return &DB{Client: client}
}

func (db*DB) GetJob(id string) *model.JobListing{
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx,cancel:=context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id,_ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err:=jobCollection.FindOne(ctx, filter).Decode(&jobListing)
	if err!=nil{
		log.Fatal(err)
	}
	return &jobListing
}

func (db*DB) GetJobs()[]*model.JobListing{
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx,cancel:=context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var jobListings []*model.JobListing
	cursor,err:=jobCollection.Find(ctx,bson.D{})
	if err!=nil{
		log.Fatal(err)
	}
	if err=cursor.All(context.TODO(),&jobListings); err!=nil{
		panic(err)
	}
	return jobListings
}

func (db*DB) CreateJobListing(jobInfo model.CreateJobListingInput)*model.JobListing{
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx,cancel:=context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserted,err:=jobCollection.InsertOne(ctx,bson.M{
		"title": jobInfo.Title,
		"description": jobInfo.Description,
		"company": jobInfo.Company,
		"url": jobInfo.URL,
	})
	if err!=nil{
		log.Fatal(err)
	}
	insertedId:=inserted.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{
		ID: insertedId,
		Title: jobInfo.Title,
		Description: jobInfo.Description,
		Company: jobInfo.Company,
		Url: jobInfo.URL, 
	}
	return &returnJobListing
}

func (db*DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing{
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx,cancel:=context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	updatedJobInfo := bson.M{}
	if jobInfo.Title!=nil{
		updatedJobInfo["title"]=jobInfo.Title
	}
	if jobInfo.Description!=nil{
		updatedJobInfo["description"]=jobInfo.Description
	}
	if jobInfo.URL!=nil{
		updatedJobInfo["url"]=jobInfo.URL
	}
	_id, err := primitive.ObjectIDFromHex(jobId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updatedJobInfo}
	var jobListing model.JobListing
	results := jobCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))
	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func (db*DB) DeleteJobListing(jobId string)*model.DeleteJobResponse{
	jobCollection := db.Client.Database("graphql-job-board").Collection("jobs")
	ctx,cancel:=context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id, err := primitive.ObjectIDFromHex(jobId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": _id}
	_,err = jobCollection.DeleteOne(ctx,filter)
	if err!=nil{
		log.Fatal(err)
	}
	return &model.DeleteJobResponse{DeleteJobID: jobId}
}
