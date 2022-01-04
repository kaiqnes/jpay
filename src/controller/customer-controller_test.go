package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	mock_service "github.com/nuno/nunes-jumia/src/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomerController(t *testing.T) {
	scenarios := []testScenario{
		*MakeScenarioExpectCustomersDtoWithOneCustomer(),
		*MakeScenarioExpectCustomersDtoWithTwoCustomers(),
		*MakeScenarioExpectErrorToExtractLimitQueryParam(),
		*MakeScenarioExpectErrorToExtractOffsetQueryParam(),
		*MakeScenarioExpectCustomersDtoWithLimit1(),
		*MakeScenarioExpectCustomersDtoWithOffset1(),
		*MakeScenarioExpectErrorInServiceLayer(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockService := mock_service.NewMockCustomerService(ctrl)
			testController := NewCustomerController(mockService)

			router := gin.New()
			testController.SetupRoutes(router)

			if scenario.ShouldMockServiceCall {
				mockService.EXPECT().GetCustomers(scenario.Limit, scenario.Offset).
					Return(scenario.MockConsumerDto, scenario.MockErr)
			}

			response := httptest.NewRecorder()
			executeRequest(response, http.MethodGet, scenario.getFullUrl(), scenario.BodyString, router)

			ctrl.Finish()

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}

func executeRequest(response *httptest.ResponseRecorder, method, requestUrl, body string, router *gin.Engine) {
	req, _ := http.NewRequest(method, requestUrl, bytes.NewBuffer([]byte(body)))
	router.ServeHTTP(response, req)
}
