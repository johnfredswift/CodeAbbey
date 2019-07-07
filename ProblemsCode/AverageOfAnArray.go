package ProblemsCode

import (
	functions "CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func AverageOfAnArray() (r []float64){
	inputString := functions.TextFileIn("\\InputFiles\\inputAverageOfAnArray")
	newLineSlice := functions.TrimSpaceNewlineInStringToArray(inputString)
	sliceLength := newLineSlice[0]
	newLineSlice = newLineSlice[1:]


	fmt.Println("Slice Length is: ", sliceLength)
	var twoDStringArray [][]string
	for _, i := range newLineSlice{
		line := strings.Split(i, " ")
		line = line[:len(line) - 1]
		twoDStringArray = append(twoDStringArray, line)
	}
	fmt.Println(twoDStringArray)
	var averagesSlice []float64

	for _, row := range twoDStringArray{
		sum := 0
		for _, stringNum := range row{
			sum += functions.StringToInt(stringNum)
		}
		averagesSlice = append(averagesSlice, float64(sum) / float64(len(row)))
	}
	fmt.Println(averagesSlice)
	fmt.Printf("%.0f", averagesSlice)
	return averagesSlice
	
}
