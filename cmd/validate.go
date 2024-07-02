package cmd

import (
	"regexp"
)

func Validate(regularExpression *string) bool {
	_, err := regexp.Compile(*regularExpression)

	return err == nil
}
