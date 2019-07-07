package ProblemsCode

import (
"CodeAbbey/FunctionsPackage"
"fmt"
"strings"
)

func BubbleSort() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputBubbleSort") //import file from a text file
	newLineSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(inputString) //Uses a function to split into arrays be linebreak
	//sliceLength := newLineSlice[0]//first element in the array
	newLineSlice = newLineSlice[1:]//remove only the first element in an array
	stringSlice := strings.Split(newLineSlice[0], " ")
	var intSlice []int
	for _, i := range stringSlice{
		intSlice = append(intSlice, FunctionsPackage.StringToInt(string(i)))
	}
	fmt.Println(intSlice)
	swaps, totalSwaps := FunctionsPackage.SortSlice(intSlice)
	fmt.Println(intSlice)
	fmt.Println(swaps, totalSwaps)
}




