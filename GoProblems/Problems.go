package main

import (
	"fmt"
	"math"
)

//VSCode shortcuts:
//Ctrl + K + 0 ... (Contract)
//Ctrl + K + J ... (Expand)
//Shift + Alt + Down (Duplicate line)

func main() {
	//Multiples()
	//Fib()
	PrimeFactor()
	//Palindrome()
	//SmallestMultiple()
	//SumSquareDiff()

	//SDG()
}

// Find the sum of all the multiples of 3 or 5 below 1000
//Result: 233168
func Multiples() {
	sum := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
			fmt.Printf("sum: %v ___ i: %v \n", sum, i)
		}
	}
}

//Find the sum of the even numbers of the Fibonacci output, under 4 million
//Result: 4613732
func Fib() {
	sum := 0
	fibOld := 1
	fibNew := 1

	for fibNew < 4000000 {
		fibNew = fibOld + fibNew

		if fibNew%2 == 0 {
			sum += fibNew
			fmt.Printf("sum: %v ___ fibNew: %v \n", sum, fibNew)
		}

		fibOld = fibNew - fibOld
	}
}

//What is the largest prime factor of the number 600851475143
func PrimeFactor() {
	var result uint64 = 1
	var number uint64 = 600851475143
	var i uint64

	for i = 2; i*i <= number; i++ {
		/* for number%2 == 0 {
			number /= 2
		} */

		//fmt.Printf("%v ,", i) 775146 -> Sqrt of 600851475143

		if number%i == 0 {
			result = i
			fmt.Println(result)
			//Killed the process after 10 minutes. Will take some further understanding to optimize
		}
	}
}

//Find the largest palindrome made from the product of two 3-digit numbers
//Ex: The largest from 2-digit numbers is 91x99 = 9009
//Result: 924x932 = 861168
func Palindrome() {
	panic("unimplemented")
}

//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
//Ex: 2520 is the smallest number divisible by all numbers from 1 to 10
func SmallestMultiple() {
	panic("unimplemented")
}

//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum
//Result: 25164150
func SumSquareDiff() {
	fmt.Println("SumSquareDiff")
}

func SDG() int {
	x := 600851475143
	c := x
	prime := 2

	// divide original number by 2 as many times as possible to eliminate
	// prime number 2 and all evens
	for c%2 == 0 {
		c /= 2
	}

	// all occurrences of 2 are eliminated; start with the next prime, 3
	// limit is sqrt of x because that is the largest possible factor
	// iterate by 2 to traverse odd numbers since we have eliminated all evens
	for i := 3; i <= int(math.Sqrt(float64(x))); i += 2 {
		for c%i == 0 {
			if i > prime {
				prime = i
			}
			c /= i
		}
	}

	if c > prime {
		prime = c
	}

	fmt.Printf("%v\n", prime)
	return 1
}
