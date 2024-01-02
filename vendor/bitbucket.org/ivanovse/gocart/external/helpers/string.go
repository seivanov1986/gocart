package helpers

import "regexp"

func CleanPhone(phone string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return phone
	}

	return reg.ReplaceAllString(phone, "")
}
