package main

import (
	"fmt"
	"math"
)

var (
	accessNumber     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	accessOperations = []string{"+", "-", "/", "*", "^", "(", ")"}
)

// Чистка от символов отличных от accessNumber и accessOperations и запись чисел и операций в массив
func cleanAndConvertString(inputData string) []string {
	var (
		dataNumberAndOperations []string
		flag                    bool
	)

	flag = false
	for _, word := range inputData {
		if existsInSlice(string(word), accessNumber) || existsInSlice(string(word), accessOperations) {
			flag = flag && !(existsInSlice(string(word), accessOperations))
			if flag {
				dataNumberAndOperations[len(dataNumberAndOperations)-1] += string(word)
			} else {
				dataNumberAndOperations = append(dataNumberAndOperations, string(word))
			}
			flag = existsInSlice(dataNumberAndOperations[len(dataNumberAndOperations)-1], accessNumber)
		}
	}

	return dataNumberAndOperations
}

// Функция для сортировки массива по первому элементу массива матрицы
// Данную функцию можно поменять или улучшить по произодитльности и скорости
// Пока что это сортировка пузырьком
func sortArrByFirstIdx(data [][]int64) [][]int64 {
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
	// Массив типа [[приорететность операции, индекс операции], [приорететность операции, индекс операции]]
	var arrayPriorityAndIdx [][]int64

	raiseFlag := 0
	for idx, data := range list {
		switch data {
		case "(":
			raiseFlag += 4
		case ")":
			raiseFlag -= 4
		case "^":
			arrayPriorityAndIdx = append(arrayPriorityAndIdx, []int64{int64(3 + raiseFlag), int64(idx)})
		case "*":
			arrayPriorityAndIdx = append(arrayPriorityAndIdx, []int64{int64(2 + raiseFlag), int64(idx)})
		case "/":
			arrayPriorityAndIdx = append(arrayPriorityAndIdx, []int64{int64(2 + raiseFlag), int64(idx)})
		case "+":
			arrayPriorityAndIdx = append(arrayPriorityAndIdx, []int64{int64(1 + raiseFlag), int64(idx)})
		case "-":
			arrayPriorityAndIdx = append(arrayPriorityAndIdx, []int64{int64(1 + raiseFlag), int64(idx)})
		}
	}
	arrayPriorityAndIdx = reverce2DArrayByVertical(arrayPriorityAndIdx)

	var arrayIdx []int64

	arrayPriorityAndIdx = sortArrByFirstIdx(arrayPriorityAndIdx)
	for _, data := range arrayPriorityAndIdx {
		arrayIdx = append(arrayIdx, data[1])
	}

	return arrayIdx
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

	for idxY := range leftPart {
		lenLineNewArr := len(rightPart[0]) + len(leftPart[0])
		lineNewArr := make([]string, 0, lenLineNewArr)

		for _, data := range leftPart[idxY] {
			lineNewArr = append(lineNewArr, data)
		}

		idxRightWithLeftRight := len(leftPart) - len(rightPart)
		idxRight := int64(math.Abs(float64((len(leftPart) - idxY) - len(rightPart))))

		if idxY >= idxRightWithLeftRight {
			for _, data := range rightPart[idxRight] {
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

// Удаление из массива скобок
func deleteBracket(data []string) []string {
	var newData []string
	for _, word := range data {
		if !(existsInSlice(word, []string{"(", ")"})) {
			newData = append(newData, word)
		}
	}
	return newData
}

func createBaseArrTree(arrayNumAndOpr []string) [][]string {
	if len(arrayNumAndOpr) == 0 {
		return [][]string{}
	}
	if len(deleteBracket(arrayNumAndOpr)) == 1 {
		return [][]string{
			{arrayNumAndOpr[0], " "},
			{">", " "},
			{"-", " "},
			{"-", " "},
		}
	}

	var (
		arrayIdxPriority = getOperationPath(arrayNumAndOpr)
		IdxPriority      = arrayIdxPriority[0]
	)
	leftPart, rightPart := arrayNumAndOpr[:(IdxPriority)], arrayNumAndOpr[(IdxPriority+1):]
	cleanLeftPart, cleanRightPart := deleteBracket(leftPart), deleteBracket(rightPart)

	if len(cleanLeftPart) == 1 && len(cleanRightPart) == 1 {
		return [][]string{
			{" ", " ", cleanLeftPart[0], " ", cleanRightPart[0], " "},
			{" ", " ", " ", " ", " ", " "},
			{" ", " ", ">", " ", ">", " "},
			{" ", " ", "-", " ", "-", " "},
			{" ", " ", "-", " ", "-", " "},
			{arrayNumAndOpr[IdxPriority], "|", "|", "|", "|", " "},
		}
	}

	leftPartAnswer, rightPartAnswer := createBaseArrTree(leftPart), createBaseArrTree(rightPart)
	lenXBloсk := len(leftPartAnswer[0]) + len(rightPartAnswer[0])

	var lineOpr []string
	for i := 0; i < lenXBloсk; i++ {
		switch i {
		case 0:
			lineOpr = append(lineOpr, arrayNumAndOpr[IdxPriority])
		case lenXBloсk - 1:
			lineOpr = append(lineOpr, " ")
		default:
			lineOpr = append(lineOpr, "|")

		}
	}
	answer := sumParts(lineOpr, leftPartAnswer, rightPartAnswer)

	return answer
}

func printTree(inputData string) {
	cleanInputData := cleanAndConvertString(inputData)
	mainFrame := createBaseArrTree(cleanInputData)
	for _, dat := range mainFrame {
		fmt.Println(dat)
	}

	mainFrame = transposition2DArray(mainFrame)
	mainFrame = reverce2DArrayByHorizontal(mainFrame)
	for _, dat := range mainFrame {
		fmt.Println(dat)
	}
}

func main() {
	printTree("5+(3+1)^((1+1)+(6-7))")
}

//func printTree(data string) {
//	arrTree := createArrTree(cleanAndConvertString(data))
//	for _, data := range arrTree {
//		fmt.Println(data)
//	}
//	//fmt.Println()
//	//arrTree = reverse(arrTree)
//	//arrTree = transposition(arrTree)
//	//for _, data := range arrTree {
//	//	fmt.Println(data)
//	//}
//}

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
