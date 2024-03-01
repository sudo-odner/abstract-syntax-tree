package tree

import (
	"abstract-syntax-tree/internal/priorityOperator"
	"abstract-syntax-tree/pkg"
	"fmt"
	"regexp"
	"strings"
)

// Настройки для спита
var (
	pattern = ""
	regex   = regexp.MustCompile(pattern)
)

// Создание путь (/ или \) до оператора
func createWayOperator(part [][]string, way string) ([][]string, int) {
	var newLineWay []string
	var indexOpera int
	newPart := make([][]string, 0, len(part)+2)

	newLineWay = make([]string, 0, len(part[0]))
	indexOpera = len(part[0]) / 2
	for i := 0; i < 5; i++ {
		check := pkg.GetFirstIndexValueInSlice(string("/*+-^"[i]), part[0])
		if check != -1 {
			indexOpera = check
		}
	}
	newLineWay = append(newLineWay, pkg.CreateArrOneElement(" ", indexOpera)...)
	newLineWay = append(newLineWay, way)
	newLineWay = append(newLineWay, pkg.CreateArrOneElement(" ", len(part[0])-indexOpera-1)...)

	newPart = append(newPart, newLineWay)
	newPart = append(newPart, part...)

	return newPart, indexOpera
}

// Сложение левой и правой части ветор относительно оператора
func sumPartsTreeFirstType(leftPart, rightPart [][]string, opera string) [][]string {
	if len(leftPart) < len(rightPart) {
		difference := len(rightPart) - len(leftPart)

		newLeftPart := make([][]string, 0, len(leftPart)+difference)

		newLeftPart = append(newLeftPart, leftPart...)
		for i := 0; i < difference; i++ {
			newLeftPart = append(newLeftPart, pkg.CreateArrOneElement(" ", len(leftPart[0])))
		}

		leftPart = newLeftPart
	}
	if len(leftPart) > len(rightPart) {
		difference := len(leftPart) - len(rightPart)

		newRightPart := make([][]string, 0, len(rightPart)+difference)

		newRightPart = append(newRightPart, rightPart...)
		for i := 0; i < difference; i++ {
			newRightPart = append(newRightPart, pkg.CreateArrOneElement(" ", len(rightPart[0])))
		}

		rightPart = newRightPart
	}
	var (
		leftIdxOpera  int
		rightIdxOpera int
	)

	leftPart, leftIdxOpera = createWayOperator(leftPart, "/")
	rightPart, rightIdxOpera = createWayOperator(rightPart, "\\")
	leftPart, rightPart = pkg.Transposition2DArray(leftPart), pkg.Transposition2DArray(rightPart)

	sum := append(leftPart, rightPart...)
	sum = pkg.Transposition2DArray(sum)
	sum = pkg.Reverce2DArrayByVertical(sum)

	var lineOpera []string
	middle := leftIdxOpera + ((((len(leftPart) + rightIdxOpera) - leftIdxOpera) + 1) / 2)

	lineOpera = append(lineOpera, pkg.CreateArrOneElement(" ", middle)...)
	lineOpera = append(lineOpera, opera)
	lineOpera = append(lineOpera, pkg.CreateArrOneElement(" ", len(sum[0])-(middle+1))...)

	sum = append(sum, lineOpera)
	sum = pkg.Reverce2DArrayByVertical(sum)

	return sum
}

// Создание дерева
func createBaseArrTree(data priorityOperator.OperatorPath) [][]string {
	if len(data.DataString) == 0 {
		return [][]string{}
	}
	if len(data.DataString) == 1 {
		mainSample := make([][]string, 1)

		mainSample[0] = append(mainSample[0], regex.Split(data.DataString[0], -1)...)

		return mainSample
	}

	idxMinPriority := data.GetMinIndexPriority()
	leftPart, opera, rightPart := data.SplitByIndexString(idxMinPriority)

	if len(leftPart.DataString) == 1 && len(rightPart.DataString) == 1 {
		mainSample := make([][]string, 3)

		lenLeft := len(leftPart.DataString[0])
		lenRight := len(rightPart.DataString[0])

		mainSample[2] = append(mainSample[2], pkg.CreateArrOneElement(" ", 2)...)
		mainSample[2] = append(mainSample[2], regex.Split(leftPart.DataString[0], -1)...)
		mainSample[2] = append(mainSample[2], pkg.CreateArrOneElement(" ", 3)...)
		mainSample[2] = append(mainSample[2], regex.Split(rightPart.DataString[0], -1)...)
		mainSample[2] = append(mainSample[2], pkg.CreateArrOneElement(" ", 2)...)

		mainSample[1] = append(mainSample[1], pkg.CreateArrOneElement(" ", lenLeft+1)...)
		mainSample[1] = append(mainSample[1], "/")
		mainSample[1] = append(mainSample[1], pkg.CreateArrOneElement(" ", 3)...)
		mainSample[1] = append(mainSample[1], "\\")
		mainSample[1] = append(mainSample[1], pkg.CreateArrOneElement(" ", lenRight+1)...)

		mainSample[0] = append(mainSample[0], pkg.CreateArrOneElement(" ", lenLeft+3)...)
		mainSample[0] = append(mainSample[0], opera)
		mainSample[0] = append(mainSample[0], pkg.CreateArrOneElement(" ", lenRight+3)...)

		return mainSample
	}

	leftPartAnswer, rightPartAnswer := createBaseArrTree(leftPart), createBaseArrTree(rightPart)
	return sumPartsTreeFirstType(leftPartAnswer, rightPartAnswer, opera)
}

func Print(inputData []string) {
	data := priorityOperator.New(inputData)
	mainFrame := createBaseArrTree(data)

	for _, dat := range mainFrame {
		fmt.Println(strings.Join(dat, ""))
	}
	fmt.Println()
}
