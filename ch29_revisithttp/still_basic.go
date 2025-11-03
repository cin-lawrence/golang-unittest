package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
}

type UserService interface {
	Register(user User) (insertedID string, err error)
}

type UserServer struct {
	service UserService
}

func NewUserServer(service UserService) *UserServer {
	return &UserServer{service: service}
}

func (s *UserServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode user payload: %v", err), http.StatusBadRequest)
		return
	}

	insertedID, err := s.service.Register(newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem registering new user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, insertedID)
}

type MongoUserService struct {
}

func NewMongoUserService() *MongoUserService {
	return &MongoUserService{}
}

func (s MongoUserService) Register(user User) (insertedID string, err error) {
	panic("implement me")
}

func main() {
	mongoService := NewMongoUserService()
	server := NewUserServer(mongoService)
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(server.RegisterUser)))
}
