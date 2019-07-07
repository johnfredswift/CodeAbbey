package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"regexp"
)

func MatchingBrackets() {
	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputMatchingBrackets")
	rows := FunctionsPackage.TrimSpaceNewlineInStringToArray(sIn)[1:]
	fmt.Println(rows)
	//var results []bool
	bracketsRegExp := regexp.MustCompile(`[\[\]\{\}\(\)<>]`)
	var cleanSlice []string
	for _, row := range rows {
		goodLocations := bracketsRegExp.FindAllStringSubmatchIndex(row, -1)
		var temp string
		for _, set := range goodLocations {
			temp += string(row[set[0]])
		}
		cleanSlice = append(cleanSlice, temp)
	}
	var results []bool
	for _, i := range cleanSlice {

		results = append(results, checkBrackets(i))
	}
	for _, i := range results {
		if i {
			fmt.Print("1 ")
		} else {
			fmt.Print("0 ")
		}
	}
}
func checkBrackets(stringIn string) bool{
	sliceCounter := []int{0, 0, 0, 0} //Counter For Brackets, []  {}  ()  <>
	for _, i := range stringIn {
		switch {
		case string(i) == "[" || string(i) == "]":
			sliceCounter[0] ++
		case string(i) == "{" || string(i) == "}":
			sliceCounter[1] ++
		case string(i) == "(" || string(i) == ")":
			sliceCounter[2] ++
		case string(i) == "<" || string(i) == ">":
			sliceCounter[3] ++
		}
	}
	for _, i := range sliceCounter{
		if i % 2 != 0{return false}
	}
	temp := ""
	for _, i := range stringIn {
		switch {
		case string(i) == "[":
			temp += "["
		case string(i) == "{":
			temp += "{"
		case string(i) == "(":
			temp += "("
		case string(i) == "<":
			temp += "<"
		case string(i) == "]":
			if len(temp) != 0 && string(temp[len(temp) - 1]) == "["{
				temp = temp [:len(temp)-1]
			} else {return false}
		case string(i) == "}":
			if len(temp) != 0 && string(temp[len(temp) - 1]) == "{"{
				temp = temp [:len(temp)-1]
			} else {return false}
		case string(i) == ")":
			if len(temp) != 0 && string(temp[len(temp) - 1]) == "("{
				temp = temp [:len(temp)-1]
			} else {return false}
		case string(i) == ">":
			if len(temp) != 0 && string(temp[len(temp) - 1]) == "<"{
				temp = temp [:len(temp)-1]
			} else {return false}

		}

	}
	return true
}