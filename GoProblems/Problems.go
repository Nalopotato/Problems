package main

import (
	"fmt"
	"strconv"
	"time"
)

//VSCode shortcuts:
//Ctrl + K + 0 ... (Contract)
//Ctrl + K + J ... (Expand)
//Shift + Alt + Down (Duplicate line)

func main() {
	start := time.Now()

	//Multiples()
	//Fib()
	//PrimeFactor()
	Palindrome()
	//SmallestMultiple()
	//SumSquareDiff()

	fmt.Printf("***** %v ***** \n", time.Since(start))
}

// Find the sum of all the multiples of 3 or 5 below 1000
//Result: 233168.  About 70ms
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
//Result: 4613732.  About 1ms
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
//I still don't fully understand this one, but that's fine
//Result: 71, 839, 1471, 6857.  About 4ms
func PrimeFactor() {
	result := 1
	number := 600851475143
	x := number

	for i := 3; i*i <= number; i += 2 {
		//fmt.Printf("%v ,", i) 775146 -> Sqrt of 600851475143

		if x%i == 0 {
			if i > result {
				result = i
				fmt.Println(result)
			}
			x /= i
		}
	}
}

//Find the largest palindrome made from the product of two 3-digit numbers
//Ex: The largest from 2-digit numbers is 91x99 = 9009
//Result: 924x932 = 861168
func Palindrome() {
	result := 0

	for i := 999; i > 0; i-- {
		for k := i; k <= 999; k++ {
			/*
				timer1 := time.NewTimer(100 * time.Millisecond)
				<-timer1.C
				fmt.Printf("i: %v ___ k: %v \n", i, k)
			*/

			num := i * k
			start := ""
			end := ""
			endFlip := ""
			numStr := strconv.Itoa(num)

			var x float64 = float64(len(numStr)) / 2 //Split in half

			if len(numStr)%2 > 0 {
				x += 0.5              //If it's an odd number of chars, shift right 1
				end = numStr[int(x):] //Start AFTER the middle character, and end at the end of the string
			} else {
				end = numStr[int(x):] //Start at the first character on the right half, and end at the end of the string
			}

			start = numStr[:int(x)]

			for _, c := range end { //Equivalent to "foreach (var c in end)" in C#.  "index" is unused, though, so we use an underscore instead
				endFlip = string(c) + endFlip
			}

			if start == endFlip {
				result = num
				fmt.Printf("i: %v __ k: %v __ result: %v \n", i, k, result)
				break
			}
		}

		if result > 0 {
			break
		}
	}
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
