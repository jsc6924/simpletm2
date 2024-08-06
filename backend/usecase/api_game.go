package usecase

import (
	"backend/entities"
	"backend/repository"
)

type ApiGameUsecase struct {
}

func (u *ApiGameUsecase) QueryByGame(game string) ([]*entities.Translate, error) {
	return repository.QueryAllTranslateByGame(game)
}
