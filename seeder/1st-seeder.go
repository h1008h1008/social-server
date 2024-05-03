package seeder

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedUsers(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database := os.Getenv("DATABASE_NAME")
	usersCollection := client.Database(database).Collection("users")

	// Predefined users
	users := []interface{}{
		bson.M{
			"userId":            "user1",
			"username":          "john_doe",
			"email":             "john@example.com",
			"passwordHash":      "hashed_password_here",
			"profilePictureUrl": "http://example.com/image.jpg",
			"bio":               "Hello, I'm John!",
			"createdAt":         time.Now(),
		},
		bson.M{
			"userId":            "user2",
			"username":          "jane_doe",
			"email":             "jane@example.com",
			"passwordHash":      "hashed_password_here",
			"profilePictureUrl": "http://example.com/image.jpg",
			"bio":               "Hello, I'm Jane!",
			"createdAt":         time.Now(),
		},
	}

	_, err := usersCollection.InsertMany(ctx, users)
	if err != nil {
		return err
	}

	return nil
}
