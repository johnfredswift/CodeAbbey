package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Palindromes() {
	stringIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputPalindromes")
	stringSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(stringIn)
	stringSlice = stringSlice[1:]
	var processedStringSlice []string
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	for _, stringSlice := range stringSlice{
		processedString := reg.ReplaceAllString(string(stringSlice), "")
		processedString = strings.ToLower(processedString)
		processedStringSlice = append(processedStringSlice, processedString)
	}
	var resultSlice []string
	for _, line := range processedStringSlice{
		isPalindrome := true
		for i := 0; i < (len(line)/2); i++{
			if line[i] != line[len(line) - i - 1]{isPalindrome = false}
		}
		if isPalindrome {resultSlice = append(resultSlice, "Y")
		} else {resultSlice = append(resultSlice, "N")}
	}
	fmt.Println(processedStringSlice)
	fmt.Println(resultSlice)
}

