package dto

import (
	"strings"
)

type CustomerOutputDto struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	CustomerName string `json:"customer_name"`
	CountryName  string `json:"country_name"`
	CountryCode  string `json:"country_code"`
	PhoneNumber  string `json:"phone_number"`
	Status       string `json:"status"`
}

func NewCustomer(name, phoneNumber, countryCode, countryName, status string) Customer {
	return Customer{
		CustomerName: strings.Title(strings.ToLower(name)),
		CountryName:  countryName,
		CountryCode:  countryCode,
		PhoneNumber:  phoneNumber,
		Status:       status,
	}
}
