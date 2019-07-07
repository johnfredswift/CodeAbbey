package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func BicycleRace() {
	s := FunctionsPackage.TextFileIn("\\InputFiles\\inputBicycleRace")
	sSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(s)[1:]
	var floatSlice [][]float64
	for _, rows := range sSlice{
		row := strings.Split(rows, " ")
		var tempSlice []float64
		for _, i := range row{
			tempSlice = append(tempSlice, FunctionsPackage.StringToFloat(i))
		}
		floatSlice = append(floatSlice, tempSlice)
	}
	for _, race := range floatSlice{
			fmt.Print((race[1] / (race[1] + race[2])) * race[0], " ")
	}


}
