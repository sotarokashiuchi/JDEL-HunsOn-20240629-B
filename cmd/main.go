// main
package main

import (
	"fmt"
	"analyze"
)


// 入力を受付、入力された整数の個数に合わせて各解析用関数を呼び出す関数。
func main() {
	var inputNumber []int
	var inputNumberTotal int

	// input
	fmt.Printf("解析する整数の個数を入力してください>>>")
	fmt.Scan(&inputNumberTotal)
	for i := 0; i < inputNumberTotal; i++ {
		var input int
		fmt.Printf("%vつ目の整数を入力してください>>>", i+1)
		fmt.Scan(&input)
		inputNumber = append(inputNumber, input)
	}

	fmt.Println("")
	fmt.Println("Analyze....")
	switch inputNumberTotal {
	case 1:
		analyze.IsOddNumber(inputNumber[0])
		analyze.IsEvenNumber(inputNumber[0])
		analyze.IsGrothendieckPrime(inputNumber[0])
		analyze.IsRepunitNumber(inputNumber[0])
		analyze.IsPrimeNumber(inputNumber[0])
		analyze.IsPerfectNumber(inputNumber[0])
		analyze.IsSquareNumber(inputNumber[0])
		analyze.IsMersenneNumber(inputNumber[0])
		analyze.IsTaxicabNumber(inputNumber[0])
		analyze.IsSophieGermainPrime(inputNumber[0])
		analyze.IsSafePrime(inputNumber[0])
		analyze.IsKaprekarNumber(inputNumber[0])
		analyze.IsFibonacciNumber(inputNumber[0])
	case 2:
		analyze.IsAmicableNumbers(inputNumber[0], inputNumber[1])
		analyze.IsBetrothedNumber(inputNumber[0], inputNumber[1])
		analyze.IsSexyPrimes(inputNumber[0], inputNumber[1])
		analyze.IsTwinPrime(inputNumber[0], inputNumber[1])
		analyze.IsCousinPrimes(inputNumber[0], inputNumber[1])
		analyze.GetLeastCommonMultiple(inputNumber[0], inputNumber[1])
		analyze.GetGreatestCommonDivisor(inputNumber[0], inputNumber[1])
		analyze.IsFriedmanNumber(inputNumber[0], inputNumber[1])
	case 3:
		analyze.IsPythagoreanTriple(inputNumber[0], inputNumber[1], inputNumber[2])
	}
}
