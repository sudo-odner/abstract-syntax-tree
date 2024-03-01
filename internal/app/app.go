package app

import (
	"abstract-syntax-tree/internal/treeType"
)

func Start() {
	inputString := "3 + 4 / 3 - 138 ^ ( 2 * 5 )"
	treeType.PrintTreeFirstType(inputString)

}
