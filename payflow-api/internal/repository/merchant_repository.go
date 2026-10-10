package repository

import (
	"database/sql"
	"errors"
	"uuid"

	"payflow/internal/model"
)
type MerchantRepository struct {
	db *sql.DB
}

func NewMerchantRepository(db *sql.DB) *MerchantRepository {
	return &MerchantRepository{
		db:db,
	}
}

func (r *MerchantRepository) GetAll() ([]model.Merchant, error) {
	query := `
		SELECT id, business_name, business_type, owner_name, phone, address, status
		FROM merchants
		WHERE status = 'ACTIVE'
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
    if err != nil {
		return nil, err
	}
    defer rows.Close()

	var merchants []model.Merchant

	for rows.Next(){
		var merchant model.Merchant
		err := rows.Scan(
			&merchant.ID,
			&merchant.BusinessName,
			&merchant.BusinessType,
			&merchant.OwnerName,
            &merchant.Phone,
			&merchant.Address,
			&merchant.Status,
		)

		if err != nil {
			return nil, err
		}

		merchants = append(merchants, merchant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return merchants, nil
}

func (r *MerchantRepository) Create(req model.Merchant) (model.Merchant, error) {
	query := `
		INSERT INTO merchants (
		id,
		business_name,
		business_type,
		owner_name,
		phone,
		address
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING status
	`

	err := r.db.QueryRow(
		query,
		req.ID,
		req.BusinessName,
		req.BusinessType,
		req.OwnerName,
		req.Phone,
		req.Address,
	).Scan(&req.Status)

	if err != nil {
		return model.Merchant{}, err
	}

	return req, nil

}

func (r *MerchantRepository) GetByID(id uuid.UUID) (model.Merchant, error) {
	query := `
	SELECT id, business_name, business_type, owner_name, phone, address, status
	FROM merchants
	WHERE id = $1 AND status = 'ACTIVE'
	`

	row := r.db.QueryRow(query, id)

    var merchant model.Merchant
	
	err := row.Scan(
		&merchant.ID,
		&merchant.BusinessName,
		&merchant.BusinessType,
		&merchant.OwnerName,
		&merchant.Phone,
		&merchant.Address,
		&merchant.Status,
	)	

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Merchant{}, errors.New("merchant not found")
		}
		return model.Merchant{}, err
	}
	
	return merchant, nil

} 

func (r *MerchantRepository) Update(merchant model.Merchant) (model.Merchant, error) {
	query := `
	UPDATE merchants
	SET business_name = $1, business_type = $2, owner_name = $3, phone = $4, address = $5
	WHERE id = $6
	`
	result, err := r.db.Exec(
		query, 
		merchant.BusinessName,
		merchant.BusinessType,
		merchant.OwnerName,
		merchant.Phone,
		merchant.Address,
		merchant.ID,
	)

	if err != nil {
		return model.Merchant{}, err
	}

	rowsAffected, err := result.RowsAffected(); 

	if err != nil {
		return model.Merchant{}, err
	}

	
	if rowsAffected == 0 {
		return model.Merchant{}, errors.New("merchant not found")
	}

	return merchant, nil
}

func (r *MerchantRepository) Delete(id uuid.UUID) error {
	query := `
	UPDATE merchants
	SET status = 'INACTIVE'
	WHERE id = $1
		AND status = 'ACTIVE'
	`
	result, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

    rowsAffected, err := result.RowsAffected();
	if err != nil {
		return err
	}

	if rowsAffected == 0 { 
		return errors.New("merchant not found")
	}

	return nil
}

func (r *MerchantRepository) Activate(id uuid.UUID) error {
	
	query := `
	UPDATE merchants 
	SET status = 'ACTIVE'
	WHERE id = $1 AND status = 'INACTIVE'
	`

	result, err := r.db.Exec(query,id)

	if err != nil {
		return  err
	}

	rowsAffected, err := result.RowsAffected() 

	if err != nil {
		return err
	}

	if rowsAffected == 0 { 
		return errors.New("merchant not found")
	}

	return nil
	
}