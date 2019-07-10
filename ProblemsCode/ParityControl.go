package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func ParityControl() {
	dataIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputParityControl")
	dataSlice := strings.Split(dataIn, " ")
	var numSlice []int
	for _, num := range dataSlice {
		numSlice = append(numSlice, FunctionsPackage.StringToInt(num))
	}
	fmt.Println(dataSlice)
	var binaryAsStringSlice []string
	for _, letter := range numSlice {

		binaryLetter := fmt.Sprintf("%08b", letter)
		binaryAsStringSlice = append(binaryAsStringSlice, binaryLetter)
	}
	fmt.Println(binaryAsStringSlice)
	var finalBinaryAsStringSlice []string
	var counter []int
	for i, letter := range binaryAsStringSlice {
		if strings.Count(letter, "1")%2 == 0 {
			finalBinaryAsStringSlice = append(finalBinaryAsStringSlice, letter)
			counter = append(counter, i)
		}
	}
	fmt.Println(finalBinaryAsStringSlice)
	for i, letter := range finalBinaryAsStringSlice {
		finalBinaryAsStringSlice[i] = "0" + letter[1:]
	}
	fmt.Println(finalBinaryAsStringSlice)
	var binarySlice []byte
	for _, letter := range finalBinaryAsStringSlice {
		binarySlice = append(binarySlice, byte(FunctionsPackage.StringToInt(letter)))
	}
	for _, letter := range binarySlice {
		fmt.Print(byte(letter), " ")
		fmt.Print(string(letter), " ,")
	}
	fmt.Println()
	for _, i := range counter {
		if numSlice[i] > 128 {
			fmt.Print(string(numSlice[i] - 128))
		} else {
			fmt.Print(string(numSlice[i]))
		}

	}
}
