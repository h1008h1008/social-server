package migration

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	usersCollection := client.Database("socialdb").Collection("users")
	_, err := usersCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"username": 1},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return err
	}

	postsCollection := client.Database("socialdb").Collection("posts")
	_, err = postsCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"authorId": 1},
	})
	if err != nil {
		return err
	}
	commentsCollection := client.Database("socialdb").Collection("comments")
	_, err = commentsCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"postId": 1},
		},
		{
			Keys: bson.M{"authorId": 1},
		},
	})
	if err != nil {
		return err
	}

	friendshipsCollection := client.Database("socialdb").Collection("friendships")
	_, err = friendshipsCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "userId1", Value: 1}, {Key: "userId2", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return err
	}

	// Indexes for Messages collection
	messagesCollection := client.Database("socialdb").Collection("messages")
	_, err = messagesCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"senderUserId": 1},
		},
		{
			Keys: bson.M{"receiverUserId": 1},
		},
		{
			Keys: bson.M{"createdAt": -1},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func EnsureCollections(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collectionsToCreate := []string{"users", "posts", "comments", "friendships", "messages"}
	for _, coll := range collectionsToCreate {
		err := client.Database("socialdb").CreateCollection(ctx, coll)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				continue
			}
			return err
		}
	}

	return nil
}
