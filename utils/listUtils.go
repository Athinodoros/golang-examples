package utils

func ReverseArgsList(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(ReverseArgsList(input[1:]), input[0])
}
