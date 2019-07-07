package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
)

func TwoPrinters() {
	dataIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputTwoPrinters")
	_, dataSlices := FunctionsPackage.StringToXRows(dataIn)

	for _, line := range dataSlices{
		a := float64(FunctionsPackage.StringToInt(line[0]))
		b := float64(FunctionsPackage.StringToInt(line[1]))
		c := float64(FunctionsPackage.StringToInt(line[2]))
		result := int(completionTime(a, b, c))
		fmt.Print(result, " ")
	}
}

func completionTime(a float64, b float64, c float64) float64{
	ot := (a * b * c) / (a + b)
	an1 := math.Floor(ot / a)
	bn1 := c - an1
	t1 := math.Max(a * an1, b * bn1)
	an2 := math.Ceil(ot/a)
	if an1 == an2{return t1}
	bn2 := c - an2
	t2 := math.Max(a * an2, b * bn2)
	return math.Min(t1, t2)

}
