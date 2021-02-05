package repository

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/seigaalghi/firestore-go/models"
)

var (
	collectionUsers = "users"
)

//FindUser by email
func FindUser(email string) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	data, err := client.Collection(collectionUsers).Where("email", "==", strings.ToLower(email)).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("Invalid email or password")
	}
	result := models.User{
		Name:     data[0].Data()["name"].(string),
		Email:    data[0].Data()["email"].(string),
		Password: data[0].Data()["password"].(string),
	}
	return &result, nil

}

//SaveUser is ...
func SaveUser(request *models.User) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	ref, _, err := client.Collection(collectionUsers).Add(ctx, map[string]interface{}{
		"name":     request.Name,
		"email":    strings.ToLower(request.Email),
		"password": request.Password,
	})
	if err != nil {
		return nil, err
	}
	request.ID = ref.ID
	return request, nil
}
