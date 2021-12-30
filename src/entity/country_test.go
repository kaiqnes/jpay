package entity

import "testing"

type testScenario struct {
	testName    string
	phone       string
	countryCode string
	expect      string
}

func TestCountry(t *testing.T) {
	t.Parallel()

	scenarios := []testScenario{
		{testName: "Valid phone from Morocco", phone: "(212) 633963130", countryCode: "212", expect: "Valid"},
		{testName: "Invalid phone from Morocco", phone: "(212) 6007989253", countryCode: "212", expect: "Invalid"},
		{testName: "Valid phone from Cameroon", phone: "(237) 673122155", countryCode: "237", expect: "Valid"},
		{testName: "Invalid phone from Cameroon", phone: "(237) 6622284920", countryCode: "237", expect: "Invalid"},
		{testName: "Valid phone from Ethiopia", phone: "(251) 911168450", countryCode: "251", expect: "Valid"},
		{testName: "Invalid phone from Ethiopia", phone: "(251) 9119454961", countryCode: "251", expect: "Invalid"},
		{testName: "Valid phone from Uganda", phone: "(256) 704244430", countryCode: "256", expect: "Valid"},
		{testName: "Invalid phone from Uganda", phone: "(256) 3142345678", countryCode: "256", expect: "Invalid"},
		{testName: "Valid phone from Mozambique", phone: "(258) 823747618", countryCode: "258", expect: "Valid"},
		{testName: "Invalid phone from Mozambique", phone: "(258) 042423566", countryCode: "258", expect: "Invalid"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.testName, func(t *testing.T) {
			currentCountry := Countries[scenario.countryCode]
			result := currentCountry.IsValidPhoneNumber(scenario.phone)
			if result != scenario.expect {
				t.Errorf("Received %s but was expected %s", result, scenario.expect)
			}
		})
	}
}
