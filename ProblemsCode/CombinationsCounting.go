package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

func CombinationsCounting() {
	dataIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputCombinationsCounting")
	_, dataSlices := FunctionsPackage.StringToXRows(dataIn)

	for _, i := range dataSlices {
		n, k := FunctionsPackage.StringToInt(i[0]), FunctionsPackage.StringToInt(i[1])
		fmt.Print(calcCombinations(n, k), " ")
	}
}

func calcFactorial(n int) (result int) {
	if n > 0 {
		result = n * calcFactorial(n-1)
		return result
	}
	return 1
}
func calcCombinations(n int, k int) int {
	return calcFactorial(n) / (calcFactorial(k) * calcFactorial(n-k))
}
