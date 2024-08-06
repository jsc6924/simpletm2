package repository

import (
	"backend/config"
	"backend/entities"
	"backend/models"
	"context"
)

func QueryAllTranslateByGame(game string) ([]*entities.Translate, error) {
	translates, err := models.Translates(models.TranslateWhere.GameID.EQ(game)).All(context.Background(), config.DB)
	results := make([]*entities.Translate, 0)
	for _, t := range translates {
		if t.TransWord.Valid {
			results = append(results, &entities.Translate{
				Raw:       t.RawWord,
				Translate: t.TransWord.String,
			})
		} else {
			results = append(results, &entities.Translate{
				Raw:       t.RawWord,
				Translate: "",
			})
		}
	}
	return results, err
}
