package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strconv"
	"strings"
)

func BitCount() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputBitCount")
	row := FunctionsPackage.TrimSpaceNewlineInStringToArray(sIn)[1]
	values := strings.Split(row, " ")

	var intSlice []int64
	for _, i := range values{
		intSlice = append(intSlice, int64(FunctionsPackage.StringToInt(i)))
	}
	for _, i := range intSlice{
		fmt.Print(StringToBinary(i), " ")


	}
}
func StringToBinary(iIn int64) int64 {
	isNeg := false
	if iIn < 0{
		isNeg = true
		iIn = iIn * - 1
	}else {
		isNeg = false
	}
	stringBinary := strconv.FormatInt(int64(iIn), 2)

	for 32 != len(stringBinary){
		stringBinary = "0" + stringBinary
	}

	var counter int64
	for _, i := range stringBinary{
		if string(i) == "1"{
			counter++
		}
	}
	if isNeg{
		var remainderCounter int64
		for i := 0; i < len(stringBinary) - 1 && string(stringBinary[len(stringBinary)-i-1]) == "0"; i++{
			remainderCounter++
		}
		return 33 - counter - remainderCounter
	} else {return counter}

}
