package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func GreatestCommonDivisor() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputGreatestCommonDivisor")
	_, inputSlice := FunctionsPackage.StringToXRows(inputString)
	var intSlice [][]int
	var resultSlice [][]int
	for _, duel := range inputSlice{

		var temp []int
		for _, num := range duel{

			temp = append(temp, FunctionsPackage.StringToInt(string(num)))
		}
		intSlice = append(intSlice, temp)
	}
	for _, pair := range intSlice{
		var temp []int
		temp = append(temp, GCD(pair[0], pair[1]))
		temp = append(temp, LCM(pair[0], pair[1], temp[0]))
		resultSlice = append(resultSlice, temp)

	}

	for _, pair := range resultSlice{
		fmt.Print("(", pair[0], " ", pair[1],") ")
	}
}
func GCD (x int, y int) int {
	for{
		if x > y {x = x - y}
		if y > x {y = y - x}
		if x == y {return x}
	}
}
func LCM (x int, y int, z int) int {
	return x * y / z
}