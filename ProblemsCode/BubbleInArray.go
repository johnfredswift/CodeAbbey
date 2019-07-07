package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func BubbleInArray() {
	stringIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputBubbleInArray")
	stringSlice := strings.Split(stringIn, " ")
	stringSlice = stringSlice[:len(stringSlice) - 1]
	//fmt.Println(stringSlice)
	var intSlice []int
	for _, i := range stringSlice{
		intSlice = append(intSlice, FunctionsPackage.StringToInt(string(i)))
	}

	swaps := 0
	for i := 0; i < len(intSlice) - 1; i++{
		if intSlice[i] > intSlice[i + 1]{
			swaps ++
			temp := intSlice[i + 1]
			intSlice[i + 1] = intSlice[i]
			intSlice[i] = temp
		}
	}
	fmt.Println(swaps, FunctionsPackage.SliceChecksum(intSlice))
}


