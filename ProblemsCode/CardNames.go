package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func CardNames() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputCardNames")
	_, inputSlice := FunctionsPackage.StringToXRows(inputString)
	suits := [4]string{"Clubs", "Spades", "Diamonds", "Hearts"}
	ranks := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

	for _, i := range inputSlice[0]{
		tempInt := FunctionsPackage.StringToInt(i)
		fmt.Print(ranks[tempInt % 13], "-of-", suits[tempInt / 13], " ")
	}
}


