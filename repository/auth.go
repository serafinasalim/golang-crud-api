package repository

import (
	"database/sql"
	"golang-crud-api/model"
	"log"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	if db == nil {
		log.Fatal("Received nil database connection")
	}

	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(params model.User) (err error) {
	const query = `INSERT INTO public.users (
					username, 
					email, 
					password, 
					created_at
				) VALUES ($1, $2, $3, $4)`

	_, err = r.db.Exec(
		query,
		params.Username,
		params.Email,
		params.Password,
		params.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByUsername(username string) (result *model.User, err error) {
	var user model.User
	const query = `SELECT 
					id, 
					username, 
					email, 
					password, 
					created_at, 
					updated_at
				FROM public.users WHERE username = $1`

	err = r.db.QueryRow(query, username).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
