package service

import (
	"errors"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/entity"
	"github.com/nuno/nunes-jumia/src/model"
)

const (
	valid                  = "Valid"
	invalid                = "Invalid"
	defaultLimit           = 10
	defaultOffset          = 0
	mockRepositoryErrorMsg = "mock repository error"
)

var (
	mockRepositoryError = errors.New(mockRepositoryErrorMsg)
)

type testScenario struct {
	testName     string
	limit        int
	offset       int
	params       map[string]string
	mockTotal    int64
	mockResult   []model.Customer
	mockErr      error
	expectResult dto.CustomerOutputDto
	expectErr    error
}

func MakeScenarioWithoutParamsExpectDtoFilledWithSingleCustomerAndErrorNil() *testScenario {
	return &testScenario{
		testName:  "Get a customer with received limit and offset, without extra params",
		limit:     defaultLimit,
		offset:    defaultOffset,
		params:    map[string]string{},
		mockTotal: int64(1),
		mockResult: []model.Customer{
			{
				Id:    1,
				Name:  "John Doe",
				Phone: "(212) 633963130",
			},
		},
		mockErr: nil,
		expectResult: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  1,
			Customers: []dto.Customer{
				{
					CustomerName: "John Doe",
					CountryName:  "Morocco",
					CountryCode:  entity.CodeMorocco,
					PhoneNumber:  "633963130",
					Status:       valid,
				},
			},
		},
		expectErr: nil,
	}
}

func MakeScenarioWithoutParamsExpectDtoEmptyAndError() *testScenario {
	return &testScenario{
		testName:     "Get a customer with received limit and offset, without extra params and returns error",
		limit:        defaultLimit,
		offset:       defaultOffset,
		params:       map[string]string{},
		mockTotal:    int64(0),
		mockResult:   []model.Customer{},
		mockErr:      mockRepositoryError,
		expectResult: dto.CustomerOutputDto{},
		expectErr:    errors.New("Fail to retrieve customers in DB. Err: mock repository error"),
	}
}

func MakeScenarioWithoutParamsExpectDtoFilledWithInvalidCustomerAndErrorNil() *testScenario {
	return &testScenario{
		testName:  "Get a customer without received limit, offset and extra params",
		limit:     defaultLimit,
		offset:    defaultOffset,
		params:    map[string]string{},
		mockTotal: int64(1),
		mockResult: []model.Customer{
			{
				Id:    1,
				Name:  "John Doe",
				Phone: "+123 633963130",
			},
		},
		mockErr: nil,
		expectResult: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  1,
			Customers: []dto.Customer{
				{
					CustomerName: "John Doe",
					CountryName:  "undefined_country_name",
					CountryCode:  "undefined_country_code",
					PhoneNumber:  "+123 633963130",
					Status:       invalid,
				},
			},
		},
		expectErr: nil,
	}
}

func MakeScenarioWithoutParamsExpectDtoFilledWithElevenCustomersAndErrorNil() *testScenario {
	return &testScenario{
		testName:  "Get 10 customers with received limit and offset, without extra params",
		limit:     defaultLimit,
		offset:    defaultOffset,
		params:    map[string]string{},
		mockTotal: int64(11),
		mockResult: []model.Customer{
			{
				Id:    1,
				Name:  "John Doe",
				Phone: "(212) 6007989253",
			},
			{
				Id:    2,
				Name:  "James Smith",
				Phone: "(212) 633963130",
			},
			{
				Id:    3,
				Name:  "Robert Jones",
				Phone: "(258) 847651504",
			},
			{
				Id:    4,
				Name:  "Michael Taylor",
				Phone: "(258) 84330678235",
			},
			{
				Id:    5,
				Name:  "William Williams",
				Phone: "(256) 775069443",
			},
			{
				Id:    6,
				Name:  "Mary Brown",
				Phone: "(256) 3142345678",
			},
			{
				Id:    7,
				Name:  "Patricia White",
				Phone: "(251) 9773199405",
			},
			{
				Id:    8,
				Name:  "Jennifer Harris",
				Phone: "(251) 914701723",
			},
			{
				Id:    9,
				Name:  "Linda Martin",
				Phone: "(237) 697151594",
			},
			{
				Id:    10,
				Name:  "Elizabeth Davies",
				Phone: "(237) 6780009592",
			},
		},
		mockErr: nil,
		expectResult: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  11,
			Customers: []dto.Customer{
				{
					CustomerName: "John Doe",
					CountryName:  "Morocco",
					CountryCode:  entity.CodeMorocco,
					PhoneNumber:  "6007989253",
					Status:       invalid,
				},
				{
					CustomerName: "James Smith",
					CountryName:  "Morocco",
					CountryCode:  entity.CodeMorocco,
					PhoneNumber:  "633963130",
					Status:       valid,
				},
				{
					CustomerName: "Robert Jones",
					CountryName:  "Mozambique",
					CountryCode:  entity.CodeMozambique,
					PhoneNumber:  "847651504",
					Status:       valid,
				},
				{
					CustomerName: "Michael Taylor",
					CountryName:  "Mozambique",
					CountryCode:  entity.CodeMozambique,
					PhoneNumber:  "84330678235",
					Status:       invalid,
				},
				{
					CustomerName: "William Williams",
					CountryName:  "Uganda",
					CountryCode:  entity.CodeUganda,
					PhoneNumber:  "775069443",
					Status:       valid,
				},
				{
					CustomerName: "Mary Brown",
					CountryName:  "Uganda",
					CountryCode:  entity.CodeUganda,
					PhoneNumber:  "3142345678",
					Status:       invalid,
				},
				{
					CustomerName: "Patricia White",
					CountryName:  "Ethiopia",
					CountryCode:  entity.CodeEthiopia,
					PhoneNumber:  "9773199405",
					Status:       invalid,
				},
				{
					CustomerName: "Jennifer Harris",
					CountryName:  "Ethiopia",
					CountryCode:  entity.CodeEthiopia,
					PhoneNumber:  "914701723",
					Status:       valid,
				},
				{
					CustomerName: "Linda Martin",
					CountryName:  "Cameroon",
					CountryCode:  entity.CodeCameroon,
					PhoneNumber:  "697151594",
					Status:       valid,
				},
				{
					CustomerName: "Elizabeth Davies",
					CountryName:  "Cameroon",
					CountryCode:  entity.CodeCameroon,
					PhoneNumber:  "6780009592",
					Status:       invalid,
				},
			},
		},
		expectErr: nil,
	}
}
