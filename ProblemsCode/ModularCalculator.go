package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func ModularCalculator() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputModularCalculator")
	_, dataSlices := FunctionsPackage.StringToXRows(inputString)
	stringSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(inputString)
	firstString := stringSlice[0]
	displayValue := uint64(FunctionsPackage.StringToInt(firstString))


	var operationPairSlice []OperationPair
	for _, pair := range dataSlices{
		var tempOP OperationPair
		tempOP.operation = pair[0]
		tempOP.value = uint64(FunctionsPackage.StringToInt(pair[1]))
		operationPairSlice = append(operationPairSlice, tempOP)
	}
	//modularPair := operationPairSlice[len(operationPairSlice) - 1]
	//fmt.Println(modularPair)
	fmt.Println(displayValue)
	for _, pair := range operationPairSlice{
		displayValue = OperationPairCalc(OperationPairCalc(displayValue, pair), operationPairSlice[len(operationPairSlice) - 1])
		fmt.Println(displayValue, pair)
	}

	fmt.Println(displayValue)

}

func OperationPairCalc(value uint64, OPIn OperationPair) uint64 {

	switch OPIn.operation {
	case "*":
		return value * OPIn.value
	case "+":
		return value + OPIn.value
	case "-":
		return value - OPIn.value
	case "%":
		return value % OPIn.value
	}
	return 0.0
}

type OperationPair struct {
operation string
value uint64
}