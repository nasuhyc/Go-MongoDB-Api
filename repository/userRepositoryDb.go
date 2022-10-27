package repository

import (
	"Go-MongoDb-Api/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryDB struct {
	UserCollection *mongo.Collection
}

type UserRepository interface {
	Insert(user models.User) (bool, error)
	GetAll() ([]models.User, error)
	Delete(id primitive.ObjectID) (bool, error)
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

func (u UserRepositoryDB) GetAll() ([]models.User, error) {
	var user models.User
	var users []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := u.UserCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&user); err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UserRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := u.UserCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil

}

func NewUserRepositoryDb(dbClient *mongo.Collection) UserRepositoryDB {
	return UserRepositoryDB{UserCollection: dbClient}
}
