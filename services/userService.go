package services

import (
	"Go-MongoDb-Api/dto"
	"Go-MongoDb-Api/models"
	"Go-MongoDb-Api/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	UserInsert(user models.User) (*dto.UserDTO, error)
	UserGetAll() ([]models.User, error)
	UserDelete(id primitive.ObjectID) (bool, error)
}

func (u DefaultUserService) UserInsert(user models.User) (*dto.UserDTO, error) {
	var res dto.UserDTO
	if len(user.Email) <= 2 {
		res.Status = false
		return &res, nil
	}
	result, err := u.Repo.Insert(user)
	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.UserDTO{Status: result}
	return &res, nil

}

func (u DefaultUserService) UserGetAll() ([]models.User, error) {
	result, err := u.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u DefaultUserService) UserDelete(id primitive.ObjectID) (bool, error) {
	result, err := u.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}
	return true, nil
}

func NewUserService(Repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{Repo: Repo}
}
