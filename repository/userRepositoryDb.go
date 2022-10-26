package repository

import (
	"Go-MongoDb-Api/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryDB struct {
	UserCollection *mongo.Collection
}

type UserRepository interface {
	Insert(user models.User) (bool, error)
}

func (u UserRepositoryDB) Insert(user models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user.Id = primitive.NewObjectID()
	result, err := u.UserCollection.InsertOne(ctx, user)

	if result.InsertedID == nil || err != nil {
		errors.New("Ekleme işlemi başarısız")
		return false, err
	}
	return true, nil
}

func NewUserRepositoryDb(dbClient *mongo.Collection) UserRepositoryDB {
	return UserRepositoryDB{UserCollection: dbClient}
}
