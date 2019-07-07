package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"sort"
)

func SortWithIndexes() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputSortWithIndexes")
	_, inputSlice := FunctionsPackage.StringToXRows(inputString)
	//fmt.Println(inputSlice)
	var intSlice []int
	for _, i := range inputSlice[0]{
		intSlice = append(intSlice, FunctionsPackage.StringToInt(string(i)))
	}
	fmt.Println("Int Slice: ", intSlice)
	var indexMap = make(map[int]Index)
	for i := 0; i < len(intSlice); i ++{
		indexMap[intSlice[i]] = Index{i + 1}
}
	fmt.Println("Index Map: ", indexMap)
	sort.Ints(intSlice)
	fmt.Print("Original Positions: ")
	for _, i := range intSlice{
		fmt.Print(indexMap[i].pos, " ")
	}
}

type Index struct {
	pos int
}

