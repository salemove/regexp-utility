package cmd

import (
	"regexp"
)

func Validate(regularExpression *string) bool {
	_, err := regexp.Compile(*regularExpression)

	if err != nil {
		return false
	}

	return true
}
