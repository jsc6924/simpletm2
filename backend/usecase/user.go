package usecase

import (
	"backend/entities"
	"backend/repository"
	"log"
)

type UserUsecase struct {
}

func (u *UserUsecase) CheckPassword(username string, hashedPassword string) (bool, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil || user == nil {
		return false, err
	}
	return user.Salt == hashedPassword, nil
}

func (u *UserUsecase) CheckToken(username string, token string) (bool, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil || user == nil {
		return false, err
	}
	return user.Token == token, nil
}

func (u *UserUsecase) CheckPermission(username string, game string, level entities.Permission) bool {
	permission, err := repository.GetPermission(username, game)
	if err != nil || int(permission) < int(level) {
		log.Printf("Error checking permission: %v\n", err)
		return false
	}
	return true
}
