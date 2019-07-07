package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func RotateString() {
	stringIn := FunctionsPackage.TextFileIn("\\inputFiles\\inputRotateString")
	fmt.Println(stringIn)
	_, stringArray := FunctionsPackage.StringToXRows(stringIn)
	fmt.Println(stringArray)
	var resultSlice []string

	for _, pair := range stringArray{
		temp := pair[1]
		len := len(pair[1])
		num := FunctionsPackage.StringToInt(pair[0])
		if num > 0{
			temp = temp[num:len] + temp[0:num]
		} else if num < 0 {
			num = 0 - num
			temp = temp[len - num: len] + temp[0:len - num]
		}

		resultSlice = append(resultSlice, temp)
	}
	fmt.Println(resultSlice)
}
