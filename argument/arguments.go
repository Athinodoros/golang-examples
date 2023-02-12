package argument

import (
	"os"
)

func CleanArguments(args []string) map[string]string {
	args = args[1:]
	if len(args)%2 == 1 {
		os.Exit(5)
	}
	argumentSets := make(map[string]string)
	for i := 0; i < len(args); i = i + 2 {
		argumentSets[args[i]] = args[i+1]
	}

	return argumentSets

}

func GetArgumentsInGivenOrder(argumentMap map[string]string, orderedList []string) map[int][2]string {
	orderedArguments := make(map[int][2]string)

	for index, argumentFlag := range orderedList {
		if argumentMap[argumentFlag] != "" {
			orderedArguments[index] = [2]string{argumentFlag, argumentMap[argumentFlag]}
		}
	}

	return orderedArguments
}
