package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func LinearFunction() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputLinearFunction")
	_, dataSlices := FunctionsPackage.StringToXRows(inputString)
	var gradSlice []int
	var yInter []int
	for _, i := range dataSlices {
		var tempSlice []int
		for _, a := range i{
			tempSlice = append(tempSlice, FunctionsPackage.StringToInt(a))
		}
		gradSlice = append(gradSlice, (tempSlice[3] - tempSlice[1])/(tempSlice[2] - tempSlice[0]))
		yInter = append(yInter, tempSlice[1] - (tempSlice[0] * (tempSlice[3] - tempSlice[1])/(tempSlice[2] - tempSlice[0])))
	}
	fmt.Println(dataSlices)
	fmt.Println(gradSlice)
	fmt.Println(yInter)
	fmt.Println()
	for i := 0; i < len(gradSlice); i++{
		fmt.Print("(", gradSlice[i], " ", yInter[i], ") ")
	}



}