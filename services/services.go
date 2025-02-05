package services

import (
	"fmt"
	"math"
)

// SumOfDigits calculates the sum of a number's digits
func SumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// GetParity returns "even" or "odd" based on the number's parity
func GetParity(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false // Numbers less than or equal to 1 are not prime
	}
	if n == 2 {
		return true // 2 is the only even prime number
	}
	if n%2 == 0 {
		return false // Other even numbers are not prime
	}

	// Check divisibility up to the square root of n
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false // Found a divisor, not prime
		}
	}
	return true // No divisors, it's prime
}

func IsArmstrong(n int) bool {
	// Convert the number to a string to count digits
	strNum := fmt.Sprintf("%d", n)
	numDigits := len(strNum)
	
	// Calculate the sum of each digit raised to the power of numDigits
	var sum int
	temp := n
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	// Return true if the sum equals the original number
	return sum == n
}

func IsPerfectNumber(n int) bool {
	if n <= 1 {
		return false
	}

	// Sum of divisors (excluding the number itself)
	sum := 0
	for i := 1; i <= n/2; i++ { // Only need to check up to n/2
		if n%i == 0 {
			sum += i
		}
	}

	// Check if sum of divisors equals the number
	return sum == n
}
