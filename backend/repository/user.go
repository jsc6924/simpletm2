package repository

import (
	"backend/config"
	"backend/entities"
	"backend/models"
	"context"
)

func GetUserByUsername(username string) (*models.User, error) {
	user, err := models.Users(models.UserWhere.ID.EQ(username)).One(context.Background(), config.DB)
	return user, err
}

func GetPermission(username string, game string) (entities.Permission, error) {
	var permission entities.Permission
	err := config.DB.QueryRow("SELECT permission FROM Permission WHERE user_id = ? AND game_id = ?", username, game).Scan(&permission)
	return permission, err
}
