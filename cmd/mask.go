package cmd

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"unicode/utf8"
)

func Mask(input *string, regexpList *[]string) string {
	inputBytes := []byte(*input)

	for _, singleRegexp := range *regexpList {
		compiledRegexp, err := regexp.Compile(singleRegexp)

		if err != nil {
			fmt.Printf("Failed to compile regexp '%s'", singleRegexp)
			os.Exit(1)
		}

		inputBytes = compiledRegexp.ReplaceAllFunc(inputBytes, replaceWithAsterisks)
	}

	return string(inputBytes)
}

func replaceWithAsterisks(match []byte) []byte {
	matchLength := utf8.RuneCount(match)
	replacement := bytes.Repeat([]byte("*"), matchLength)

	return replacement
}
