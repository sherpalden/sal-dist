package goCli

import (
	"flag"
	"fmt"
)

type ArgConfig struct {
	Name     string
	Required bool
	Validate func(string) error
}

func GetArgs(argsConfig []ArgConfig) (map[string]interface{}, error) {
	argsMap := make(map[string]interface{})
	flags := make(map[string]*string)

	// Create a flag for each argument
	for _, arg := range argsConfig {
		flags[arg.Name] = flag.String(arg.Name, "", fmt.Sprintf("description for %s", arg.Name))
	}

	// Parse the flags
	flag.Parse()

	// Validate and store the results
	for _, arg := range argsConfig {
		value := *flags[arg.Name]
		if arg.Required && value == "" {
			return nil, fmt.Errorf("required argument '%s' is missing", arg.Name)
		}
		if arg.Validate != nil {
			if err := arg.Validate(value); err != nil {
				return nil, fmt.Errorf("validation failed for argument '%s': %v", arg.Name, err)
			}
		}
		argsMap[arg.Name] = value
	}

	return argsMap, nil
}
