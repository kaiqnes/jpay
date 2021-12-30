package country

import "regexp"

var (
	Countries = map[string]Country{
		"237": {
			Name:  "Cameroon",
			Regex: "\\(237\\)\\ ?[2368]\\d{7,8}$",
		},
		"251": {
			Name:  "Ethiopia",
			Regex: "\\(251\\)\\ ?[1-59]\\d{8}$",
		},
		"212": {
			Name:  "Morocco",
			Regex: "\\(212\\)\\ ?[5-9]\\d{8}$",
		},
		"258": {
			Name:  "Mozambique",
			Regex: "\\(258\\)\\ ?[28]\\d{7,8}$",
		},
		"256": {
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
