package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strconv"
	"strings"
)

func NeumannsRandomGenerator() {
	stringIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputNeumannsRandomGenerator")
	stringSlice := strings.Split(FunctionsPackage.TrimSpaceNewlineInStringToArray(stringIn)[1], " ")
	var intSlice []int
	var resultSlice []int
	for _, i := range stringSlice{
		intSlice = append(intSlice, FunctionsPackage.StringToInt(i))
	}
	for _, num := range intSlice{
		isRepeated := false
		iterations := 0
		result := num
		var previousSlice []int
		for !(isRepeated && iterations != 0){
			iterations++
			result = result * result
			result = truncate(result)
			for _, i := range previousSlice{
				fmt.Println(i, result)
				if i == result{isRepeated = true}
			}
			previousSlice = append(previousSlice, result)
		}
		resultSlice = append(resultSlice, len(previousSlice))
		fmt.Println()
	}
	fmt.Println(resultSlice)
}
func truncate(intIn int)int{
	temp := strconv.Itoa(intIn)
	for i := 0; len(temp) < 8; i++{
		temp = "0" + temp
	}
	fmt.Println(temp[2:6])

	return FunctionsPackage.StringToInt(temp[2:6])
}