package repository

import (
	"github.com/nuno/nunes-jumia/src/model"
)

const (
	defaultLimit  = 10
	defaultOffset = 0
)

type testScenario struct {
	TestName           string
	Rows               []model.Customer
	Limit              int
	Offset             int
	ExpectTotal        int64
	ExpectLengthResult int
	ExpectError        error
}

func MakeScenarioWithLimitOffsetDefaultReturnsTwoCustomers() *testScenario {
	return &testScenario{
		TestName: "Send default limit and offset to retrieve two customers",
		Rows: []model.Customer{
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
		},
		Limit:              defaultLimit,
		Offset:             defaultOffset,
		ExpectTotal:        2,
		ExpectLengthResult: 2,
		ExpectError:        nil,
	}
}

func MakeScenarioWithLimit1AndOffsetDefaultReturnsOneCustomer() *testScenario {
	return &testScenario{
		TestName: "Send limit 1 and default offset to retrieve one customer",
		Rows: []model.Customer{
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
		},
		Limit:              1,
		Offset:             defaultOffset,
		ExpectTotal:        2,
		ExpectLengthResult: 1,
		ExpectError:        nil,
	}
}

func MakeScenarioWithLimitDefaultAndOffset2ReturnsNoneCustomer() *testScenario {
	return &testScenario{
		TestName: "Send limit default and offset 2 to retrieve none customer",
		Rows: []model.Customer{
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
		},
		Limit:              defaultLimit,
		Offset:             2,
		ExpectTotal:        2,
		ExpectLengthResult: 0,
		ExpectError:        nil,
	}
}

func MakeScenarioWithLimitOffsetDefaultReturnsNoneCustomer() *testScenario {
	return &testScenario{
		TestName:           "Send limit and offset default to retrieve none customer",
		Rows:               []model.Customer{},
		Limit:              defaultLimit,
		Offset:             defaultOffset,
		ExpectTotal:        0,
		ExpectLengthResult: 0,
		ExpectError:        nil,
	}
}
