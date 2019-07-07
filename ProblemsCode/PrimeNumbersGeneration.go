package ProblemsCode

import (
	"CodeAbbey/FunctionsPackage"
	"fmt"
	"math"
)

func PrimeNumbersGeneration() {

	sIn := FunctionsPackage.TextFileIn("\\InputFiles\\inputPrimeNumbersGeneration")
	_, stringSlice := FunctionsPackage.StringToXRows(sIn)
	primes := calcPrimesToN(5000000)

	fmt.Println()
	fmt.Println(primes)
	for _, i := range  stringSlice[0]{
		fmt.Print(primes[FunctionsPackage.StringToInt(i)], " ")
	}
}

func calcPrimesToX(xIn int) map[int]int{
	var primesSlice []int
	for i := 1; i <= xIn; i++{
		primesSlice = append(primesSlice, i)
	}
	for pass := 1; pass <= xIn; pass++{
		for i := 1; i <= xIn; i++{
			if i % pass == 0 && i != pass {
				primesSlice[i - 1] = 0
				}
			}
		}
	primes := make(map[int]int)
	counter := 1
	for i := 1; i < len(primesSlice); i++{
		if primesSlice[i] != 0{
			primes[counter] = primesSlice[i]
			counter ++
		}
	}
	return primes
}

func calcPrimesToN(xIn int) map[int]int{
	primes := make(map[int]int)
	primes[1] = 2
	primeNumber := 1
	previousPercent := 0
	for i := 2; len(primes) < xIn; i++ {
		maxMultiple := int(math.Floor(math.Sqrt(float64(i)))) + 1
		isPrime := true
		for a := 2; a <= maxMultiple && isPrime; a++{
			if i % a == 0{isPrime = false}
		}
		if isPrime{
			primeNumber ++
			primes[primeNumber] = i
		}
		x := int(math.Round(float64(len(primes)) / float64(xIn) * 100))
		if previousPercent != x{
			previousPercent = x
			fmt.Println(previousPercent, "%")
		}
	}
	return primes
}