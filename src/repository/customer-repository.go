package repository

import (
	"github.com/nuno/nunes-jumia/src/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source=./customer-repository.go -destination=./mocks/customer-repository_mock.go
type CustomerRepository interface {
	GetCustomers() ([]model.Customer, error)
}

type customerRepository struct {
	session *gorm.DB
}

func NewCustomerRepository(session *gorm.DB) CustomerRepository {
	return &customerRepository{
		session: session,
	}
}

func (repository customerRepository) GetCustomers() ([]model.Customer, error) {
	var customers []model.Customer

	result := repository.session.Find(&customers)
	return customers, result.Error
}
