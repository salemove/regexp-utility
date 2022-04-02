package main

import (
	"flag"
	"fmt"
	"os"
	"regexp_utility/cmd"
)

type stringArrayFlags []string

func (i *stringArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// needed by flag
func (i *stringArrayFlags) String() string {
	return ""
}

func mask(input *string, regexpList *[]string) {
	result := cmd.Mask(input, regexpList)
	fmt.Print(result)
}

func validate(regularExpression *string) {
	isValid := cmd.Validate(regularExpression)

	if !isValid {
		fmt.Print("Invalid regular expression")
		os.Exit(1)
	}
}

func main() {
	// mask command flags
	maskCmd := flag.NewFlagSet("mask", flag.ExitOnError)
	maskInput := maskCmd.String("input", "", "string to mask")

	var regexpList stringArrayFlags
	maskCmd.Var(&regexpList, "regexp", "Regular expression to match against, multiple --regexp flags are allowed")

	// validate command flags
	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	validateRegexp := validateCmd.String("regexp", "", "Regular expression to validate")

	unexpectedCommandError := "Usage: regexp_utility <subcommand> --help\n\tavailable subcommands: mask, validate"

	if len(os.Args) < 2 {
		fmt.Println(unexpectedCommandError)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "mask":
		maskCmd.Parse(os.Args[2:])
		regexpStringList := []string(regexpList)

		if *maskInput == "" {
			fmt.Println("Input cannot be empty")
			os.Exit(1)
		}

		if len(regexpStringList) == 0 {
			fmt.Println("At least one regular expression has to be specified")
			os.Exit(1)
		}

		mask(maskInput, &regexpStringList)
	case "validate":
		validateCmd.Parse(os.Args[2:])

		if *validateRegexp == "" {
			fmt.Println("Regular expression cannot be empty")
			os.Exit(1)
		}

		validate(validateRegexp)
	default:
		fmt.Println(unexpectedCommandError)
		os.Exit(1)
	}
}
