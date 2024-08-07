package repository

import (
	"backend/config"
	"backend/entities"
	"backend/models"
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func InsertOrUpdateTranslate(game string, raw string, translate string) error {
	// start a transaction
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// check if the record exists
	count, err := models.Translates(models.TranslateWhere.GameID.EQ(game), models.TranslateWhere.RawWord.EQ(raw)).Count(context.Background(), tx)
	if err != nil {
		return err
	}
	newRecord := models.Translate{
		GameID:    game,
		RawWord:   raw,
		TransWord: null.String{String: translate, Valid: true},
	}
	if count == 0 {
		// Insert new record
		err = newRecord.Insert(context.Background(), tx, boil.Infer())
		if err != nil {
			return err
		}
	} else {
		// Update existing record
		_, err = newRecord.Update(context.Background(), tx, boil.Infer())
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

// DeleteTranslate deletes a translation record from the database
func DeleteTranslate(gameID string, rawWord string) error {
	// Start a transaction
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // Ensure rollback in case of error

	// Find the record to delete
	record, err := models.Translates(
		models.TranslateWhere.GameID.EQ(gameID),
		models.TranslateWhere.RawWord.EQ(rawWord),
	).One(context.Background(), tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No record found, nothing to delete
			return nil
		}
		return err
	}

	// Delete the record
	_, err = record.Delete(context.Background(), tx)
	if err != nil {
		return err
	}

	// Commit the transaction
	return tx.Commit()
}
