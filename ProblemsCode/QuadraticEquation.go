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
		x1, x2, isI := findXForQuadratic(a, b, c)
		if !isI {
			answer := FunctionsPackage.FloatToString(x1) + " " + FunctionsPackage.FloatToString(x2)
			answerSlice = append(answerSlice, answer)
		} else {
			answer := "0"
			if x1 >= 0 {
				answer += "+" + FunctionsPackage.FloatToString(x1) + "i 0"
			} else {
				answer += FunctionsPackage.FloatToString(x1) + "i 0"
			}
			if x2 >= 0 {
				answer += "+" + FunctionsPackage.FloatToString(x2) + "i"
			} else {
				answer += FunctionsPackage.FloatToString(x2) + "i"
			}
			answerSlice = append(answerSlice, answer)
		}
	}
	for _, i := range answerSlice {
		fmt.Print(i + "; ")
	}
	fmt.Println()
	fmt.Println(len(answerSlice))

}

func findXForQuadratic(a int, b int, c int) (x1 float64, x2 float64, isI bool) {
	if math.Pow(float64(b), 2)-4*float64(a)*float64(c) < 0 {
		isI = true
		x1 = (float64(-b) + math.Sqrt(-(math.Pow(float64(b), 2) - 4*float64(a)*float64(c)))) / (float64(2 * a))
		x2 = (float64(-b) - math.Sqrt(-(math.Pow(float64(b), 2) - 4*float64(a)*float64(c)))) / (float64(2 * a))
	} else {
		isI = false
		x1 = (float64(-b) + math.Sqrt(math.Pow(float64(b), 2)-4*float64(a)*float64(c))) / (float64(2 * a))
		x2 = (float64(-b) - math.Sqrt(math.Pow(float64(b), 2)-4*float64(a)*float64(c))) / (float64(2 * a))
	}
	return
}
