package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func FoolsDay2014() {
	dataIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputFoolsDay2014")
	_, rowSlices := FunctionsPackage.StringToXRows(dataIn)

	for _, i := range rowSlices{
		total := 0
		for _, num := range i{
			total += FunctionsPackage.StringToInt(num) * FunctionsPackage.StringToInt(num)
		}
		fmt.Print(total, " ")
	}
}
