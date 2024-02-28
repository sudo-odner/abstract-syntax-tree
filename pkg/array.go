package pkg

// Проверяет существует ли елемент в слайсе
func ExistsInSlice[T comparable](value T, array []T) bool {
	for _, valueArray := range array {
		if valueArray == value {
			return true
		}
	}
	return false
}

func CreateArrOneElement[T comparable](value T, size int) []T {
	newArray := make([]T, 0, size)
	for i := 0; i < size; i++ {
		newArray = append(newArray, value)
	}

	return newArray
}

// Копирование матрицы рамером 2
func copy2DArray[T any](array [][]T) [][]T {
	if len(array) == 0 {
		return array
	}
	newArray := make([][]T, len(array), cap(array))
	for idx := range array {
		newArray[idx] = make([]T, len(array[idx]), cap(array[idx]))
		copy(newArray[idx][:], array[idx][:])
	}

	return newArray
}

// Транспонирование матрицы
func Transposition2DArray[T any](array [][]T) [][]T {
	if len(array) == 0 {
		return array
	}
	lenY, lenX := len(array), len(array[0])

	var maxLen int
	if lenY > lenX {
		maxLen = lenY
	} else {
		maxLen = lenX
	}

	// Создание квадратной матрицы для удобного транспонирования
	copyArray := make([][]T, maxLen)
	for idx := range maxLen {
		copyArray[idx] = make([]T, maxLen)
		if idx < len(array) {
			copy(copyArray[idx][:], array[idx][:])
		} else {
			copy(copyArray[idx][:], array[0][:])
		}
	}

	// Транспонирование матрицы
	for idxY := range lenY {
		for idxX := range lenX {
			copyArray[idxX][idxY] = array[idxY][idxX]
		}
	}

	// Обрезка матрицы от мусора
	copyCutArray := make([][]T, lenX)
	for idxX := range lenX {
		copyCutArray[idxX] = make([]T, lenY)
		copy(copyCutArray[idxX][:], copyArray[idxX][:])
	}

	return copyCutArray
}

// Разворачивает матрицу по вертикали
func reverce2DArrayByVertical[T any](array [][]T) [][]T {
	if len(array) == 0 {
		return array
	}
	copyArray := copy2DArray(array)

	lenY := len(array)
	for idxY := range lenY {
		copy(copyArray[idxY][:], array[(lenY - idxY - 1)][:])
	}

	return copyArray
}

// Разворачивает матрицу по горизонтали
func Reverce2DArrayByHorizontal[T any](array [][]T) [][]T {
	if len(array) == 0 {
		return array
	}
	copyArray := copy2DArray(array)

	lenY, lenX := len(array), len(array[0])
	for idxY := range lenY {
		for idxX := range lenX {
			copyArray[idxY][idxX] = array[idxY][(lenX - idxX - 1)]
		}
	}

	return copyArray
}
