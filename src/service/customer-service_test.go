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
		t.Run(scenario.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepository := mock_repository.NewMockCustomerRepository(ctrl)
			testCustomerService := NewCustomerService(mockRepository)

			mockRepository.EXPECT().GetCustomers(scenario.limit, scenario.offset).Return(scenario.mockTotal, scenario.mockResult, scenario.mockErr)

			result, err := testCustomerService.GetCustomers(scenario.limit, scenario.offset, scenario.params)

			isCustomersLengthCorrect := !(len(result.Customers) < scenario.limit && len(result.Customers) != int(scenario.mockTotal)) ||
				(len(result.Customers) >= scenario.limit && len(result.Customers) != scenario.limit)

			ctrl.Finish()

			assert.Equal(t, result, scenario.expectResult)
			assert.Equal(t, err, scenario.expectErr)
			assert.Equal(t, isCustomersLengthCorrect, true)
		})
	}
}
