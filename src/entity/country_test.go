package entity

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCountryEntity(t *testing.T) {
	scenarios := []testScenario{
		*MakeScenarioWithValidPhoneFromCameroon(),
		*MakeScenarioWithInvalidPhoneFromCameroon(),
		*MakeScenarioWithValidPhoneFromEthiopia(),
		*MakeScenarioWithInvalidPhoneFromEthiopia(),
		*MakeScenarioWithValidPhoneFromMorocco(),
		*MakeScenarioWithInvalidPhoneFromMorocco(),
		*MakeScenarioWithValidPhoneFromMozambique(),
		*MakeScenarioWithInvalidPhoneFromMozambique(),
		*MakeScenarioWithValidPhoneFromUganda(),
		*MakeScenarioWithInvalidPhoneFromUganda(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			currentCountry := Countries[scenario.CountryCode]

			result := currentCountry.IsValidPhoneNumber(scenario.Phone)

			assert.Equal(t, result, scenario.ExpectResult)
		})
	}
}
