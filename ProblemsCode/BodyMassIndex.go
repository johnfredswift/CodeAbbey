package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
)

func BodyMassIndex(){
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputBodyMassIndex")
	lengthOfSlice, leftSlice, rightSlice := FunctionsPackage.StringToTwoRows(inputString)
	fmt.Println(leftSlice, rightSlice)
	var resultsArray []string
	for i := 0; i < lengthOfSlice; i++{
		bmi := FunctionsPackage.StringToFloat(leftSlice[i]) / math.Pow(FunctionsPackage.StringToFloat(rightSlice[i]), 2)
		if bmi < 18.5{resultsArray = append(resultsArray, "under")
		} else if bmi < 25.0 {resultsArray = append(resultsArray, "normal")
		} else if bmi < 30.0{resultsArray = append(resultsArray, "over")
		} else {resultsArray = append(resultsArray, "obese") }
	}
	fmt.Println(resultsArray)
}



