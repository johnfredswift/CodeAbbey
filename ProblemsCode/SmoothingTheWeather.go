package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func SmoothingTheWeather() {
	s := FunctionsPackage.TextFileIn("\\inputFiles\\inputSmoothingTheWeather")
	sSlice := strings.Split(FunctionsPackage.TrimSpaceNewlineInStringToArray(s)[1], " ")
	var floatSlice []float64
	for _, i := range sSlice{
		floatSlice = append(floatSlice, FunctionsPackage.StringToFloat(i))
	}
	fmt.Println(Smooth(floatSlice))
}
func Smooth(floatSliceIn []float64)(tempFloatSlice []float64) {

	tempFloatSlice = append(tempFloatSlice, floatSliceIn[0])
	for i := 1; i < len(floatSliceIn) - 1; i++{
		tempFloatSlice = append(tempFloatSlice, (floatSliceIn[i - 1] + floatSliceIn[i] +
			floatSliceIn[i + 1])/3)
	}
	tempFloatSlice = append(tempFloatSlice, floatSliceIn[len(floatSliceIn)-1])
	return tempFloatSlice
}
