package InterviewCode

import (
"CodeAbbey/FunctionsPackage"
"fmt"
"strings"
)

func BubbleSort2() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputBubbleSort") //import file from a text file
	newLineSlice := FunctionsPackage.TrimSpaceNewlineInStringToArray(inputString) //Uses a function to split into arrays be linebreak
	//sliceLength := newLineSlice[0]//first element in the array
	newLineSlice = newLineSlice[1:]//remove only the first element in an array
	stringSlice := strings.Split(newLineSlice[0], " ")
	var intSlice []int
	for _, i := range stringSlice{
		intSlice = append(intSlice, FunctionsPackage.StringToInt(string(i)))
	}
	fmt.Println(intSlice)
	swaps, totalSwaps := sortSlice(intSlice)
	fmt.Println(intSlice)
	fmt.Println(swaps, totalSwaps)


}

func sortSlice(sliceIn []int) (passes int,totalSwaps int){
	swaps := 0
	for ; !(passes != 0 && swaps == 0) ; passes ++{ //&& swaps != 0
		fmt.Println("Pass number:", passes)
		swaps = 0
		for i := 0; i < len(sliceIn) - 1; i++{
			if sliceIn[i] > sliceIn[i + 1]{
				swaps ++
				temp := sliceIn[i + 1]
				sliceIn[i + 1] = sliceIn[i]
				sliceIn[i] = temp
			}

		}
		totalSwaps += swaps
	}
	return
}


