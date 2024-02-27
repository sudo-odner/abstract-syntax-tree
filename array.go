package main

// Копирование матрицы рамером 2
func copy2DArray[T any](array [][]T) [][]T {
	newArray := make([][]T, len(array), cap(array))
	for idx := range array {
		copy(newArray[idx][:], array[idx][:])
	}

	return newArray
}

// Транспонирование матрицы
func transposition2DArray[T any](array [][]T) [][]T {
	copyArray := copy2DArray(array)
	if len(array) == 0 {
		return copyArray
	}

	lenY, lenX := len(array), len(array[0])
	for idxY := range lenY {
		for idxX := range lenX {
			copyArray[idxY][idxX] = array[idxX][idxY]
		}
	}

	return copyArray
}

// Разворачивает матрицу по вертикали
func reverce2DArrayByVertical[T any](array [][]T) [][]T {
	copyArray := copy2DArray(array)
	if len(array) == 0 {
		return copyArray
	}

	lenY := len(array)
	for idxY := range lenY / 2 {
		copy(copyArray[idxY][:], array[(lenY - idxY - 1)][:])
	}

	return copyArray
}

// Разворачивает матрицу по вертикали
func reverce2DArrayByHorizontal[T any](array [][]T) [][]T {
	copyArray := copy2DArray(array)
	if len(array) == 0 {
		return copyArray
	}

	lenY, lenX := len(array), len(array[0])
	for idxY := range lenY {
		for idxX := range lenX / 2 {
			copyArray[idxY][idxX] = array[idxY][(lenX - idxX - 1)]
		}
	}

	return copyArray
}
