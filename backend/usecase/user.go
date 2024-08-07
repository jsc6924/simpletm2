package usecase

import (
	"backend/entities"
	"backend/repository"
	"backend/utils"
	"log"
)

type UserUsecase struct {
}

func (u *UserUsecase) CreateUser(username string, password string) error {
	token, err := utils.GenToken(32)
	if err != nil {
		return err
	}
	hashedPassword := utils.Hash(password)
	return repository.CreateUser(username, hashedPassword, token)
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

func (u *UserUsecase) GetPermission(username string, game string) (entities.Permission, error) {
	return repository.GetPermission(username, game)
}

func (u *UserUsecase) SetPermission(username string, game string, permission entities.Permission) error {
	return repository.SetPermission(username, game, permission)
}

func (u *UserUsecase) GetGamesByUser(username string) ([]*entities.GameWithPermission, error) {
	return repository.GetGamesByUser(username)
}

func (u *UserUsecase) GetToken(username string) (string, error) {
	return repository.GetToken(username)
}

func (u *UserUsecase) UpdateToken(username string) error {
	token, err := utils.GenToken(32)
	if err != nil {
		return err
	}
	return repository.SetToken(username, token)
}

func (u *UserUsecase) GenerateSharedURL(username string, game string) (string, error) {
	tempUserSuffix, err := utils.GenToken(8)
	if err != nil {
		return "", err
	}
	tempUser := "_" + username + "_share_" + tempUserSuffix
	tempUserPassword, err := utils.GenToken(8)
	if err != nil {
		return "", err
	}
	tempUserToken, err := utils.GenToken(32)
	if err != nil {
		return "", err
	}
	tempUserPermission := entities.EDIT
	err = repository.CreateUser(tempUser, utils.Hash(tempUserPassword), tempUserToken)
	if err != nil {
		return "", err
	}
	err = repository.SetPermission(tempUser, game, tempUserPermission)
	if err != nil {
		return "", err
	}
	return utils.MakeSharedURL(tempUser, tempUserToken, game), nil
}
