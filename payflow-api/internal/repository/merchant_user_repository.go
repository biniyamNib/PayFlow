package repository

import (
	"fmt"
	"errors"
	"database/sql"
	"payflow/internal/model"

)

type MerchantUserRepository struct {
	db DBExecutor
}

func NewMerchantUserRepository(db DBExecutor) *MerchantUserRepository {
	return &MerchantUserRepository{
		db: db,
	}
}

func (r *MerchantUserRepository) Create(req model.MerchantUser) (model.MerchantUser, error) {
	query := `
	INSERT INTO merchant_users (
		id,
		merchant_id,
		email,
		password_hash,
		role
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING created_at
	`
	err := r.db.QueryRow(
		query,
		req.ID,
		req.MerchantID,
		req.Email,
		req.PasswordHash,
		req.Role,
	).Scan(&req.CreatedAt)

	if err != nil {
		return model.MerchantUser{}, fmt.Errorf("create merchant user: %w", err)
	}

	return req, nil

}

func (r *MerchantUserRepository) GetByEmail (email string) (model.MerchantUser, error) {
	query := `
	SELECT id, merchant_id, email, password_hash, role, created_at
	FROM merchant_users
	WHERE email = $1
	`
	row := r.db.QueryRow(query, email)

	var merchantUser model.MerchantUser

	err := row.Scan(
		&merchantUser.ID,
		&merchantUser.MerchantID,
		&merchantUser.Email,
		&merchantUser.PasswordHash,
		&merchantUser.Role,
		&merchantUser.CreatedAt,
	)

	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.MerchantUser{}, ErrMerchantUserNotFound
		}
		return model.MerchantUser{}, fmt.Errorf("get merchant user by email: %w", err)
	}

	return merchantUser, nil

}