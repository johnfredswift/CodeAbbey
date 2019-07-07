package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func LinearCongruentialGenerator() {
	input := FunctionsPackage.TextFileIn("\\InputFiles\\inputLinearCongruentialGenerator")
	_, slices := FunctionsPackage.StringToXRows(input)
	var rows [][]int //  A, C, M, X0, N
	for _, i := range slices{
		var temp []int
		for _, num := range i{
			temp = append(temp, FunctionsPackage.StringToInt(num))
		}
		rows = append(rows, temp)
	}
	var results []int
	for _, row := range rows{
		runningTotal := row[3]
		for counter := 0; counter < row[4]; counter++{
			runningTotal = (row[0] * runningTotal + row[1]) % row[2] //Xnext = (A * Xcur + C) % M
		}
		results = append(results, runningTotal)
	}
	fmt.Println(results)
}

