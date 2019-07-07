package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func RockPaperScissors() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputRockPaperScissors")
	_, sSlice := FunctionsPackage.StringToXRows(sIn)
	fmt.Println(sSlice)

	for _, set := range sSlice{
		fmt.Print(calcWin(set), " ")
	}
}
func calcWin(stringSlice []string) string{
	oneWins := 0
	twoWins := 0
	for i := 0; i < len(stringSlice); i++{

		playerOne, playerTwo := string(stringSlice[i][0]), string(stringSlice[i][1])
		if (playerOne == "P" && playerTwo == "R") || (playerOne == "R" && playerTwo == "S") ||(playerOne == "S" && playerTwo == "P"){
			oneWins++
		} else if playerOne != playerTwo{
			twoWins++
		}
	}
	if oneWins > twoWins{
		return "1"
	}
	if twoWins > oneWins{
		return "2"
	}
	return "D"
}
