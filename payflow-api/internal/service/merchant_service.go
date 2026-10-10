package service

import (
	"payflow/internal/model"
	"payflow/internal/repository"
	"uuid"

	"payflow/internal/utils"
)

type MerchantService struct {
	repository *repository.MerchantRepository
}

func NewMerchantService(repo *repository.MerchantRepository) *MerchantService {
	return &MerchantService{
		repository: repo,
	}
}

func (s *MerchantService) GetAll() ([]model.Merchant, error) {
	return s.repository.GetAll()
}

func (s *MerchantService) Create(req model.CreateMerchantRequest) (model.Merchant, error) {
	if err := utils.ValidateRequired(req.BusinessName, "business_name"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.BusinessType, "business_type"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.OwnerName, "owner_name"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.Phone, "phone"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.Address, "address"); err != nil {
		return model.Merchant{}, err
	}

	merchant := model.Merchant{
		ID:           uuid.New(),
		BusinessName: req.BusinessName,
		BusinessType: req.BusinessType,
		OwnerName:    req.OwnerName,
		Phone:        req.Phone,
		Address:      req.Address,
	}

	createdMerchant, err := s.repository.Create(merchant)

	if err != nil {
		return model.Merchant{}, err
	}

	return createdMerchant, nil

}

func (s *MerchantService) GetByID(id uuid.UUID) (model.Merchant, error) {
	return s.repository.GetByID(id)
}

func (s *MerchantService) Update(id uuid.UUID, req model.CreateMerchantRequest) (model.Merchant, error) {
	if err := utils.ValidateRequired(req.BusinessName, "business_name"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.BusinessType, "business_type"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.OwnerName, "owner_name"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.Phone, "phone"); err != nil {
		return model.Merchant{}, err
	}
	if err := utils.ValidateRequired(req.Address, "address"); err != nil {
		return model.Merchant{}, err
	}

	merchant := model.Merchant{
		ID:           id,
		BusinessName: req.BusinessName,
		BusinessType: req.BusinessType,
		OwnerName:    req.OwnerName,
		Phone:        req.Phone,
		Address:      req.Address,
	}

	updatedMerchant, err := s.repository.Update(merchant)

	if err != nil {
		return model.Merchant{}, err
	}

	return updatedMerchant, nil
}

func (s *MerchantService) Delete(id uuid.UUID) error {
	return s.repository.Delete(id)
}

func (s *MerchantService) Activate(id uuid.UUID) error {
	return s.repository.Activate(id)
}
