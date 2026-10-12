package service

import (
	"context"
	"database/sql"
	"payflow/internal/model"
	"payflow/internal/repository"
	"payflow/internal/utils"
	"uuid"
)

type MerchantUserService struct {
	db         *sql.DB
	repository *repository.MerchantUserRepository
}

func NewMerchantUserService(db *sql.DB) *MerchantUserService {
	return &MerchantUserService{
		db: db,
	}
}

func (s *MerchantUserService) Register(req model.RegisterMerchantRequest) (model.MerchantUser, error) {
	if err := utils.ValidateRequired(req.BusinessName, "business_name"); err != nil {
		return model.MerchantUser{}, err
	}
	if err := utils.ValidateRequired(req.BusinessType, "business_type"); err != nil {
		return model.MerchantUser{}, err
	}
	if err := utils.ValidateRequired(req.OwnerName, "owner_name"); err != nil {
		return model.MerchantUser{}, err
	}
	if err := utils.ValidateRequired(req.Phone, "phone"); err != nil {
		return model.MerchantUser{}, err
	}
	if err := utils.ValidateRequired(req.Email, "email"); err != nil {
		return model.MerchantUser{}, err
	}
	if err := utils.ValidateRequired(req.Password, "password"); err != nil {
		return model.MerchantUser{}, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return model.MerchantUser{}, err
	}

	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return model.MerchantUser{}, err
	}
	defer tx.Rollback()

	merchantRepository := repository.NewMerchantRepository(tx)
	merchantUserRepository := repository.NewMerchantUserRepository(tx)

	merchant := model.Merchant{
		ID:           uuid.New(),
		BusinessName: req.BusinessName,
		BusinessType: req.BusinessType,
		OwnerName:    req.OwnerName,
		Phone:        req.Phone,
		Address:      req.Address,
	}

	createdMerchant, err := merchantRepository.Create(merchant)
	if err != nil {
		return model.MerchantUser{}, err
	}

	merchantUser := model.MerchantUser{
		ID:           uuid.New(),
		MerchantID:   createdMerchant.ID,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Role:         "OWNER",
	}

	createdMerchantUser, err := merchantUserRepository.Create(merchantUser)
	if err != nil {
		return model.MerchantUser{}, err
	}

	if err := tx.Commit(); err != nil {
		return model.MerchantUser{}, err
	}

	return createdMerchantUser, nil
}
