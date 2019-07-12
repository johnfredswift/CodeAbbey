package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
)

func QuadraticEquation() {
	dataIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputQuadraticEquation") //-9+3i -9-3i;
	_, dataSlice := FunctionsPackage.StringToXRows(dataIn)
	var answerSlice []string
	for _, i := range dataSlice {
		fmt.Println(i[0], i[1], i[2])
		a, b, c := FunctionsPackage.StringToInt(i[0]), FunctionsPackage.StringToInt(i[1]), FunctionsPackage.StringToInt(i[2])
		x1, x2 := findXForQuadratic(a, b, c)

		answerSlice = append(answerSlice, x1+" "+x2)

	}
	for _, i := range answerSlice {
		fmt.Print(i + "; ")
	}

}

func findXForQuadratic(a int, b int, c int) (x1 string, x2 string) {
	if math.Pow(float64(b), 2)-4*float64(a)*float64(c) < 0 {
		x1PartOne := FunctionsPackage.FloatToString(float64(-b) / (float64(2 * a)))
		x2PartOne := FunctionsPackage.FloatToString(float64(-b) / (float64(2 * a)))
		x1PartTwo := FunctionsPackage.FloatToString(
			math.Sqrt(-(math.Pow(float64(b), 2)-4*float64(a)*float64(c)))/(float64(2*a))) + "i"
		x2PartTwo := FunctionsPackage.FloatToString(
			-(math.Sqrt(-(math.Pow(float64(b), 2) - 4*float64(a)*float64(c))) / (float64(2 * a)))) + "i"
		if string(x1PartTwo[:1]) != "-" {
			x1PartTwo = "+" + x1PartTwo
		}
		if string(x2PartTwo[:1]) != "-" {
			x2PartTwo = "+" + x2PartTwo
		}
		x1 = x1PartOne + x1PartTwo
		x2 = x2PartOne + x2PartTwo
	} else {
		x1 = FunctionsPackage.FloatToString((float64(-b) + math.Sqrt(math.Pow(float64(b), 2)-4*float64(a)*float64(c))) / (float64(2 * a)))
		x2 = FunctionsPackage.FloatToString((float64(-b) - math.Sqrt(math.Pow(float64(b), 2)-4*float64(a)*float64(c))) / (float64(2 * a)))
	}
	return

}
