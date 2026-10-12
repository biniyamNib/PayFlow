package model

import (
	"time"
	"uuid"
)

type Merchant struct {
	ID           uuid.UUID `json:"id"`
	BusinessName string    `json:"business_name"`
	BusinessType string    `json:"business_type"`
	OwnerName    string    `json:"owner_name"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Status       string    `json:"status"`
}

type MerchantUser struct {
	ID           uuid.UUID `json:"id"`
	MerchantID   uuid.UUID `json:"merchant_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateMerchantRequest struct {
	BusinessName string `json:"business_name"`
	BusinessType string `json:"business_type"`
	OwnerName    string `json:"owner_name"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
}

type RegisterMerchantRequest struct {
	BusinessName string `json:"business_name"`
	BusinessType string `json:"business_type"`
	OwnerName    string `json:"owner_name"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  MerchantUser `json:"user"`
}
