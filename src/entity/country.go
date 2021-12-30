package entity

import "regexp"

const (
	CodeCameroon   = "237"
	CodeEthiopia   = "251"
	CodeMorocco    = "212"
	CodeMozambique = "258"
	CodeUganda     = "256"
)

var (
	Countries = map[string]Country{
		CodeCameroon: {
			Name:  "Cameroon",
			Regex: "\\(237\\)\\ ?[2368]\\d{7,8}$",
		},
		CodeEthiopia: {
			Name:  "Ethiopia",
			Regex: "\\(251\\)\\ ?[1-59]\\d{8}$",
		},
		CodeMorocco: {
			Name:  "Morocco",
			Regex: "\\(212\\)\\ ?[5-9]\\d{8}$",
		},
		CodeMozambique: {
			Name:  "Mozambique",
			Regex: "\\(258\\)\\ ?[28]\\d{7,8}$",
		},
		CodeUganda: {
			Name:  "Uganda",
			Regex: "\\(256\\)\\ ?\\d{9}$",
		},
	}
)

type Country struct {
	Name  string
	Regex string
}

func (country *Country) IsValidPhoneNumber(phone string) string {
	regex, _ := regexp.Compile(country.Regex)

	if regex.MatchString(phone) {
		return "Valid"
	}

	return "Invalid"
}
