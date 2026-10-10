package model

import "uuid"

type Merchant struct {
	ID           uuid.UUID     `json:"id"`
	BusinessName string `json:"business_name"`
	BusinessType string `json:"business_type"`
	OwnerName    string `json:"owner_name"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Status       string `json:"status"`
}

type CreateMerchantRequest struct {
	BusinessName string `json:"business_name"`
	BusinessType string `json:"business_type"`
	OwnerName    string `json:"owner_name"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
}