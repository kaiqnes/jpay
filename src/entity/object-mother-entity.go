package entity

const (
	valid   = "Valid"
	invalid = "Invalid"
)

type testScenario struct {
	TestName     string
	Phone        string
	CountryCode  string
	ExpectResult string
}

func MakeScenarioWithValidPhoneFromCameroon() *testScenario {
	return &testScenario{
		TestName:     "Valid Phone from Cameroon",
		Phone:        "(237) 673122155",
		CountryCode:  CodeCameroon,
		ExpectResult: valid,
	}
}

func MakeScenarioWithInvalidPhoneFromCameroon() *testScenario {
	return &testScenario{
		TestName:     "Invalid Phone from Cameroon",
		Phone:        "(237) 6622284920",
		CountryCode:  CodeCameroon,
		ExpectResult: invalid,
	}
}

func MakeScenarioWithValidPhoneFromEthiopia() *testScenario {
	return &testScenario{
		TestName:     "Valid Phone from Ethiopia",
		Phone:        "(251) 911168450",
		CountryCode:  CodeEthiopia,
		ExpectResult: valid,
	}
}

func MakeScenarioWithInvalidPhoneFromEthiopia() *testScenario {
	return &testScenario{
		TestName:     "Invalid Phone from Ethiopia",
		Phone:        "(251) 9119454961",
		CountryCode:  CodeEthiopia,
		ExpectResult: invalid,
	}
}

func MakeScenarioWithValidPhoneFromMorocco() *testScenario {
	return &testScenario{
		TestName:     "Valid Phone from Morocco",
		Phone:        "(212) 633963130",
		CountryCode:  CodeMorocco,
		ExpectResult: valid,
	}
}

func MakeScenarioWithInvalidPhoneFromMorocco() *testScenario {
	return &testScenario{
		TestName:     "Invalid Phone from Morocco",
		Phone:        "(212) 6007989253",
		CountryCode:  CodeMorocco,
		ExpectResult: invalid,
	}
}

func MakeScenarioWithValidPhoneFromMozambique() *testScenario {
	return &testScenario{
		TestName:     "Valid Phone from Mozambique",
		Phone:        "(258) 823747618",
		CountryCode:  CodeMozambique,
		ExpectResult: valid,
	}
}

func MakeScenarioWithInvalidPhoneFromMozambique() *testScenario {
	return &testScenario{
		TestName:     "Invalid Phone from Mozambique",
		Phone:        "(258) 042423566",
		CountryCode:  CodeMozambique,
		ExpectResult: invalid,
	}
}

func MakeScenarioWithValidPhoneFromUganda() *testScenario {
	return &testScenario{
		TestName:     "Valid Phone from Uganda",
		Phone:        "(256) 704244430",
		CountryCode:  CodeUganda,
		ExpectResult: valid,
	}
}

func MakeScenarioWithInvalidPhoneFromUganda() *testScenario {
	return &testScenario{
		TestName:     "Invalid Phone from Uganda",
		Phone:        "(256) 3142345678",
		CountryCode:  CodeUganda,
		ExpectResult: invalid,
	}
}
