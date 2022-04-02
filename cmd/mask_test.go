package cmd

import (
	"os"
	"os/exec"
	"testing"
)

func TestMaskWithSingleRegexp(t *testing.T) {
	input := "foo bar baz"
	regexpList := []string{"bar"}

	expectedResult := "foo *** baz"

	result := Mask(&input, &regexpList)

	assertEquals(t, result, expectedResult)
}

func TestMaskRespectsCaseRegexp(t *testing.T) {
	input := "foo bar baz BAR foo"
	regexpList := []string{"bar"}

	expectedResult := "foo *** baz BAR foo"

	result := Mask(&input, &regexpList)

	assertEquals(t, result, expectedResult)
}

func TestMaskWithMultipleRegexp(t *testing.T) {
	input := "My secret is abcdefg and my number is 369!"
	regexpList := []string{"abcdefg", "([(0-9)]{3})"}

	expectedResult := "My secret is ******* and my number is ***!"

	result := Mask(&input, &regexpList)

	assertEquals(t, result, expectedResult)
}

func TestMaskWithMultipleRegexpAndMultipleOccurrences(t *testing.T) {
	input := "My secret is \"abc\" and my number is 777 and code is defgh!"
	regexpList := []string{"(abc|defgh)", "([(0-9)]{3})"}

	expectedResult := "My secret is \"***\" and my number is *** and code is *****!"

	result := Mask(&input, &regexpList)

	assertEquals(t, result, expectedResult)
}

func TestRegexWithLookaround(t *testing.T) {
	if os.Getenv("RUN_TEST") == "1" {
		input := "Hello"
		// SSN regex that uses features not supported by RE2
		regexpList := []string{"^((?!666|000)[0-8][0-9\\_]{2}\\-(?!00)[0-9\\_]{2}\\-(?!0000)[0-9\\_]{4})*$"}
		Mask(&input, &regexpList)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestRegexWithLookaround")
	cmd.Env = append(os.Environ(), "RUN_TEST=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("process exited with exit code %v, expected exit status 1", err)
}
