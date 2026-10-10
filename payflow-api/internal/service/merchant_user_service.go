package service

import (
	"payflow/internal/model"
	"payflow/internal/repository"
	"payflow/internal/utils"
	"uuid"
)

type MerchantUserService struct {
	repository *repository.MerchantUserRepository
}

func NewMerchantUserService(repo *repository.MerchantUserRepository) *MerchantUserService {
	return &MerchantUserService{
		repository: repo,
	}
}

func (s *MerchantUserService) Register(req model.RegisterMerchantRequest) (model.MerchantUser, error) {
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

	merchant_user := model.MerchantUser {
		ID: uuid.New(),
		Email: req.Email,
		PasswordHash: hashedPassword,
	}

	createdMerchantUser, err := s.repository.Create(merchant_user)

	 if err != nil {
		return model.MerchantUser{}, err
	 }

	 return createdMerchantUser, nil

}