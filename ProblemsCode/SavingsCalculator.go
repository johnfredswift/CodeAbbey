package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
	"strings"
)

func SavingsCalculator() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputSavingsCalculator")
	sInSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(sIn)
	sInSlice = sInSlice[1:]
	var floatSlice [][]float64
	for _, i := range sInSlice{
		var temp []float64
		strings := strings.Split(i, " ")
		for _, a := range strings{
			temp = append(temp, FunctionsPackage.StringToFloat(a))
		}
		floatSlice = append(floatSlice, temp)
	}
	var resultingYears []int
	for _, i := range floatSlice{
		resultingYears = append(resultingYears, calcInterest(i)) //[0] Initial, [1] Stop Amount, [2] Interest Rate
	}
	fmt.Println(resultingYears)
}
func calcInterest(iIn []float64)int{ //[0] Initial Capital, [1] Stop Amount, [2] Interest Rate
	x := iIn[0] //[0] Initial Capital
	i := 0
	for  ;x < iIn[1]; x = x * (1 + iIn[2]/100){
		x = math.Floor(x*100)/100
		i++
	}
	return i
}
