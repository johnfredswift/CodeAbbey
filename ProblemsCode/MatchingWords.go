package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"sort"
	"strings"
)

func MatchingWords() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputMatchingWords")
	sSlice := strings.Split(sIn, " ")
	sSlice = sSlice[:len(sSlice) - 1]
	fmt.Println(sSlice)
	resultsMap := make(map[string]int)
	for _, i := range sSlice{
		resultsMap[i]++
	}
	fmt.Println(resultsMap)
	var resultsSlice []string
	for key, value := range resultsMap{
		if value > 1{
			resultsSlice = append(resultsSlice, key)
		}
	}
	sort.Strings(resultsSlice)
	fmt.Println(resultsSlice)
}
