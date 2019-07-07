package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func SquareRoot() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputSquareRoot")
	_, inputSlice := FunctionsPackage.StringToXRows(inputString)
	fmt.Println(inputSlice)
	var intSlice [][]int
	var resultSlice []float64
	for _, pair := range inputSlice{
		var temp []int
		for _, num := range pair{
			temp = append(temp, FunctionsPackage.StringToInt(string(num)))
		}
		intSlice = append(intSlice, temp)
	}
	for _, pair := range intSlice{
		r := 1.0
		for i := 0; i < pair[1]; i++{
			r = (r + (float64(pair[0]) / r)) / 2
		}
		resultSlice = append(resultSlice, r)
	}
	fmt.Println(resultSlice)
}
