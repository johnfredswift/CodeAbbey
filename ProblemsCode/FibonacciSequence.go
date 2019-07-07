package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math/big"
	"strings"
)
func FibonacciSequence() {
	stringIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputFibonacciSequence")
	stringLine := FunctionsPackage.TrimSpaceNewlineInString(stringIn)
	stringSlice := strings.Split(stringLine, " ")[1:]
	highest := getHighestNum(stringSlice)
	fmt.Println("Biggest Number is: ", highest)
	fmt.Println("Biggest Number is: ", len(highest), "digits long")
	fmt.Print("Of the list of Fib Numbers, these are there index in the Fib Sequence: ",)
	fibs := calcFibUptoX(highest)
	for _, i := range stringSlice{
		fmt.Print(fibs[i], " ")
	}
}
func getHighestNum(sliceIn []string)string{

	currentHighest := "0"
	for _, i := range sliceIn{
		if len(i) > len(currentHighest) {
			currentHighest = i
		} else if len(i) == len(currentHighest) {
			fmt.Println(i, currentHighest)
			for digit := 0; digit < len(currentHighest); digit++ {
				if FunctionsPackage.StringToInt(string(i[digit])) > FunctionsPackage.StringToInt(string(currentHighest[digit])){
					currentHighest = i
				}
				if i[digit] != currentHighest[digit]{break}
			}
		}
	}
	return currentHighest
}
func calcFibUptoX(highest string) map[string]int{
	var s []string
	highestInt, curr, prev := new(big.Int), new(big.Int), new(big.Int)
	highestInt.SetString(highest, 10)
	prev.SetInt64(0)
	curr.SetInt64(1)
	fibs := make(map[string]int)
	for i := 1; curr.Cmp(highestInt) < 0; i++ {
		curr.Add(curr, prev)
		s = append(s, curr.String())
		fibs[curr.String()] = i
		prev, curr = curr, prev
	}
	return fibs
}
