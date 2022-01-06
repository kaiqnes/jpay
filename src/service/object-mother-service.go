package service

import (
	"errors"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/entity"
	"github.com/nuno/nunes-jumia/src/model"
)

const (
	valid   = "Valid"
	invalid = "Invalid"
)

type testScenario struct {
	TestName     string
	MockResult   []model.Customer
	MockErr      error
	ExpectResult dto.CustomerOutputDto
	ExpectErr    error
}

func MakeScenarioExpectDtoWithSingleValidCustomer() *testScenario {
	return &testScenario{
		TestName: "Get a customer",
		MockResult: []model.Customer{
			{
				Id:    2,
				Name:  "James Smith",
				Phone: "(212) 633963130",
			},
		},
		MockErr: nil,
		ExpectResult: dto.CustomerOutputDto{
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
		ExpectErr: nil,
	}
}

func MakeScenarioExpectDtoWithTenCustomers() *testScenario {
	return &testScenario{
		TestName: "Get 10 customers with received Limit and Offset",
		MockResult: []model.Customer{
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
		MockErr: nil,
		ExpectResult: dto.CustomerOutputDto{
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
				{
					CustomerName: "Robert Jones",
					CountryName:  entity.NameMozambique,
					CountryCode:  entity.CodeMozambique,
					PhoneNumber:  "847651504",
					Status:       valid,
				},
				{
					CustomerName: "Michael Taylor",
					CountryName:  entity.NameMozambique,
					CountryCode:  entity.CodeMozambique,
					PhoneNumber:  "84330678235",
					Status:       invalid,
				},
				{
					CustomerName: "William Williams",
					CountryName:  entity.NameUganda,
					CountryCode:  entity.CodeUganda,
					PhoneNumber:  "775069443",
					Status:       valid,
				},
				{
					CustomerName: "Mary Brown",
					CountryName:  entity.NameUganda,
					CountryCode:  entity.CodeUganda,
					PhoneNumber:  "3142345678",
					Status:       invalid,
				},
				{
					CustomerName: "Patricia White",
					CountryName:  entity.NameEthiopia,
					CountryCode:  entity.CodeEthiopia,
					PhoneNumber:  "9773199405",
					Status:       invalid,
				},
				{
					CustomerName: "Jennifer Harris",
					CountryName:  entity.NameEthiopia,
					CountryCode:  entity.CodeEthiopia,
					PhoneNumber:  "914701723",
					Status:       valid,
				},
				{
					CustomerName: "Linda Martin",
					CountryName:  entity.NameCameroon,
					CountryCode:  entity.CodeCameroon,
					PhoneNumber:  "697151594",
					Status:       valid,
				},
				{
					CustomerName: "Elizabeth Davies",
					CountryName:  entity.NameCameroon,
					CountryCode:  entity.CodeCameroon,
					PhoneNumber:  "6780009592",
					Status:       invalid,
				},
			},
		},
		ExpectErr: nil,
	}
}

func MakeScenarioExpectDtoEmptyAndError() *testScenario {
	return &testScenario{
		TestName:     "Receive error from repository layer",
		MockResult:   []model.Customer{},
		MockErr:      errors.New("mock repository error"),
		ExpectResult: dto.CustomerOutputDto{},
		ExpectErr:    errors.New("Fail to retrieve customers in DB. Err: mock repository error"),
	}
}
