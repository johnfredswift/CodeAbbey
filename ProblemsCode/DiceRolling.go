package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func DiceRolling() {
	inputString := FunctionsPackage.TextFileIn("//InputFiles//inputDiceRolling")
	//fmt.Println(inputString)
	diceSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(inputString)
	diceSlice = diceSlice[1:]
	//fmt.Println(diceSlice)
	var sides  = 6.0
	var boundary  = 1.0/sides
	var resultsSlice [] float64

	for _, roll := range diceSlice{
		for i :=  boundary; i < sides; i += boundary{
			if FunctionsPackage.StringToFloat(roll) < i{
				resultsSlice = append(resultsSlice, i / boundary)
				break
			}
		}
	}
	fmt.Println(resultsSlice)

}
