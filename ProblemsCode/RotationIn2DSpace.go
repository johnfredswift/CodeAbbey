package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
	"strings"
)

func RotationIn2DSpace() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputRotationIn2DSpace")
	rows := FunctionsPackage.TrimSpaceNewlineInStringToArray(sIn)
	degrees := FunctionsPackage.StringToInt(strings.Split(rows[0], " ")[1])
	rows = rows[1:]
	var stars []star
	for _, row := range rows{
		temp := strings.Split(row, " ")
		stars = append(stars, star{temp[0], coordinates{
			FunctionsPackage.StringToFloat(temp[1]),
			FunctionsPackage.StringToFloat(temp[2])}})
	}
	//fmt.Println(stars)
	//fmt.Println(multipleMatrix(stars[0].coor, rotation(degrees, false)))

	var results []star
	//fmt.Println(rotation(degrees, false))
	for _, i := range stars{
		results = append(results,
			star{i.name, multipleMatrix(i.coor, rotation(degrees))})
	}
	var rounded []star
	for _, i := range results{
		var temp star
		temp.name = i.name
		temp.coor = coordinates{math.Round(i.coor.x*10000)/10000, math.Round(i.coor.y*100)/100}
		rounded = append(rounded, temp)
		fmt.Println(temp)
	}
	swaps := 0
	passes := 0
	for ; !(passes != 0 && swaps == 0) ; passes ++{
		swaps = 0
		for i := 0; i < len(stars) - 1; i++{
			if rounded[i].coor.x > rounded[i + 1].coor.x{
				swaps ++
				temp := rounded[i + 1]
				rounded[i + 1] = rounded[i]
				rounded[i] = temp
			}
		}
	}
	swaps = 0
	passes = 0
	for ; !(passes != 0 && swaps == 0) ; passes ++{
		swaps = 0
		for i := 0; i < len(stars) - 1; i++{
			if rounded[i].coor.y > rounded[i + 1].coor.y{
				swaps ++
				temp := rounded[i + 1]
				rounded[i + 1] = rounded[i]
				rounded[i] = temp
			}
		}
	}
	for i := 0; i < len(stars); i++{
		fmt.Print(rounded[i].name, " ")
	}

}

type star struct{
	name string
	coor coordinates
}
type coordinates struct {
	x float64
	y float64
}

func rotation(angleIn int) (r [][]float64) {
	var floatAngleIn float64
	var temp []float64
	r = append(r, temp)
	r = append(r, temp)

	floatAngleIn = float64(angleIn)
	r[0] = append(r[0], math.Cos(float64(floatAngleIn * math.Pi/180)))
	r[0] = append(r[0], math.Sin(float64(floatAngleIn * math.Pi/180)))
	r[1] = append(r[1], math.Sin(float64(floatAngleIn * math.Pi/180)))
	r[1] = append(r[1], math.Cos(float64(floatAngleIn * math.Pi/180)))
	return
}


func multipleMatrix(coorIn coordinates, rotationIn [][]float64) coordinates{
	var outCoor coordinates
	outCoor.x = (rotationIn[0][0] * coorIn.x) - (rotationIn[0][1] * coorIn.y)
	outCoor.y = (rotationIn[1][0] * coorIn.x) + (rotationIn[1][1] * coorIn.y)

	return outCoor
	}




