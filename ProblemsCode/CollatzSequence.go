package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func CollatzSequence() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputCollatzSequence")
	inputSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(inputString)
	inputSlice = inputSlice[1:]
	stringSlice := strings.Split(inputSlice[0], " ")
	fmt.Println(stringSlice)
	var intSlice []int
	for _, num := range stringSlice{
		intSlice = append(intSlice, FunctionsPackage.StringToInt(string(num)))
	}
	var resultSlice []int
	for _, num := range intSlice{
		tempNum := num
		count := 0
		for ; tempNum != 1;{
			if tempNum % 2 == 0{
				tempNum = tempNum / 2
			} else{
				tempNum = (tempNum * 3) + 1
			}
			count++
		}
		resultSlice = append(resultSlice, count)
	}
	fmt.Println(resultSlice)
}
