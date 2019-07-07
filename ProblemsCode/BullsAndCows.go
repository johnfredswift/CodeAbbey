package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strings"
)

func BullsAndCows() {
	dataIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputBullsAndCows")
	firstRow, rows := FunctionsPackage.StringToXRows(dataIn)
	fmt.Println(firstRow[0])
	for _, i := range rows[0] {
		a, b := 0, 0
		for _, letter := range i {
			if strings.Contains(strings.Split(firstRow[0], " ")[0], string(letter)) {
				a++
			}
		}
		counter := 0
		for _, letter := range i {
			if string(letter) == string(strings.Split(firstRow[0], " ")[0][counter]) {
				b++
				a--
			}
			counter++
		}
		fmt.Print(b, "-", a, " ")
	}
}
