package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func ReverseString() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputReverseString")
	inputSlice := strings.Split(inputString, " ")
	fmt.Println(inputSlice)
	var reversedSlice []string
	for _, word := range inputSlice {
		var tempString string
		var wordLen = len(word)
		for letter := 0; letter < wordLen; letter++ {
			tempString += string(word[wordLen - letter - 1])
		}
		reversedSlice = append(reversedSlice, tempString)

	}
	reservedOrderSlice := ReverseSlice(reversedSlice)
	fmt.Println(reversedSlice)
	fmt.Println(reservedOrderSlice)

}
func ReverseSlice(stringIn []string) []string {
	var tempSlice []string
	sliceLen := len(stringIn)
	for i := 0; i < sliceLen; i++{
		tempSlice = append(tempSlice, stringIn[sliceLen - i - 1])
	}
	return tempSlice

}
