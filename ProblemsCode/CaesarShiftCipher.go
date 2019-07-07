package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"regexp"
	"strings"
)

func CaesarShiftCipher() {
	s := FunctionsPackage.TextFileIn("\\InputFiles\\inputCaesarShiftCipher")
	sSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(s)
	cipherNumber := int32(FunctionsPackage.StringToInt(strings.Split(sSlice[0], " ")[1]))
	sSlice = sSlice[1:]
	var resultsSlice []string
	reLower := regexp.MustCompile("^[a-z]*$")
	reHigher := regexp.MustCompile("^[A-Z]*$")
	fmt.Println(sSlice, cipherNumber, resultsSlice)

	for _, line := range sSlice{
		var temp string
		for _, letter := range line{
			var i int32
			if reLower.MatchString(string(letter)){
				i = letter - cipherNumber
				//i to string
				//temp += i
				fmt.Println(string(letter),string(i), "Lower")
			} else if reHigher.MatchString(string(letter)){
				if letter - cipherNumber < 65{
					i = (letter - cipherNumber) + 26
				} else{i = letter - cipherNumber}
				} else {i = letter}
				fmt.Println(string(letter),string(i), "Upper")
				temp += string(i)
			}

		resultsSlice = append(resultsSlice, temp)
		}

	fmt.Println(resultsSlice)
	}

