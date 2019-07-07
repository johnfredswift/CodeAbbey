package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func JosephusProblem() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputJosephusProblem")
	slice := strings.Split(sIn, " ")
	n, k := FunctionsPackage.StringToInt(slice[0]), FunctionsPackage.StringToInt(slice[1])
	fmt.Println(calcJP(n, k))
}
func calcJP(nIn int, kIn int) int {
	var intSlice []int
	for i := 0; i < nIn; i++{
		intSlice = append(intSlice, i + 1)
	}
	//kCounter := 0
	//sliceCounter := 0
	indexCounter := kIn - 1

	for len(intSlice) != 1 {

		if indexCounter >= len(intSlice){
			indexCounter = indexCounter % len(intSlice)
		}
		fmt.Println(intSlice[indexCounter])
		intSlice = FunctionsPackage.Remove(intSlice, indexCounter)

		indexCounter += kIn - 1
	}

	return intSlice[0]
}



