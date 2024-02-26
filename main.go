package main

import "fmt"

func stringInString(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func clearInput(data string) []string {
	var (
		accessNumber     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
		accessOperations = []string{"+", "-", "/", "*", "^", "(", ")"}
	)
	var (
		dataNumberAndOperations []string
		flag                    bool
	)

	flag = false
	for _, word := range data {
		if stringInString(string(word), accessNumber) || stringInString(string(word), accessOperations) {
			flag = flag && !(stringInString(string(word), accessOperations))
			fmt.Println(flag, string(word))
			if flag {
				dataNumberAndOperations[len(dataNumberAndOperations)-1] += string(word)
			} else {
				dataNumberAndOperations = append(dataNumberAndOperations, string(word))
			}
			flag = stringInString(dataNumberAndOperations[len(dataNumberAndOperations)-1], accessNumber)
		}
	}

	return dataNumberAndOperations
}

func main() {
	fmt.Println(clearInput("12 + 4 - 5a"))
}

// +, -, /, *, ^, (, )

/*
	3 + 4 / 3 - 18 ^ ( 2 * 5 )
	["3", "+", "4", "/", "3", "-", "18", "^", "(", "2", "*", "5", ")"]


	  *
	/   \ /   \
	2	5 	5
*/
