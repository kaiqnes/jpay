package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mock_service "github.com/nuno/nunes-jumia/src/service/mocks"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCustomerController(t *testing.T) {
	scenarios := []testScenario{
		*MakeScenarioWithoutParamsExpectCustomersDtoWithOneCustomer(),
		*MakeScenarioWithoutParamsExpectCustomersDtoWithTwoCustomers(),
		*MakeScenarioExpectErrorToExtractLimitQueryParam(),
		*MakeScenarioExpectErrorToExtractOffsetQueryParam(),
		*MakeScenarioExpectCustomersDtoWithLimit1(),
		*MakeScenarioExpectCustomersDtoWithOffset1(),
		*MakeScenarioExpectCustomersDtoWithCountryNameFilter(),
		*MakeScenarioExpectCustomersDtoWithStatusFilter(),
		*MakeScenarioExpectErrorInServiceLayer(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockService := mock_service.NewMockCustomerService(ctrl)
			testController := NewCustomerController(mockService)

			gin.SetMode("release")
			router := gin.New()
			testController.SetupRoutes(router)

			if scenario.ShouldMockServiceCall {
				mockService.EXPECT().GetCustomers(scenario.Limit, scenario.Offset, scenario.MockParams).
					Return(scenario.MockConsumerDto, scenario.MockErr)
			}

			response := httptest.NewRecorder()
			executeRequest(response, http.MethodGet, scenario.getFullUrl(), scenario.BodyString, router)

			if !reflect.DeepEqual(scenario.ExpectStatus, response.Result().StatusCode) {
				//t.Errorf("Test result is '%v' but was expected '%v'", response.Result().StatusCode, scenario.ExpectStatus)
				t.Errorf("\n\nresult: '%v'\nexpect: '%v'", response.Result().StatusCode, scenario.ExpectStatus)
			}

			if !reflect.DeepEqual(scenario.ExpectResponse, response.Body.String()) {
				//t.Errorf("Test result is '%v' but was expected '%v'", response.Body.String(), scenario.ExpectResponse)
				t.Errorf("\n\nresult: '%v'\nexpect: '%v'", response.Body.String(), scenario.ExpectResponse)
			}

			ctrl.Finish()
		})
	}
}

func executeRequest(response *httptest.ResponseRecorder, method, requestUrl, body string, router *gin.Engine) {
	req, _ := http.NewRequest(method, requestUrl, bytes.NewBuffer([]byte(body)))
	router.ServeHTTP(response, req)
}
