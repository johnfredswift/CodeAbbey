package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func ArrayChecksum() { // /sp Slice
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputArrayChecksum")
	newLineSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(inputString)
	newLineSlice = newLineSlice[1:]
	inputSlice := strings.Split(newLineSlice[0], " ")
	fmt.Println(inputSlice)
	result := 0
	var inputSliceInt []int
	for _, i := range inputSlice{
		inputSliceInt = append(inputSliceInt, FunctionsPackage.StringToInt(i))
	}
	fmt.Println(inputSliceInt)
	for _, i := range inputSliceInt{
		result = (result + i) * 113
		result = result % 10000007
	}
	fmt.Println(result)
}
