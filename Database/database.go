package database

import (
	"context"
	"fmt"
	"os"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// Connecting to the database
	client,err:=mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	// Showing Error if occured
	if err!=nil{
		log.Fatal(err)
	}

	// Setting a time limit so that the operation will not take more time than mentioned
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)

	// Trying to connect
	err=client.Connect(ctx);

	// cancel releases the memory taken by temporary values
	// and defer runs it after the surrounding function returns
	defer cancel()

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB")

	return client
}