package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func PythagoreanTheorem() {
	stringIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputPythagoreanTheorem")

	rows := FunctionsPackage.TrimSpaceNewlineInStringToArray(stringIn)[1:]
	var triangleSlice [][]int
	for _, row := range rows {
		stringSlice := strings.Split(row, " ")
		var temp []int
		for _, num := range stringSlice {
			temp = append(temp, FunctionsPackage.StringToInt(num))
		}
		triangleSlice = append(triangleSlice, temp)
	}
	for _, tri := range triangleSlice{

		fmt.Print(CalcPythag(tri), " ")
	}
}
func CalcPythag(tri []int) string {
	a, b, c := tri[0], tri[1], tri[2]
	result := a*a + b*b - c*c
	if result > 0{
		return "A"
	}else if result == 0{
		return "R"
	}else{
		return "O"
	}
}