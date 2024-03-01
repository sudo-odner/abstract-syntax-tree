package calculator

import (
	"abstract-syntax-tree/internal/priorityOperator"
	"math"
	"strconv"
)

func baseCalculate(data priorityOperator.OperatorPath) float64 {
	if len(data.DataString) == 1 {
		dataFloat, _ := strconv.ParseFloat(data.DataString[0], 64)
		return dataFloat
	}

	idxMinPriority := data.GetMinIndexPriority()
	leftPart, opera, rightPart := data.SplitByIndexString(idxMinPriority)
	if len(leftPart.DataString) == 1 && len(rightPart.DataString) == 1 {
		leftPartFloat, _ := strconv.ParseFloat(leftPart.DataString[0], 64)
		rightPartFloat, _ := strconv.ParseFloat(rightPart.DataString[0], 64)
		switch opera {
		case "-":
			return leftPartFloat - rightPartFloat
		case "+":
			return leftPartFloat + rightPartFloat
		case "*":
			return leftPartFloat * rightPartFloat
		case "/":
			return leftPartFloat / rightPartFloat
		case "^":
			return math.Pow(leftPartFloat, rightPartFloat)
		}
	}
	leftPartFloat, rightPartFloat := baseCalculate(leftPart), baseCalculate(rightPart)

	switch opera {
	case "-":
		return leftPartFloat - rightPartFloat
	case "+":
		return leftPartFloat + rightPartFloat
	case "*":
		return leftPartFloat * rightPartFloat
	case "/":
		return leftPartFloat / rightPartFloat
	case "^":
		return math.Pow(leftPartFloat, rightPartFloat)
	}

	return 0
}

func Сalculate(inputData []string) float64 {
	data := priorityOperator.New(inputData)

	return baseCalculate(data)
}
