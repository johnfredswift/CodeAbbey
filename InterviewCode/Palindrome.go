package InterviewCode

import "fmt"

func Palindrome(s string) bool{
	var returnBool bool = true
	var charSlice []string
	for _, char := range s{
		charSlice = append(charSlice, string(char))
	}
	fmt.Print(charSlice)
	for i := 0; i < len(charSlice); i++{
		if charSlice[i] != charSlice[len(charSlice) - 1 - i]{returnBool = false}
	}

	return returnBool
}
