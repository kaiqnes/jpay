package dto

import (
	"github.com/nuno/nunes-jumia/src/entity"
	"github.com/nuno/nunes-jumia/src/model"
	"regexp"
	"strings"
)

type CustomerOutputDto struct {
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Total  int64      `json:"total"`
	Phones []Customer `json:"phones"`
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

func NewCustomerOutputDto(total int64, limit, offset int, customersData []model.Customer) (outputDto CustomerOutputDto) {
	matcher, _ := regexp.Compile("^\\((\\d{3})\\)\\s((?:.*))$")

	for _, customerData := range customersData {
		var (
			matches = matcher.FindStringSubmatch(customerData.Phone)
		)

		outputDto.Phones = append(outputDto.Phones, NewCustomer(customerData, matches))
	}

	outputDto.Total = total
	outputDto.Limit = limit
	outputDto.Offset = offset

	return
}

func NewCustomer(customerData model.Customer, matches []string) (customerInfo Customer) {
	customerInfo.SetFormattedName(customerData.Name)

	if matches == nil || len(matches) == 0 {
		customerInfo.PhoneNumber = customerData.Phone
		customerInfo.CountryName = "undefined_country_name"
		customerInfo.CountryCode = "undefined_country_code"
		customerInfo.Status = "Invalid"
		return customerInfo
	}

	countryCode := matches[1]
	phoneNumber := matches[2]
	customerCountry := entity.Countries[countryCode]

	customerInfo.PhoneNumber = phoneNumber
	customerInfo.CountryCode = countryCode
	customerInfo.CountryName = customerCountry.Name
	customerInfo.Status = customerCountry.IsValidPhoneNumber(customerData.Phone)

	return customerInfo
}
