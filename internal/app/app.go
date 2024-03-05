package app

import (
	"abstract-syntax-tree/internal/calculator"
	"abstract-syntax-tree/internal/tree"
	"abstract-syntax-tree/pkg"
	"fmt"
	"strconv"
)

// Проверка что строка это число
func checkNumberInString(number string) bool {
	_, err := strconv.ParseFloat(number, 64)
	if err == nil {
		return true
	}
	return false
}

// Преобразование строки в массив операторов и чисел
func cleanString(inputString string) []string {
	var (
		dataNumAndOpera  []string
		accessNumber     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "."}
		accessOperations = []string{"+", "-", "/", "*", "^", "(", ")"}
	)

	for _, stringRune := range inputString {
		if pkg.ExistsInSlice(string(stringRune), accessNumber) {
			if len(dataNumAndOpera) != 0 && checkNumberInString(dataNumAndOpera[len(dataNumAndOpera)-1]) {
				dataNumAndOpera[len(dataNumAndOpera)-1] += string(stringRune)
			} else {
				dataNumAndOpera = append(dataNumAndOpera, string(stringRune))
			}
		}
		if pkg.ExistsInSlice(string(stringRune), accessOperations) {
			dataNumAndOpera = append(dataNumAndOpera, string(stringRune))
		}
	}

	return dataNumAndOpera
}

func Start() {
	inputString := "1 * 4 + 2 + 4 + 3 + 3 * 4"
	//inputString := "3 + 4 + 5"
	cleanInput := cleanString(inputString)
	tree.Print(cleanInput)
	fmt.Println("Answer: ", calculator.Сalculate(cleanInput))
}
