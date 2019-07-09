package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"strconv"
)

func BlackjackCounting() {
	dataIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputBlackjackCounting")
	_, dataSlice := FunctionsPackage.StringToXRows(dataIn)
	for _, game := range dataSlice {
		total := 0
		aceCount := 0
		for _, card := range game {
			if card != "J" && card != "Q" && card != "K" && card != "A" {
				i, _ := strconv.Atoi(card)
				total += i
			}
			if card == "J" || card == "Q" || card == "K" || card == "T" {
				total += 10
			}
			if card == "A" {
				total++
				aceCount++
			}

		}
		for total <= 11 && aceCount != 0 {
			total += 10
			aceCount--
		}
		if total <= 21 {
			fmt.Print(total, " ")
		} else {
			fmt.Print("Bust ")
		}
	}
}
