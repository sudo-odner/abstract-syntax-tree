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

func (a *OperatorPath) getMinIndexPriority() int {
	var (
		minPriority      = 999999999
		indexMinPriority int
	)
	for _, data := range a.OperatorPath {
		if data.priority < minPriority {
			minPriority = data.priority
			indexMinPriority = data.idxInArray
		}
	}

	return indexMinPriority
}

func (a *OperatorPath) SplitByIndexString() (OperatorPath, string, OperatorPath) {
	var (
		opera     string
		leftPart  OperatorPath
		rightPart OperatorPath
	)
	index := a.getMinIndexPriority()

	leftPart.DataString = append(leftPart.DataString, a.DataString[:index]...)
	rightPart.DataString = append(rightPart.DataString, a.DataString[(index+1):]...)

	for i := range a.OperatorPath {
		if a.OperatorPath[i].idxInArray == index {
			opera = a.OperatorPath[i].name
			leftPart.OperatorPath = append(leftPart.OperatorPath, a.OperatorPath[:i]...)
		}
		if a.OperatorPath[i].idxInArray > index {
			newOperator := operator{a.OperatorPath[i].name, a.OperatorPath[i].idxInArray - index - 1, a.OperatorPath[i].priority}
			rightPart.OperatorPath = append(rightPart.OperatorPath, newOperator)
		}
	}
	return leftPart, opera, rightPart
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
			for i := range a.OperatorPath {
				if a.OperatorPath[i].priority >= a.OperatorPath[idx].priority {
					a.OperatorPath[i].priority++
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
