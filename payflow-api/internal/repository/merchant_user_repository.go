package repository

import (
	"fmt"
	"payflow/internal/model"
)

type MerchantUserRepository struct {
	db DBExecutor
}

func NewMerchantUserRepository (db DBExecutor) *MerchantUserRepository {
	return &MerchantUserRepository{
		db:db,
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

