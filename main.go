package main

import "fmt"

var (
	accessNumber     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	accessOperations = []string{"+", "-", "/", "*", "^", "(", ")"}
)

func stringInString(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getOperationPath(list []string) []int64 {
	var (
		arrayIdx       []int64
		firstPriority  = []string{"^"}
		secondPriority = []string{"*", "/"}
		thirdPriority  = []string{"+", "-"}
	)

	for idx, data := range list {
		if stringInString(data, firstPriority) {
			arrayIdx = append(arrayIdx, int64(idx))
		}
	}

	for idx, data := range list {
		if stringInString(data, secondPriority) {
			arrayIdx = append(arrayIdx, int64(idx))
		}
	}

	for idx, data := range list {
		if stringInString(data, thirdPriority) {
			arrayIdx = append(arrayIdx, int64(idx))
		}
	}

	return arrayIdx
}

func clearInput(data string) []string {
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
	fmt.Println(clearInput("12 + 4 - 5a * 4"))
}

// +, -, /, *, ^, (, )

/*
	3 + 4 / 3 - 18 ^ ( 2 * 5 )
	["3", "+", "4", "/", "3", "-", "18", "^", "(", "2", "*", "5", ")"]
	(3 * 2) + (3 / 3) + 4
              +
         /         \
        +         4
     /     \
     *     /
	/ \   / \
	3 2   3 3
+
|--> +
|    |--> *
|	 |    |
|    |    |--> 3
|    |	  |
|	 |    |--> 2
|  	 |
|	 |--> /
|		  |--> 3
|		  |
|		  |--> 3
|
|--> 4

*/
