package cmd

import (
	"testing"
)

func TestValidateWithValidRegexp(t *testing.T) {
	regularExpression := "foo"

	result := Validate(&regularExpression)
	expectedResult := true

	assertEquals(t, result, expectedResult)
}

func TestValidateWithRegexpWithLookaroundFeature(t *testing.T) {
	regularExpression := "^((?!666|000)[0-8][0-9\\_]{2}\\-(?!00)[0-9\\_]{2}\\-(?!0000)[0-9\\_]{4})*$"

	result := Validate(&regularExpression)
	expectedResult := false

	assertEquals(t, result, expectedResult)
}
