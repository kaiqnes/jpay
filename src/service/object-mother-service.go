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
	TestName     string
	Limit        int
	Offset       int
	Params       map[string]string
	MockResult   []model.Customer
	MockErr      error
	ExpectResult dto.CustomerOutputDto
	ExpectErr    error
}

func MakeScenarioWithoutParamsExpectDtoFilledWithSingleCustomerAndErrorNil() *testScenario {
	return &testScenario{
		TestName: "Get a customer with received Limit and Offset, without extra Params",
		Limit:    defaultLimit,
		Offset:   defaultOffset,
		Params:   map[string]string{},
		MockResult: []model.Customer{
			{
				Id:    1,
				Name:  "John Doe",
				Phone: "(212) 633963130",
			},
		},
		MockErr: nil,
		ExpectResult: dto.CustomerOutputDto{
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  1,
			Customers: []dto.Customer{
				{
					CustomerName: "John Doe",
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

func MakeScenarioWithoutParamsExpectDtoEmptyAndError() *testScenario {
	return &testScenario{
		TestName:     "Get a customer with received Limit and Offset, without extra Params and returns error",
		Limit:        defaultLimit,
		Offset:       defaultOffset,
		Params:       map[string]string{},
		MockResult:   []model.Customer{},
		MockErr:      mockRepositoryError,
		ExpectResult: dto.CustomerOutputDto{},
		ExpectErr:    errors.New("Fail to retrieve customers in DB. Err: mock repository error"),
	}
}

func MakeScenarioWithoutParamsExpectDtoFilledWithInvalidCustomerAndErrorNil() *testScenario {
	return &testScenario{
		TestName: "Get a customer without received Limit, Offset and extra Params",
		Limit:    defaultLimit,
		Offset:   defaultOffset,
		Params:   map[string]string{},
		MockResult: []model.Customer{
			{
				Id:    1,
				Name:  "John Doe",
				Phone: "+123 633963130",
			},
		},
		MockErr: nil,
		ExpectResult: dto.CustomerOutputDto{
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
		ExpectErr: nil,
	}
}

func MakeScenarioWithoutParamsExpectDtoFilledWithTenCustomersAndErrorNil() *testScenario {
	return &testScenario{
		TestName: "Get 10 customers with received Limit and Offset, without extra Params",
		Limit:    defaultLimit,
		Offset:   defaultOffset,
		Params:   map[string]string{},
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
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  10,
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

func MakeScenarioFilteringByCountryNameExpectDtoFilledWithCustomers() *testScenario {
	return &testScenario{
		TestName: "Get 2 customers filtered by country_name morocco with received Limit and Offset",
		Limit:    defaultLimit,
		Offset:   defaultOffset,
		Params:   map[string]string{"country_name": entity.NameMorocco},
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
		ExpectErr: nil,
	}
}

func MakeScenarioFilteringByStatusExpectDtoFilledWithCustomers() *testScenario {
	return &testScenario{
		TestName: "Get 5 customers filtered by status valid with received Limit and Offset",
		Limit:    defaultLimit,
		Offset:   defaultOffset,
		Params:   map[string]string{"status": valid},
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
			Limit:  defaultLimit,
			Offset: defaultOffset,
			Total:  5,
			Customers: []dto.Customer{
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
					CustomerName: "William Williams",
					CountryName:  entity.NameUganda,
					CountryCode:  entity.CodeUganda,
					PhoneNumber:  "775069443",
					Status:       valid,
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
			},
		},
		ExpectErr: nil,
	}
}

func MakeScenarioFilteringByCountryNameAndStatusExpectDtoFilledWithCustomers() *testScenario {
	return &testScenario{
		TestName: "Get 1 customers filtered by country_name morocco and status valid with received Limit and Offset",
		Limit:    defaultLimit,
		Offset:   defaultOffset,
		Params:   map[string]string{"country_name": entity.NameMorocco, "status": valid},
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
			Limit:  defaultLimit,
			Offset: defaultOffset,
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
		ExpectErr: nil,
	}
}
