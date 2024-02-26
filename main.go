package main

import (
	"fmt"
)

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

// Функция для сортировки массива по первому элементу массива матрицы
// Данную функцию можно поменять или улучшить по произодитльности и скорости
func sortedByFirstIdx(data [][]int64) [][]int64 {
	for i := 0; i < len(data); i++ {
		for j := 0; j < (len(data) - i); j++ {
			if j != 0 {
				if data[j-1][0] > data[j][0] {
					data[j-1], data[j] = data[j], data[j-1]
				}
			}
		}
	}
	return data
}

// Генерация списка последовательности выполнения операций, где элемен это операция,
// а его индекс это приоретет выполения
func getOperationPath(list []string) []int64 {
	var arrayIdx []int64
	raiseFlag := 0
	var arrayFilter [][]int64

	for idx, data := range list {
		switch data {
		case "(":
			raiseFlag += 4
		case ")":
			raiseFlag -= 4
		case "^":
			arrayFilter = append(arrayFilter, []int64{int64(3 + raiseFlag), int64(idx)})
		case "*":
			arrayFilter = append(arrayFilter, []int64{int64(2 + raiseFlag), int64(idx)})
		case "/":
			arrayFilter = append(arrayFilter, []int64{int64(2 + raiseFlag), int64(idx)})
		case "+":
			arrayFilter = append(arrayFilter, []int64{int64(1 + raiseFlag), int64(idx)})
		case "-":
			arrayFilter = append(arrayFilter, []int64{int64(1 + raiseFlag), int64(idx)})
		}
	}

	arrayFilter = sortedByFirstIdx(arrayFilter)

	for _, data := range arrayFilter {
		arrayIdx = append(arrayIdx, data[1])
	}

	return arrayIdx
}

// Объедениение 2 одномерных массивов по горизонтали
func sumStringSlice(first, second []string) []string {
	sizeNewSlice := len(first) + len(second)
	newSlice := make([]string, 0, sizeNewSlice)

	for _, data := range first {
		newSlice = append(newSlice, data)
	}

	for _, data := range second {
		newSlice = append(newSlice, data)
	}

	return newSlice
}

// Совмещение трех блоков
func sumParts(mainLine []string, leftPart, rightPart [][]string) [][]string {
	var newArr [][]string

	if len(leftPart) < len(rightPart) {
		leftPart, rightPart = rightPart, leftPart
	}

	lenFakeSlice := len(rightPart[0])
	fakeSlice := make([]string, 0, lenFakeSlice)

	for i := 0; i < lenFakeSlice; i++ {
		fakeSlice = append(fakeSlice, " ")
	}

	for idxY, _ := range leftPart {
		lenLineNewArr := len(rightPart[0]) + len(rightPart[0])
		lineNewArr := make([]string, 0, lenLineNewArr)

		for _, data := range leftPart[idxY] {
			lineNewArr = append(lineNewArr, data)
		}

		if idxY >= (len(leftPart) - len(rightPart)) {
			for _, data := range rightPart[((len(leftPart) - len(rightPart)) - idxY)] {
				lineNewArr = append(lineNewArr, data)
			}
		} else {
			for _, data := range fakeSlice {
				lineNewArr = append(lineNewArr, data)
			}
		}
		newArr = append(newArr, lineNewArr)
	}
	newArr = append(newArr, mainLine)

	return newArr
}

func deleteBracket(data []string) []string {
	var newData []string
	for _, word := range data {
		if !(stringInString(word, []string{"(", ")"})) {
			newData = append(newData, word)
		}
	}
	return newData
}

func tree(arrayNumAndOpr []string) [][]string {
	if len(arrayNumAndOpr) == 0 {
		return [][]string{}
	}
	if len(arrayNumAndOpr) == 1 {
		return [][]string{
			{arrayNumAndOpr[0], " "},
		}
	}
	var (
		arrayIdxPriority        = getOperationPath(arrayNumAndOpr)
		lastIdxArrayIdxPriority = arrayIdxPriority[(len(arrayIdxPriority) - 1)]
	)
	leftPart, rightPart := arrayNumAndOpr[:(lastIdxArrayIdxPriority)], arrayNumAndOpr[(lastIdxArrayIdxPriority+1):]

	cleanLeftPart, cleanRightPart := deleteBracket(leftPart), deleteBracket(rightPart)

	if len(cleanLeftPart) == 1 && len(cleanRightPart) == 1 {
		return [][]string{
			{" ", " ", cleanLeftPart[0], " ", cleanRightPart[0], " "},
			{" ", " ", " ", " ", " ", " "},
			{" ", " ", ">", " ", ">", " "},
			{" ", " ", "-", " ", "-", " "},
			{" ", " ", "-", " ", "-", " "},
			{arrayNumAndOpr[lastIdxArrayIdxPriority], "|", "|", "|", "|", " "},
		}
	}

	leftPartAnswer, rightPartAnswer := elementTree(leftPart), elementTree(rightPart)

	var lineOpr []string
	for i := 0; i < len(leftPartAnswer[0])+len(rightPartAnswer[0]); i++ {
		switch i {
		case 0:
			lineOpr = append(lineOpr, arrayNumAndOpr[lastIdxArrayIdxPriority])
		case len(leftPartAnswer[0]) + len(rightPartAnswer[0]) - 1:
			lineOpr = append(lineOpr, " ")
		default:
			lineOpr = append(lineOpr, "|")

		}
	}
	answer := sumParts(lineOpr, leftPartAnswer, rightPartAnswer)

	return answer
}

func printTree(data string) {

}
func main() {
	fmt.Println(tree(clearInput("1+3-4")))
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
//func checkNumber(num string, list []string) bool {
//	for _, elementNum := range num {
//		for _, elementList := range list {
//			if string(elementNum) != elementList {
//				return false
//			}
//		}
//	}
//	return true
//}

// Удалить если получиться код ниже
//
//func transposition(data [][]string) [][]string {
//	var newData [][]string
//
//	for i := range data {
//		copy(newData[i][:], newData[i][:])
//	}
//
//	for y := range newData {
//		for x := range newData {
//			data[y][x] = newData[x][y]
//		}
//	}
//	return data
//}
//
//// First array small, second big
//func sumSmallAndBigSlice(smallArr, bigArr [][]string) [][]string {
//	if len(smallArr[0]) > len(bigArr[0]) {
//		// Отладка
//	}
//
//	var newArr [][]string
//	for _, data1 := range smallArr {
//		var newLineSmallArr []string
//		var flag = len(bigArr[0])
//		for _, data2 := range data1 {
//			newLineSmallArr = append(newLineSmallArr, data2)
//			flag--
//		}
//		for i := 0; i < flag; i++ {
//			newLineSmallArr = append(newLineSmallArr, " ")
//		}
//		newArr = append(newArr, newLineSmallArr)
//	}
//	for _, data1 := range bigArr {
//		newArr = append(newArr, data1)
//	}
//	return newArr
//}
