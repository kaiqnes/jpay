package repository

import (
	"github.com/nuno/nunes-jumia/src/model"
)

type testScenario struct {
	TestName           string
	Rows               []model.Customer
	ExpectLengthResult int
	ExpectError        error
}

func MakeScenarioReturnsTwoCustomers() *testScenario {
	return &testScenario{
		TestName: "Retrieve two customers",
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
		ExpectLengthResult: 2,
		ExpectError:        nil,
	}
}

func MakeScenarioReturnsNoneCustomer() *testScenario {
	return &testScenario{
		TestName:           "Send limit and offset default to retrieve none customer",
		Rows:               []model.Customer{},
		ExpectLengthResult: 0,
		ExpectError:        nil,
	}
}
