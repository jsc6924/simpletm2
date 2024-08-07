package usecase

import (
	"backend/entities"
	"backend/repository"
)

type GameUsecase struct{}

func (u *GameUsecase) GetUsersByGame(game string) ([]*entities.UserWithPermission, error) {
	return repository.GetUsersByGame(game)
}

func (u *GameUsecase) CreateGameAsUser(gameId, gameTitle, username string) error {
	return repository.CreateGameAsUser(gameId, gameTitle, username)
}

func (u *GameUsecase) DeleteGame(gameId string) error {
	return repository.DeleteGame(gameId)
}
