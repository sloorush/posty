package main_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"posty/api/handler"
	"posty/pkg/post"
	"posty/pkg/user"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() (*mongo.Database, error) {
	// log.Println(os.Getenv("DB_URL"))
	DB_URL := os.Getenv("DB_URL")
	decoded_db_url, err := url.QueryUnescape(DB_URL)

	// fmt.Println("\n" + decoded_db_url)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(decoded_db_url))
	if err != nil {
		return nil, err
	}
	db := client.Database("posty")
	return db, nil
}

func initialize() (userService user.Service, postService post.Service) {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("DB connected")
	userCollection := db.Collection("user")
	userRepo := user.NewRepo(userCollection)
	userService = user.NewService(userRepo)

	postCollection := db.Collection("post")
	postRepo := post.NewRepo(postCollection)
	postService = post.NewService(postRepo)

	return userService, postService
}

func TestServer(t *testing.T) {

	//Initial Setup
	// userService, postService := initialize()

	//Testing GetMeetingDetailsFromID
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.PingHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestReadUser(t *testing.T) {

	//Initial Setup
	userService, _ := initialize()

	//Testing GetMeetingDetailsFromID
	req, err := http.NewRequest("GET", "/users/61617212ce87a8d785217103", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetUser(userService))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestReadPost(t *testing.T) {

	//Initial Setup
	_, postService := initialize()

	//Testing GetMeetingDetailsFromID
	req, err := http.NewRequest("GET", "/posts/6161a86b9cb5cab0da8c0c82", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetPost(postService))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
