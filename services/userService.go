package services

import (
	"Go-MongoDb-Api/dto"
	"Go-MongoDb-Api/models"
	"Go-MongoDb-Api/repository"
)

type DefaultUserService struct {
	Repo repository.UserRepository
}

type UserService interface {
	UserInsert(user models.User) (*dto.UserDTO, error)
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

func NewUserService(Repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{Repo: Repo}
}
