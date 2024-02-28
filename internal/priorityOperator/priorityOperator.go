package priorityOperator

import "slices"

type OperatorPath struct {
	DataString   []string
	OperatorPath []operator
}

type operator struct {
	name       string
	idxInArray int
	priority   int
}

func New(std []string) OperatorPath {
	var newOperatorPath OperatorPath
	newOperatorPath.DataString = std
	newOperatorPath.OperatorPath = prioritizationOperation(std)

	newOperatorPath.cleanRepeatPriority()
	newOperatorPath.cleanBracket()
	return newOperatorPath
}

func prioritizationOperation(array []string) []operator {
	arrayOperatorStruct := make([]operator, 0, len(array))
	raiseFlag := 0
	countBracket := 0
	var newOperator operator

	for idx, data := range array {
		switch data {
		case "(":
			raiseFlag += 4
			countBracket++
		case ")":
			raiseFlag -= 4
			countBracket++
		case "^":
			newOperator = operator{data, idx - countBracket, 3 + raiseFlag}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "*":
			newOperator = operator{data, idx - countBracket, 2 + raiseFlag}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "/":
			newOperator = operator{data, idx - countBracket, 2 + raiseFlag}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "+":
			newOperator = operator{data, idx - countBracket, 1 + raiseFlag}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		case "-":
			newOperator = operator{data, idx - countBracket, 1 + raiseFlag}
			arrayOperatorStruct = append(arrayOperatorStruct, newOperator)
		}
	}
	arrayOperatorStruct = slices.Clip(arrayOperatorStruct)

	return arrayOperatorStruct
}

func (a *OperatorPath) checkSamePriority(index int) bool {
	for _, data := range a.OperatorPath[(index + 1):] {
		if a.OperatorPath[index].priority == data.priority {
			return true
		}
	}
	return false
}

func (a *OperatorPath) rotateOperatorPathByHorizontal() {
	var element int
	var copyOperatorPath []operator
	copyOperatorPath = append(copyOperatorPath, a.OperatorPath...)

	for idx := range a.OperatorPath {
		element = len(a.OperatorPath) - 1 - idx
		a.OperatorPath[idx] = copyOperatorPath[element]
	}
}

func (a *OperatorPath) cleanRepeatPriority() {

	a.rotateOperatorPathByHorizontal()
	for idx := range a.OperatorPath {
		for a.checkSamePriority(idx) {
			for i := range a.OperatorPath[(idx + 1):] {
				if a.OperatorPath[idx+i+1].priority >= a.OperatorPath[idx].priority {
					a.OperatorPath[idx+i+1].priority++
				}
			}
		}
	}
	a.rotateOperatorPathByHorizontal()
	a.OperatorPath = slices.Clip(a.OperatorPath)
}

func (a *OperatorPath) cleanBracket() {
	newDataString := make([]string, 0, len(a.DataString))

	for _, data := range a.DataString {
		if data == ")" || data == "(" {
		} else {
			newDataString = append(newDataString, data)
		}
	}

	a.DataString = newDataString
	a.DataString = slices.Clip(a.DataString)
}
