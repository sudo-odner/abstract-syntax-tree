package main

import (
	"abstract-syntax-tree/internal/app"
)

func main() {
	app.Start()
}

/*
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

func prioritizationOperation(array []string) []operator {
	arrayOperatorStruct := make([]operator, len(array))
	raiseFlag := 0
	var newOperator operator

	for idx, data := range array {
		switch data {
		case "(":
			raiseFlag += 4
		case ")":
			raiseFlag -= 4
		case "^":
			newOperator = operator{data, idx, int(3 + raiseFlag)}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "*":
			newOperator = operator{data, idx, int(2 + raiseFlag)}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "/":
			newOperator = operator{data, idx, int(2 + raiseFlag)}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "+":
			newOperator = operator{data, idx, int(1 + raiseFlag)}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "-":
			newOperator = operator{data, idx, int(1 + raiseFlag)}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		}
	}

	return arrayOperatorStruct
}

// Генерация массива с последовательностью индексов операторов относительно строки, где индекс массива это приоретет выполеннеия с конца
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
	arrayPriorityAndIdx = priorityOperator.reverce2DArrayByVertical(arrayPriorityAndIdx)
	arrayPriorityAndIdx = sortArrByFirstIdx(arrayPriorityAndIdx)

	var arrayIdx []int64

	for _, data := range arrayPriorityAndIdx {
		arrayIdx = append(arrayIdx, data[1])
	}

	return arrayIdx
}

// Совмещение трех блоков
func sumParts(leftPart, rightPart [][]string) [][]string {
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

	return newArr
}

// Удаление из массива скобок
func deleteBracket(data []string) []string {
	var newData []string
	for _, word := range data {
		if !(priorityOperator.existsInSlice(word, []string{"(", ")"})) {
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
			{" ", arrayNumAndOpr[0]},
			{" ", ">"},
			{" ", "-"},
			{" ", "-"},
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
			{" ", " ", " ", cleanLeftPart[0], " ", cleanRightPart[0]},
			{" ", " ", " ", " ", " ", " "},
			{" ", " ", " ", ">", " ", ">"},
			{" ", " ", " ", "-", " ", "-"},
			{" ", " ", " ", "-", " ", "-"},
			{" ", arrayNumAndOpr[IdxPriority], "|", "|", "|", "|"},
			{" ", ">", " ", " ", " ", " "},
			{" ", "-", " ", " ", " ", " "},
			{" ", "-", " ", " ", " ", " "},
		}
	}

	leftPartAnswer, rightPartAnswer := createBaseArrTree(leftPart), createBaseArrTree(rightPart)
	answer := sumParts(leftPartAnswer, rightPartAnswer)

	lineOpre := make([]string, len(answer[0]))
	for i := range lineOpre {
		if i >= (len(answer[0])-4) && len(leftPartAnswer) != 4 {
			lineOpre[i] = " "
		} else {
			lineOpre[i] = "|"
		}
	}

	answer = append(answer, lineOpre)

	answer = priorityOperator.transposition2DArray(answer)
	answer = priorityOperator.reverce2DArrayByVertical(answer)
	lineSpace := make([]string, len(answer[0]))
	for i := range lineSpace {
		if i == len(lineSpace)-1 {
			lineSpace[i] = arrayNumAndOpr[IdxPriority]
		} else {
			lineSpace[i] = " "
		}
	}
	answer = append(answer, lineSpace)

	lineSpace = make([]string, len(answer[0]))
	for i := range lineSpace {
		lineSpace[i] = " "
	}
	answer = append(answer, lineSpace)
	answer = priorityOperator.reverce2DArrayByVertical(answer)
	answer = priorityOperator.transposition2DArray(answer)

	// написатоь чтуку
	printOpreVisual := make([]string, len(answer[0]))
	for i := range printOpreVisual {
		if i == 1 {
			printOpreVisual[i] = ">"
		} else {
			printOpreVisual[i] = " "
		}
	}
	answer = append(answer, printOpreVisual)

	printOpreVisual = make([]string, len(answer[0]))
	for i := range printOpreVisual {
		if i == 1 {
			printOpreVisual[i] = "-"
		} else {
			printOpreVisual[i] = " "
		}
	}
	answer = append(answer, printOpreVisual)
	answer = append(answer, printOpreVisual)

	return answer
}

func printTree(inputData string) {
	cleanInputData := cleanString(inputData)
	mainFrame := createBaseArrTree(cleanInputData)
	for _, dat := range mainFrame {
		fmt.Println(dat)
	}
	fmt.Println()
	mainFrame = pkg.Transposition2DArray(mainFrame)
	mainFrame = pkg.Reverce2DArrayByHorizontal(mainFrame)
	for _, dat := range mainFrame {
		fmt.Println(dat)
	}
}

//func main() {
//	//printTree("5+(3+1)^((1+1)+(6-7))")
//	printTree("5+(3+1)^((1+1)+(6-7.4)^(232+1-(3-1)))")
//}

// +, -, /, *, ^, (, )


*/
/*
	3 + 4 / 3 - 18 ^ ( 2 * 5 )
	["3", "+", "4", "/", "3", "-", "18", "^", "(", "2", "*", "5", ")"]
	(3 * 2) + (3 / 3) + 4
                  +
           |          |
           +          4
      |        |
      *        /
	/   \    /   \
	3   2    3   3
+
|--> +
|    |--> *
|	 |    |
|    |    |--> 3
|    |	  |
|	 |    |--> 2
|  	 |
|	 |--> /
|         |
|		  |--> 3
|		  |
|		  |--> 3
|
|--> 4

*/
