package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
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

//integration service

type TwelveDataResponse struct {
	Values []struct {
		Close string `json:"close"`
	} `json:"values"`
}

func TwelveDemo(from string, to string) (string){
	demoKey := "3086f380e87b4353a4fd98f1a2c71b42"
	url := fmt.Sprintf("https://api.twelvedata.com/time_series?symbol=%s/%s&interval=1day&apikey=%s", from, to, demoKey)
	fmt.Println("In twelvedemo start")
	var ex string
	if from == to {
		ex = "1"
	} else {
		// Make HTTP GET request
		res, err := http.Get(url)
		if err != nil {
			return err.Error()
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}

		var response TwelveDataResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}

		if len(response.Values) > 0 {
			_, err := fmt.Sscanf(response.Values[0].Close, "%s", &ex)
			if err != nil {
				fmt.Println(err)
				return err.Error()
			}
		} else {
			return "error"
		}
	}
	fmt.Println("In twelvedemo end")
	fmt.Println(ex)

	return ex
}