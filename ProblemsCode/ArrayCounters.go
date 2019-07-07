package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func ArrayCounters() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputArrayCounters")
	inputSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(inputString)
	highestNum := FunctionsPackage.StringToInt(string(inputSlice[0][1]))
	inputSlice = inputSlice[1:]
	stringSlice := strings.Split(inputSlice[0], " ")
	fmt.Println(inputSlice)
	var intSlice []int
	for _, num := range stringSlice{
		intSlice = append(intSlice, FunctionsPackage.StringToInt(string(num)))
	}
	fmt.Println(inputSlice)
	var countSlice []int
	for i := 1; i < highestNum; i++{
		countSlice = append(countSlice, 0)
	}
	for _, i := range intSlice{
		countSlice[i - 1] = countSlice[i - 1] + 1
	}
	fmt.Println(countSlice)

}
