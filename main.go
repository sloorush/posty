package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"posty/api/handler"

	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("DB connection success!!", db)

	http.HandleFunc("/ping", handler.PingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DatabaseConnection() (*mongo.Database, error) {
	// log.Println(os.Getenv("DB_URL"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		return nil, err
	}
	db := client.Database("posty")
	return db, nil
}
