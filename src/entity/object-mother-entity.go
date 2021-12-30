package entity

const (
	valid   = "Valid"
	invalid = "Invalid"
)

type testScenario struct {
	testName    string
	phone       string
	countryCode string
	expect      string
}

func MakeEntityWithValidPhoneFromCameroon() *testScenario {
	return &testScenario{
		testName:    "Valid phone from Cameroon",
		phone:       "(237) 673122155",
		countryCode: CodeCameroon,
		expect:      valid,
	}
}

func MakeEntityWithInvalidPhoneFromCameroon() *testScenario {
	return &testScenario{
		testName:    "Invalid phone from Cameroon",
		phone:       "(237) 6622284920",
		countryCode: CodeCameroon,
		expect:      invalid,
	}
}

func MakeEntityWithValidPhoneFromEthiopia() *testScenario {
	return &testScenario{
		testName:    "Valid phone from Ethiopia",
		phone:       "(251) 911168450",
		countryCode: CodeEthiopia,
		expect:      valid,
	}
}

func MakeEntityWithInvalidPhoneFromEthiopia() *testScenario {
	return &testScenario{
		testName:    "Invalid phone from Ethiopia",
		phone:       "(251) 9119454961",
		countryCode: CodeEthiopia,
		expect:      invalid,
	}
}

func MakeEntityWithValidPhoneFromMorocco() *testScenario {
	return &testScenario{
		testName:    "Valid phone from Morocco",
		phone:       "(212) 633963130",
		countryCode: CodeMorocco,
		expect:      valid,
	}
}

func MakeEntityWithInvalidPhoneFromMorocco() *testScenario {
	return &testScenario{
		testName:    "Invalid phone from Morocco",
		phone:       "(212) 6007989253",
		countryCode: CodeMorocco,
		expect:      invalid,
	}
}

func MakeEntityWithValidPhoneFromMozambique() *testScenario {
	return &testScenario{
		testName:    "Valid phone from Mozambique",
		phone:       "(258) 823747618",
		countryCode: CodeMozambique,
		expect:      valid,
	}
}

func MakeEntityWithInvalidPhoneFromMozambique() *testScenario {
	return &testScenario{
		testName:    "Invalid phone from Mozambique",
		phone:       "(258) 042423566",
		countryCode: CodeMozambique,
		expect:      invalid,
	}
}

func MakeEntityWithValidPhoneFromUganda() *testScenario {
	return &testScenario{
		testName:    "Valid phone from Uganda",
		phone:       "(256) 704244430",
		countryCode: CodeUganda,
		expect:      valid,
	}
}

func MakeEntityWithInvalidPhoneFromUganda() *testScenario {
	return &testScenario{
		testName:    "Invalid phone from Uganda",
		phone:       "(256) 3142345678",
		countryCode: CodeUganda,
		expect:      invalid,
	}
}
