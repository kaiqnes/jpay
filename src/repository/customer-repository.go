package repository

import (
	"github.com/nuno/nunes-jumia/src/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source=./customer-repository.go -destination=./mocks/customer-repository_mock.go
type CustomerRepository interface {
	GetCustomers(limit, offset int, queryParams map[string]string) (int64, []model.Customer, error)
}

type customerRepository struct {
	session *gorm.DB
}

func NewCustomerRepository(session *gorm.DB) CustomerRepository {
	return &customerRepository{
		session: session,
	}
}

func (repository customerRepository) GetCustomers(limit, offset int, queryParams map[string]string) (int64, []model.Customer, error) {
	var (
		customers []model.Customer
		total     int64
	)

	result := repository.session.Find(&customers)
	if result.Error != nil {
		return total, customers, result.Error
	}

	total = result.RowsAffected

	result.Limit(limit).Offset(offset).Order("phone asc").Find(&customers)
	return total, customers, result.Error
}
