package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func DoubleDiceRoll() {
	xIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputDoubleDiceRoll")
	_, xSlice := FunctionsPackage.StringToXRows(xIn)
	intSlice := FunctionsPackage.SliceStringToInt(xSlice)
	for _, i := range intSlice{
		fmt.Print(i[0] % 6 + i[1] % 6 + 2, " ")
	}


}


