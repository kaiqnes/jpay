package controller

import (
	"errors"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/entity"
	"net/http"
)

const (
	defaultUrl = "http://localhost:8080/customers"
	valid      = "Valid"
	invalid    = "Invalid"
)

type testScenario struct {
	TestName              string
	Method                string
	Uri                   string
	ShouldMockServiceCall bool
	MockConsumerDto       dto.CustomerOutputDto
	MockErr               error
	ExpectStatus          int
	ExpectResponse        string
}

func MakeScenarioExpectCustomersDtoWithOneCustomer() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search resource receiving customerDto with one invalid customer",
		Method:                http.MethodGet,
		Uri:                   defaultUrl,
		ShouldMockServiceCall: true,
		MockConsumerDto: dto.CustomerOutputDto{
			Customers: []dto.Customer{
				{
					CustomerName: "John Doe",
					CountryName:  entity.NameMorocco,
					CountryCode:  entity.CodeMorocco,
					PhoneNumber:  "6007989253",
					Status:       invalid,
				},
			},
		},
		MockErr:        nil,
		ExpectStatus:   http.StatusOK,
		ExpectResponse: `{"customers":[{"customer_name":"John Doe","country_name":"Morocco","country_code":"212","phone_number":"6007989253","status":"Invalid"}]}`,
	}
}

func MakeScenarioExpectCustomersDtoWithTwoCustomers() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search resource receiving customerDto with two customers - invalid and valid",
		Method:                http.MethodGet,
		Uri:                   defaultUrl,
		ShouldMockServiceCall: true,
		MockConsumerDto: dto.CustomerOutputDto{
			Customers: []dto.Customer{
				{
					CustomerName: "John Doe",
					CountryName:  entity.NameMorocco,
					CountryCode:  entity.CodeMorocco,
					PhoneNumber:  "6007989253",
					Status:       invalid,
				},
				{
					CustomerName: "James Smith",
					CountryName:  entity.NameMorocco,
					CountryCode:  entity.CodeMorocco,
					PhoneNumber:  "633963130",
					Status:       valid,
				},
			},
		},
		MockErr:        nil,
		ExpectStatus:   http.StatusOK,
		ExpectResponse: `{"customers":[{"customer_name":"John Doe","country_name":"Morocco","country_code":"212","phone_number":"6007989253","status":"Invalid"},{"customer_name":"James Smith","country_name":"Morocco","country_code":"212","phone_number":"633963130","status":"Valid"}]}`,
	}
}

func MakeScenarioExpectErrorInServiceLayer() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search resource receiving internal server error",
		Method:                http.MethodGet,
		Uri:                   defaultUrl,
		ShouldMockServiceCall: true,
		MockConsumerDto:       dto.CustomerOutputDto{},
		MockErr:               errors.New("mock_internal_server_error"),
		ExpectStatus:          http.StatusInternalServerError,
		ExpectResponse:        `{"error":"mock_internal_server_error"}`,
	}
}
