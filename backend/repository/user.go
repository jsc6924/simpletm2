package repository

import (
	"backend/config"
	"backend/entities"
	"backend/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateUser(username, hashedPassword, token string) error {
	user := models.User{
		ID:    username,
		Salt:  hashedPassword,
		Token: token,
	}
	err := user.Insert(context.Background(), config.DB, boil.Infer())
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	user, err := models.Users(models.UserWhere.ID.EQ(username)).One(context.Background(), config.DB)
	return user, err
}

func GetPermission(username string, game string) (entities.Permission, error) {
	var permission entities.Permission
	err := config.DB.QueryRow("SELECT permission FROM Permission WHERE user_id = ? AND game_id = ?", username, game).Scan(&permission)
	return permission, err
}

func InsertOrSetPermission(username string, game string, permission entities.Permission) error {
	_, err := config.DB.Exec("INSERT INTO Permission (user_id, game_id, permission) VALUES (?, ?, ?) ON CONFLICT(user_id, game_id) DO UPDATE SET permission = ?", username, game, permission, permission)
	return err
}

func GetToken(username string) (string, error) {
	var token string
	err := config.DB.QueryRow("SELECT token FROM User WHERE id = ?", username).Scan(&token)
	return token, err
}

func SetToken(username string, token string) error {
	_, err := config.DB.Exec("UPDATE User SET token = ? WHERE id = ?", token, username)
	return err
}

func GetGamesByUser(username string) ([]*entities.GameWithPermission, error) {
	query := `SELECT p.game_id, g.title, p.permission from Permission AS p 
            JOIN User AS u ON p.user_id=u.id
            JOIN Game AS g ON p.game_id=g.id
            WHERE p.user_id=?`
	rows, err := config.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var games []*entities.GameWithPermission
	for rows.Next() {
		var game entities.GameWithPermission
		err = rows.Scan(&game.Game, &game.GameTitle, &game.Permission)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	return games, nil
}

func GetUsersByGame(game string) ([]*entities.UserWithPermission, error) {
	query := `SELECT p.user_id, p.permission from Permission AS p 
            JOIN User AS u ON p.user_id=u.id
            JOIN Game AS g ON p.game_id=g.id
            WHERE p.game_id=?`
	rows, err := config.DB.Query(query, game)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*entities.UserWithPermission
	for rows.Next() {
		var user entities.UserWithPermission
		err = rows.Scan(&user.Username, &user.Permission)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
