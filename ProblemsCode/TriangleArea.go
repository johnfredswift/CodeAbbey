package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
)

func TriangleArea() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputTriangleArea")
	_, numberSlice := FunctionsPackage.StringToXRows(sIn)
	//var triangleSliceByCood []triangleByCood
	var triangleSliceByLen []triangleByLen
	for _, value := range numberSlice{
		tempTriangle := triangleByCood{
			cood{FunctionsPackage.StringToInt(value[0]),FunctionsPackage.StringToInt(value[1])},
			cood{FunctionsPackage.StringToInt(value[2]),FunctionsPackage.StringToInt(value[3])},
			cood{FunctionsPackage.StringToInt(value[4]),FunctionsPackage.StringToInt(value[5])}}
		triangleSliceByLen = append(triangleSliceByLen, triangleByLen{
			math.Sqrt(
				math.Pow(float64(tempTriangle.bCood.x - tempTriangle.aCood.x), 2) +
					math.Pow(float64(tempTriangle.bCood.y - tempTriangle.aCood.y), 2)),
			math.Sqrt(
				math.Pow(float64(tempTriangle.cCood.x - tempTriangle.aCood.x), 2) +
					math.Pow(float64(tempTriangle.cCood.y - tempTriangle.aCood.y), 2)),
			math.Sqrt(
				math.Pow(float64(tempTriangle.cCood.x - tempTriangle.bCood.x), 2) +
					math.Pow(float64(tempTriangle.cCood.y - tempTriangle.bCood.y), 2))})
	}


	var resultSlice []float64
	for _, triangle := range triangleSliceByLen{
		var s float64
		s = (triangle.ab + triangle.ac + triangle.bc) / 2
		result := math.Sqrt(s * (s - triangle.ab) * (s - triangle.ac) * (s - triangle.bc))
		resultSlice = append(resultSlice, result)
		fmt.Print(math.Round(result*100000)/100000)
		fmt.Print(" ")
	}
	fmt.Printf("%0.00000001f", resultSlice)
}

type triangleByCood struct{
	aCood cood
	bCood cood
	cCood cood
}
type triangleByLen struct{
	ab float64
	ac float64
	bc float64
}
type cood struct {
	x int
	y int
}
