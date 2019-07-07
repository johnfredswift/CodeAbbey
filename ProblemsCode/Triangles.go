package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"sort"
)

func Triangles() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputTriangles")
	lengthOfSlice, leftSlice, middleSlice, rightSlice := FunctionsPackage.StringToThreeRows(inputString)


	var trianglesSlice [][]int
	//fmt.Println(trianglesArray)

	for i := 0; i < lengthOfSlice; i ++{
		l := FunctionsPackage.StringToInt(leftSlice[i])
		m := FunctionsPackage.StringToInt(middleSlice[i])
		r := FunctionsPackage.StringToInt(rightSlice[i])
		a := []int{l, m, r}
		trianglesSlice = append(trianglesSlice, a)
	}

	//fmt.Println(trianglesArray)
	var resultSlice []int
	for i := 0; i < lengthOfSlice; i++{
		sort.Ints(trianglesSlice[i])
		if trianglesSlice[i][0] + trianglesSlice[i][1] > trianglesSlice[i][2] {
			resultSlice = append(resultSlice, 1)
		} else { resultSlice = append(resultSlice, 0) }
	}
	fmt.Println(trianglesSlice[0])
}



