package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
)

type DateTime struct {
	days int
	hours int
	minutes int
	seconds int
	totalSeconds int

}
func ModuloAndTimeDifference() {
	inputString := FunctionsPackage.TextFileIn("\\InputFiles\\inputModuloAndTimeDifference")
	_, dataSlices := FunctionsPackage.StringToXRows(inputString)

	//fmt.Println(dataSlices)
	var leftDateTimes []DateTime
	var rightDateTimes []DateTime
	var resultDateTimes []DateTime


	for _, line := range dataSlices {
		var leftTemp []int
		var rightTemp []int
		for i := 0; i < len(line) / 2; i ++{
			leftTemp = append(leftTemp, FunctionsPackage.StringToInt(line[i]))
			rightTemp = append(rightTemp, FunctionsPackage.StringToInt(line[i + (len(line) / 2)]))
		}
		leftDateTimes = append(leftDateTimes, *NewDateTime(leftTemp...))
		rightDateTimes = append(rightDateTimes, *NewDateTime(rightTemp...))
		resultDateTimes = append(resultDateTimes, *NewDateTime(calcDifference(*NewDateTime(leftTemp...), *NewDateTime(rightTemp...))))
	}
	//fmt.Println(resultDateTimes)
	for _, dataTime := range resultDateTimes{
		fmt.Print("(", dataTime.days," ", dataTime.hours, " ", dataTime.minutes, " ", dataTime.seconds, ") ")
	}


}
func NewDateTime(intSlice ...int) *DateTime{
	temp := new(DateTime)
	if len(intSlice) == 4{
		temp.days = intSlice[0]
		temp.hours = intSlice[1]
		temp.minutes = intSlice[2]
		temp.seconds = intSlice[3]
		temp.totalSeconds = calcSeconds(temp)
		return temp
	} else if len(intSlice) == 1 {
		remainder := 0
		temp.days = intSlice[0] / 86400
		remainder = intSlice[0] % 86400
		temp.hours = remainder / 3600
		remainder = remainder % 3600
		temp.minutes = remainder / 60
		remainder = remainder % 60
		temp.seconds = remainder
		temp.totalSeconds = intSlice[0]
		return temp
	}
	return temp

}
func calcSeconds(dateTimeIn *DateTime) int{
	total := dateTimeIn.seconds
	total += dateTimeIn.days * 86400
	total += dateTimeIn.hours * 3600
	total += dateTimeIn.minutes * 60
	return total
}
func calcDifference (time DateTime, time2 DateTime) int {
	return time2.totalSeconds - time.totalSeconds
}