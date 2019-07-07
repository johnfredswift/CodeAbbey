package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
	"math/rand"
)

func DuelChances() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputDuelChances")
	_, inputSlice := FunctionsPackage.StringToXRows(inputString)
	fmt.Println(inputSlice)
	var intSlice [][]int
	var resultSlice []int
	for _, duel := range inputSlice{
		var temp []int
		for _, num := range duel{

			temp = append(temp, FunctionsPackage.StringToInt(string(num)))
		}
		intSlice = append(intSlice, temp)
	}
	for _, i := range intSlice{
		resultSlice = append(resultSlice, int(math.Round(CalcOdds(i[0], i[1])*100)))
	}
	fmt.Println(resultSlice)

}

func CalcOdds(x int, y int) float64{
	xWins := 0
	numberOfDuels := 1000000
 	for i := 0; i < numberOfDuels; i++{
 		for won := false; !won;{

 			if rand.Intn(100) < x{
 				xWins ++
 				won = true
			}
 			if won || rand.Intn(100) < y{
 				won = true
			}
		}
	}
 	return float64(xWins) / float64(numberOfDuels)
}
