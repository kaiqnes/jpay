package entity

import "testing"

func TestCountry(t *testing.T) {
	t.Parallel()

	scenarios := []testScenario{
		*MakeEntityWithValidPhoneFromCameroon(),
		*MakeEntityWithInvalidPhoneFromCameroon(),
		*MakeEntityWithValidPhoneFromEthiopia(),
		*MakeEntityWithInvalidPhoneFromEthiopia(),
		*MakeEntityWithValidPhoneFromMorocco(),
		*MakeEntityWithInvalidPhoneFromMorocco(),
		*MakeEntityWithValidPhoneFromMozambique(),
		*MakeEntityWithInvalidPhoneFromMozambique(),
		*MakeEntityWithValidPhoneFromUganda(),
		*MakeEntityWithInvalidPhoneFromUganda(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.testName, func(t *testing.T) {
			currentCountry := Countries[scenario.countryCode]
			result := currentCountry.IsValidPhoneNumber(scenario.phone)
			if result != scenario.expect {
				t.Errorf("Test result is '%s' but was expected '%s'", result, scenario.expect)
			}
		})
	}
}
