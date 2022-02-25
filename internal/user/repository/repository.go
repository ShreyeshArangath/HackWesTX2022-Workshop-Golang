package repository

import (
	"context"
	"github.com/ShreyeshArangath/hackwestx-workshop/internal/user/entity"
	"github.com/ShreyeshArangath/hackwestx-workshop/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// Get returns the user with the specified userId.
	Get(ctx context.Context, userId string) (entity.User, error)
	// Create saves a new user in the storage.
	Create(ctx context.Context, user entity.User) error
	// Update updates the user with given ID in the storage.
	Update(ctx context.Context, userId string, user entity.User) error
	// Delete removes the user with given userId from the storage.
	Delete(ctx context.Context, userId string) error
}

type repository struct {
	//TODO: Implement the DB Context
	db     db.DBhelper
	logger log.Logger
}

func NewRepository(db db.DBhelper, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, userId string) (entity.User, error) {
	var user entity.User
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return entity.User{}, err
	}

	filter := bson.M{"_id": id}
	collection, _ := r.db.GetMongoDBCollection("users")
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r repository) Create(ctx context.Context, user entity.User) error {
	collection, _ := r.db.GetMongoDBCollection("users")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) Update(ctx context.Context, userId string, user entity.User) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	// Figure out which ID we need to update
	filter := bson.M{"_id": id}
	collection, _ := r.db.GetMongoDBCollection("users")
	var newUser entity.User
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "avatarURL", Value: user.AvatarURL},
			{Key: "firstName", Value: user.FirstName},
			{Key: "lastName", Value: user.LastName},
			{Key: "email", Value: user.Email},
			{Key: "createdAt", Value: user.CreatedAt},
			{Key: "updatedAt", Value: time.Now()},
		},
		},
	}
	err = collection.FindOneAndUpdate(ctx, filter, update).Decode(&newUser)
	if err != nil {
		return err
	}
	newUser.UserId = id.String()

	return nil
}

func (r repository) Delete(ctx context.Context, userId string) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	collection, _ := r.db.GetMongoDBCollection("users")
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
