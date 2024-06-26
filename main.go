package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"social-server-module/database"
	"social-server-module/migration"
	"social-server-module/model"
	"social-server-module/seeder"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	client, err := database.ConnectMongo(uri)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}
	defer client.Disconnect(context.Background())

	if err := migration.EnsureCollections(client); err != nil {
		log.Fatal("Error ensuring collections: ", err)
	}

	if err := migration.CreateIndexes(client); err != nil {
		log.Fatal("Error creating indexes: ", err)
	}

	if err := seeder.SeedUsers(client); err != nil {
		log.Fatal("Failed to seed users: ", err)
	}

	usersCollection := client.Database("socialdb").Collection("users")

	database := os.Getenv("DATABASE_NAME")
	usersCollection := client.Database(database).Collection("users")
	fmt.Println(usersCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error finding users: ", err)
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor
	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal("Error decoding user: ", err)
		}
		fmt.Printf("User: %+v\n", user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor error: ", err)
	}

	fmt.Println("Database setup complete.")
}
