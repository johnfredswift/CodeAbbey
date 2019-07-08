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
		fmt.Print(calcCombinations(uint64(n), uint64(k)), " ")
	}
}

func calcFactorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * calcFactorial(n-1)
		return result
	}
	return 1
}
func calcFactorialToBound(n uint64, lowerBound uint64) (result uint64) {
	if n > lowerBound {
		result = n * calcFactorialToBound(n-1, lowerBound)
		return result
	}
	return lowerBound
}

func calcCombinations(n uint64, k uint64) uint64 {
	if n-k > k {
		return calcFactorialToBound(n, n-k+1) / calcFactorial(k)
	}
	return calcFactorialToBound(n, k+1) / calcFactorial(n-k)
}
