package database

import (
"go.mongodb.org/mongo-driver/mongo"
"fmt"
"log"
"time"
"os"
"context"
"github.com/joho/godotenv"
"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error Loading .env file")
	}
	MongoDb := os.Getenv("MONGODB_URL")
	if err!=nil{
		log.Fatal(err)
	}
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	defer cancel()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("connected to MongoDB")
	return client
}

var Client*mongo.Client = DBInstance()

func OpenCollection(client*mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}