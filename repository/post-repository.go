package repository

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/seigaalghi/firestore-go/models"
	"google.golang.org/api/iterator"
)

var (
	collectionPosts = "posts"
)

// FindAllPosts is a function to get all data of posts from database
func FindAllPosts() ([]models.Post, error) {
	var data []models.Post
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	iter := client.Collection(collectionPosts).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		result := models.Post{
			ID:        doc.Ref.ID,
			Title:     doc.Data()["title"].(string),
			Text:      doc.Data()["text"].(string),
			Date:      doc.Data()["date"].(time.Time),
			Price:     doc.Data()["price"].(int64),
			Authors:   doc.Data()["authors"].([]interface{}),
			Published: doc.Data()["published"].(bool),
		}

		data = append(data, result)
	}
	return data, nil
}

//SavePost is to Create new post
func SavePost(request *models.Post) (*models.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	ref, res, err := client.Collection(collectionPosts).Add(ctx, map[string]interface{}{
		"title":     request.Title,
		"text":      request.Text,
		"date":      firestore.ServerTimestamp,
		"price":     request.Price,
		"authors":   request.Authors,
		"published": request.Published,
	})
	if err != nil {
		return nil, err
	}
	request.ID = ref.ID
	request.Date = res.UpdateTime
	return request, nil
}

//UpdatePost is to Update a post
func UpdatePost(request *models.Post, ID string) (*models.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, err = client.Collection(collectionPosts).Doc(ID).Update(ctx, []firestore.Update{
		{
			Path:  "title",
			Value: request.Title,
		},
		{
			Path:  "text",
			Value: request.Text,
		},
		{
			Path:  "date",
			Value: request.Date,
		},
		{
			Path:  "price",
			Value: request.Price,
		},
		{
			Path:  "auhtors",
			Value: request.Authors,
		},
		{
			Path:  "published",
			Value: request.Published,
		},
	})
	if err != nil {
		return nil, err
	}
	return request, nil
}

//RemovePost is to Create new post
func RemovePost(ID string) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return err
	}
	defer client.Close()

	if _, err := client.Collection(collectionPosts).Doc(ID).Delete(ctx); err != nil {
		return err
	}
	return nil
}
