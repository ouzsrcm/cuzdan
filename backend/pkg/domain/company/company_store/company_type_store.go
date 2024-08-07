package company_store

import (
	"immortality/pkg/domain/company/company_models"

	"gorm.io/gorm"
)

// GetCompanyTypeList gets a list of all company types
func (store *CompanyStore) GetCompanyTypeList() (*[]company_models.CompanyType, error) {

	res := []company_models.CompanyType{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&res).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetCompanyType gets a company type by id
func (store *CompanyStore) GetCompanyType(id int) (*company_models.CompanyType, error) {

	res := company_models.CompanyType{}

	err := store.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).First(&res).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateCompanyType creates a new company type
func (store *CompanyStore) CreateCompanyType(model *company_models.CompanyType) error {

	res := store.Db.Transaction(func(tx *gorm.DB) error {

		err := tx.Create(&model)
		if err.Error != nil {
			return err.Error
		}
		return nil
	})
	if res != nil {
		return res
	}

	return nil
}

// RemoveCompanyType removes a company type
func (store *CompanyStore) RemoveCompanyType(id int) error {

	res := store.Db.Transaction(func(tx *gorm.DB) error {

		err := tx.Where("id = ?", id).Delete(&company_models.CompanyType{})

		if err.Error != nil {
			return err.Error
		}

		return nil
	})

	if res != nil {
		return res
	}

	return nil
}
