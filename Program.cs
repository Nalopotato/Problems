﻿
//Multiples();
//Fib();
//PrimeFactor();
//Palindrome();
SmallestMultiple();


// Find the sum of all the multiples of 3 or 5 below 1000
void Multiples()
{
    var sum = 0;

    for (int i = 0; i < 1000; i++)
    {
        if (i % 3 == 0 || i % 5 == 0)
        {
            sum += i;
            Console.WriteLine($"sum: {sum} ___ i: {i}");
        }
    }
}


//Find the sum of the even numbers of the Fibonacci output, under 4 million
void Fib()
{
    var sum = 0;
    var fibOld = 1;
    var fibNew = 1;

    while (fibNew < 4000000)
    {
        fibNew = fibOld + fibNew;

        if (fibNew % 2 == 0)
        {
            sum += fibNew;
            Console.WriteLine($"sum: {sum} ___ fibNew: {fibNew}");
        }

        fibOld = fibNew - fibOld;
    }
}


//What is the largest prime factor of the number 600851475143 -> O(n log log n)
void PrimeFactor()
{
    long result = 1;
    long number = 600851475143;

    for (int i = 2; i * i <= number; i++)
    {
        if (number % i == 0)
        {
            result = i;
            Console.WriteLine(result);
            //Killed the process after 10 minutes. Will take some further understanding to optimize
        }
    }
}


//Find the largest palindrome made from the product of two 3-digit numbers
//Ex: The largest from 2-digit numbers is 91x99 = 9009
//Result: 924x932 = 861168
//I did *NOT* google for any help or answers to this problem, to see what I could come up with
//At first, my result was 999x91, because within the 2nd for loop, I did not have my conditions set properly
void Palindrome()
{
    var result = 0;

    for (var i = 999; i > 0; i--)
    {
        for (var k = i; k <= 999; k++)
        {
            var num = i * k;
            var start = "";
            var end = "";
            var endFlip = "";

            double x = num.ToString().Length / 2; //Even number of chars

            if (num.ToString().Length % 2 > 0) //Has a middle character
            {
                x = num.ToString().Length / 2 + 0.5;
                end = num.ToString().Substring((int)x + 1, (int)x);
            }
            else
            {
                end = num.ToString().Substring((int)x, (int)x);
            }

            start = num.ToString().Substring(0, (int)x);

            foreach (var c in end)
            {
                endFlip = c + endFlip;
            }

            if (start == endFlip)
            {
                result = num;
                Console.WriteLine($"i: {i} __ k: {k} __ result: {result}");
                //Console.WriteLine($"start: {start} __ middle: {middle} __ end: {end}");
                break;
            }
        }

        if (result > 0)
            break;
    }
}


//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
//Ex: 2520 is the smallest number divisible by all numbers from 1 to 10
void SmallestMultiple()
{

}