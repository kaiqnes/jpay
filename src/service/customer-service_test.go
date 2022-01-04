package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	mock_repository "github.com/nuno/nunes-jumia/src/repository/mocks"
	"testing"
)

func TestCustomerService(t *testing.T) {
	scenarios := []testScenario{
		*MakeScenarioWithoutParamsExpectDtoFilledWithSingleCustomerAndErrorNil(),
		*MakeScenarioWithoutParamsExpectDtoEmptyAndError(),
		*MakeScenarioWithoutParamsExpectDtoFilledWithInvalidCustomerAndErrorNil(),
		*MakeScenarioWithoutParamsExpectDtoFilledWithElevenCustomersAndErrorNil(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepository := mock_repository.NewMockCustomerRepository(ctrl)
			testCustomerService := NewCustomerService(mockRepository)

			mockRepository.EXPECT().GetCustomers(scenario.Limit, scenario.Offset).Return(scenario.MockTotal, scenario.MockResult, scenario.MockErr)

			result, err := testCustomerService.GetCustomers(scenario.Limit, scenario.Offset, scenario.Params)

			isCustomersLengthCorrect := !(len(result.Customers) < scenario.Limit && len(result.Customers) != int(scenario.MockTotal)) ||
				(len(result.Customers) >= scenario.Limit && len(result.Customers) != scenario.Limit)

			ctrl.Finish()

			assert.Equal(t, result, scenario.ExpectResult)
			assert.Equal(t, err, scenario.ExpectErr)
			assert.Equal(t, isCustomersLengthCorrect, true)
		})
	}
}
