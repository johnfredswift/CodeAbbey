package FunctionsPackage

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func StringToInt(numIn string) (returnValue int) {
	num, err := strconv.Atoi(numIn)
	if err == nil {
		//fmt.Println(num)
		return num
		//fmt.Println(i)
		//intArray.append(il)
	}
	return 0
}
func StringToUInt64(numIn string) (returnValue uint64) {
	num, err := strconv.ParseUint(numIn, 10, 64)
	if err == nil {
		//fmt.Println(num)
		return num
		//fmt.Println(i)
		//intArray.append(il)
	}
	return 0
}
func TrimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(`\r\n`)
	return re.ReplaceAllString(s, " ")
}
func TrimSpaceNewlineInStringToArray(s string) (stringArrayOut []string) { //Function to take string in and split it into a slice by linebreaks
	re := regexp.MustCompile(`\r\n`) //Create a regexp to detect Linebreaks
	stringArrayOut = re.Split(s, -1) // Split the stringIn by the regexp above
	return                           //return value is already set so no need to pass back the variable
}
func StringToFloat(numIn string) float64 {
	f, _ := strconv.ParseFloat(numIn, 64)
	return f
}
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func CountDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}
func TextFileIn(fileName string) string {
	inputPath := CurrentWorkingDirectory() + fileName
	absPath, _ := filepath.Abs(inputPath)
	b, err := ioutil.ReadFile(absPath) //Read File of Input
	if err != nil {
		fmt.Print(err)
	}
	d := string(b)
	return d
}
func CurrentWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
func StringToTwoRows(inputString string) (int, []string, []string) {
	stringArray := strings.Split(TrimSpaceNewlineInString(inputString), " ") //Split " " into Slices

	lengthOfArray := StringToInt(stringArray[0])
	stringArray = stringArray[1:] //Drops first value as this is just the length

	var leftSlice []string
	var rightSlice []string

	for i := 0; i < lengthOfArray*2-1; i += 2 {
		leftSlice = append(leftSlice, stringArray[i])
		rightSlice = append(rightSlice, stringArray[i+1])
	}
	return lengthOfArray, leftSlice, rightSlice
}

func StringToThreeRows(inputString string) (int, []string, []string, []string) {
	stringArray := strings.Split(TrimSpaceNewlineInString(inputString), " ") //Split " " into Slices
	lengthOfArray := StringToInt(stringArray[0])
	stringArray = stringArray[1:] //Drops first value as this is just the length

	var leftSlice []string
	var middleSlice []string
	var rightSlice []string

	for i := 0; i < lengthOfArray*3-1; i += 3 {
		leftSlice = append(leftSlice, stringArray[i])
		middleSlice = append(middleSlice, stringArray[i+1])
		rightSlice = append(rightSlice, stringArray[i+2])
	}
	return lengthOfArray, leftSlice, middleSlice, rightSlice
}
func StringToXRows(inputString string) (firstRow []string, outputArray [][]string) {
	stringArray := TrimSpaceNewlineInStringToArray(inputString)
	firstRow = stringArray[:1]
	stringArray = stringArray[1:]
	columnLen := len(stringArray)
	for c := 0; c < columnLen; c++ {
		outputArray = append(outputArray, strings.Split(stringArray[c], " "))
	}
	return firstRow, outputArray
}
func Average(numsIn []int) (numOut int) {
	sum := 0
	for _, x := range numsIn {
		sum += x
	}
	numOut = sum / len(numsIn)
	return

}
func SortSlice(sliceIn []int) (passes int, totalSwaps int) {
	swaps := 0
	for ; !(passes != 0 && swaps == 0); passes++ { //&& swaps != 0
		//fmt.Println("Pass number:", passes)
		swaps = 0
		for i := 0; i < len(sliceIn)-1; i++ {
			if sliceIn[i] > sliceIn[i+1] {
				swaps++
				temp := sliceIn[i+1]
				sliceIn[i+1] = sliceIn[i]
				sliceIn[i] = temp
			}
		}
		totalSwaps += swaps
	}
	return
}

func SliceChecksum(sliceIn []int) (result int) { // /sp Slice
	for _, i := range sliceIn {
		result = (result + i) * 113
		result = result % 10000007
	}
	return
}

func SliceStringToInt(xSlice [][]string) [][]int {
	var intSlice [][]int
	for _, row := range xSlice {
		var temp []int
		for _, i := range row {
			temp = append(temp, StringToInt(i))
		}
		intSlice = append(intSlice, temp)
	}
	return intSlice
}
func SliceStringToFloat(xSlice [][]string) (floatSlice [][]float64) {
	for _, row := range xSlice {
		var temp []float64
		for _, i := range row {
			temp = append(temp, StringToFloat(i))
		}
		floatSlice = append(floatSlice, temp)
	}
	return
}
func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', -1, 64)
}
