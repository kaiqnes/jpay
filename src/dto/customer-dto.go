package dto

import (
	"strings"
)

type CustomerOutputDto struct {
	Limit     int        `json:"limit"`
	Offset    int        `json:"offset"`
	Total     int        `json:"total"`
	Customers []Customer `json:"customers"`
}

type Customer struct {
	CustomerName string `json:"customer_name"`
	CountryName  string `json:"country_name"`
	CountryCode  string `json:"country_code"`
	PhoneNumber  string `json:"phone_number"`
	Status       string `json:"status"`
}

func (customer *Customer) SetFormattedName(name string) {
	customer.CustomerName = strings.Title(strings.ToLower(name))
}

func NewIdentifiedCustomer(name, phoneNumber, countryCode, countryName, status string) Customer {
	return Customer{
		CustomerName: strings.Title(strings.ToLower(name)),
		CountryName:  countryName,
		CountryCode:  countryCode,
		PhoneNumber:  phoneNumber,
		Status:       status,
	}
}

func NewUnidentifiedCustomer(name, phoneNumber string) Customer {
	return Customer{
		CustomerName: strings.Title(strings.ToLower(name)),
		CountryName:  "undefined_country_name",
		CountryCode:  "undefined_country_code",
		Status:       "Invalid",
		PhoneNumber:  phoneNumber,
	}
}
