package controller

import (
	"errors"
	"fmt"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/entity"
	"net/http"
	"strings"
)

const (
	baseUrl    = "http://localhost:8080"
	defaultUri = "/customer/search"
	emptyBody  = ""
	valid      = "Valid"
	invalid    = "Invalid"
)

type testScenario struct {
	TestName              string
	Method                string
	Uri                   string
	ReqParams             map[string]interface{}
	BodyString            string
	ShouldMockServiceCall bool
	Limit                 int
	Offset                int
	MockParams            map[string]string
	MockConsumerDto       dto.CustomerOutputDto
	MockErr               error
	ExpectStatus          int
	ExpectResponse        string
}

func MakeScenarioWithoutParamsExpectCustomersDtoWithOneCustomer() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search resource receiving customerDto with one invalid customer",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             nil,
		BodyString:            emptyBody,
		ShouldMockServiceCall: true,
		Limit:                 defaultLimit,
		Offset:                defaultOffset,
		MockParams:            make(map[string]string),
		MockConsumerDto: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  1,
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
		ExpectResponse: `{"limit":10,"offset":0,"total":1,"customers":[{"customer_name":"John Doe","country_name":"Morocco","country_code":"212","phone_number":"6007989253","status":"Invalid"}]}`,
	}
}

func MakeScenarioWithoutParamsExpectCustomersDtoWithTwoCustomers() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search resource receiving customerDto with two customers - invalid and valid",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             nil,
		BodyString:            emptyBody,
		ShouldMockServiceCall: true,
		Limit:                 defaultLimit,
		Offset:                defaultOffset,
		MockParams:            make(map[string]string),
		MockConsumerDto: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  2,
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
		ExpectResponse: `{"limit":10,"offset":0,"total":2,"customers":[{"customer_name":"John Doe","country_name":"Morocco","country_code":"212","phone_number":"6007989253","status":"Invalid"},{"customer_name":"James Smith","country_name":"Morocco","country_code":"212","phone_number":"633963130","status":"Valid"}]}`,
	}
}

func MakeScenarioExpectErrorToExtractLimitQueryParam() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search?limit=ABC resource sending a invalid limit value and receiving an error",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             map[string]interface{}{limitKey: "ABC"},
		BodyString:            emptyBody,
		ShouldMockServiceCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"error":"error to parse limit"}`,
	}
}

func MakeScenarioExpectErrorToExtractOffsetQueryParam() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search?offset=ABC resource sending a invalid offset value and receiving an error",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             map[string]interface{}{offsetKey: "ABC"},
		BodyString:            emptyBody,
		ShouldMockServiceCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"error":"error to parse offset"}`,
	}
}

func MakeScenarioExpectCustomersDtoWithLimit1() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search?limit=1 resource sending limit value and receiving customerDto with one invalid customer",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             map[string]interface{}{limitKey: "1"},
		BodyString:            emptyBody,
		ShouldMockServiceCall: true,
		Limit:                 1,
		Offset:                defaultOffset,
		MockParams:            make(map[string]string),
		MockConsumerDto: dto.CustomerOutputDto{
			Limit:  1,
			Offset: defaultOffset,
			Total:  1,
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
		ExpectResponse: `{"limit":1,"offset":0,"total":1,"customers":[{"customer_name":"John Doe","country_name":"Morocco","country_code":"212","phone_number":"6007989253","status":"Invalid"}]}`,
	}
}

func MakeScenarioExpectCustomersDtoWithOffset1() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search?offset=1 resource sending offset value and receiving customerDto with one valid customer",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             map[string]interface{}{offsetKey: 1},
		BodyString:            emptyBody,
		ShouldMockServiceCall: true,
		Limit:                 defaultLimit,
		Offset:                1,
		MockParams:            make(map[string]string),
		MockConsumerDto: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: 1,
			Total:  1,
			Customers: []dto.Customer{
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
		ExpectResponse: `{"limit":10,"offset":1,"total":1,"customers":[{"customer_name":"James Smith","country_name":"Morocco","country_code":"212","phone_number":"633963130","status":"Valid"}]}`,
	}
}

func MakeScenarioFilteringByCountryNameExpectCustomersDtoWithOneCustomer() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search?country_name=morocco resource sending country_name value and receiving customerDto with one invalid customer",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             map[string]interface{}{countryNameKey: entity.NameMorocco},
		BodyString:            emptyBody,
		ShouldMockServiceCall: true,
		Limit:                 defaultLimit,
		Offset:                defaultOffset,
		MockParams:            map[string]string{countryNameKey: entity.NameMorocco},
		MockConsumerDto: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  1,
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
		ExpectResponse: `{"limit":10,"offset":0,"total":1,"customers":[{"customer_name":"John Doe","country_name":"Morocco","country_code":"212","phone_number":"6007989253","status":"Invalid"}]}`,
	}
}

func MakeScenarioFilteringByStatusExpectCustomersDtoWithOneCustomer() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search?status=invalid resource sending status value and receiving customerDto with one invalid customer",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             map[string]interface{}{statusKey: invalid},
		BodyString:            emptyBody,
		ShouldMockServiceCall: true,
		Limit:                 defaultLimit,
		Offset:                defaultOffset,
		MockParams:            map[string]string{statusKey: invalid},
		MockConsumerDto: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  1,
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
		ExpectResponse: `{"limit":10,"offset":0,"total":1,"customers":[{"customer_name":"John Doe","country_name":"Morocco","country_code":"212","phone_number":"6007989253","status":"Invalid"}]}`,
	}
}

func MakeScenarioExpectErrorInServiceLayer() *testScenario {
	return &testScenario{
		TestName:              "Test /customer/search resource receiving internal server error",
		Method:                http.MethodGet,
		Uri:                   defaultUri,
		ReqParams:             nil,
		BodyString:            emptyBody,
		ShouldMockServiceCall: true,
		Limit:                 defaultLimit,
		Offset:                defaultOffset,
		MockParams:            map[string]string{},
		MockConsumerDto:       dto.CustomerOutputDto{},
		MockErr:               errors.New("mock_internal_server_error"),
		ExpectStatus:          http.StatusInternalServerError,
		ExpectResponse:        `{"error":"mock_internal_server_error"}`,
	}
}

func (testScenario *testScenario) getFullUrl() (fullUrl string) {
	fullUrl = fmt.Sprintf("%s%s", baseUrl, testScenario.Uri)

	if len(testScenario.ReqParams) > 0 {
		fullUrl += "?"
		for key, param := range testScenario.ReqParams {
			fullUrl += fmt.Sprintf("%s=%v&", key, param)
		}
		fullUrl = strings.TrimSuffix(fullUrl, "&")
	}
	return
}
