package ProblemsCode

import (
	functions "CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func WeightedSumOfDigits() {
	inputString := functions.TextFileIn("\\InputFiles\\inputWeightedSumOfDigits")
	newLineSlice := functions.TrimSpaceNewlineInStringToArray(inputString)
	sliceLength := newLineSlice[0]
	stringSlice := strings.Split(newLineSlice[1], " ")
	fmt.Println("Length of Input Data: ", sliceLength)



	var digitsArray []int
	for _, i := range stringSlice{
		sum := 0
		for counter, char := range i{
			sum += (counter + 1) * functions.StringToInt(string(char))
		}
		digitsArray = append(digitsArray, sum)
	}
	fmt.Println(digitsArray)
}
