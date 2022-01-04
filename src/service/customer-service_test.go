package service

import (
	"github.com/golang/mock/gomock"
	mock_repository "github.com/nuno/nunes-jumia/src/repository/mocks"
	"reflect"
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

			if !reflect.DeepEqual(scenario.expectResult, result) {
				t.Errorf("Test result is '%v' but was expected '%v'", result, scenario.expectResult)
			}
			if !reflect.DeepEqual(scenario.expectErr, err) {
				t.Errorf("Test result is '%v' but was expected '%s'", err, scenario.expectErr)
			}
			if (len(result.Customers) < scenario.limit && len(result.Customers) != int(scenario.mockTotal)) ||
				(len(result.Customers) >= scenario.limit && len(result.Customers) != scenario.limit) {
				t.Errorf("Test result is '%v' but was expected '%s'", scenario.mockTotal, result.Customers)
			}

			ctrl.Finish()
		})
	}
}
