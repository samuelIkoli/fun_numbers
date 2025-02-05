package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuelIkoli/fun_numbers/entity"
	"github.com/samuelIkoli/fun_numbers/services"
)

func Test(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "Fun numbers is working",
	})
}

func Task(ctx *gin.Context){
	now:= time.Now().UTC()
	result := entity.Response{
		Email: "ayibanimiikoli@gmail.com",
		Current_datetime: now.Format(time.RFC3339),
		Github_url: "https://github.com/samuelIkoli/HNG12_BE_0",
	}
	ctx.JSON(200, result)
}

func Ping(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "Pong ",
	})
}

func Numerate(c *gin.Context) {
	// Get the 'number' query parameter
	numberStr := c.Query("number")

	// Validate the input
	if numberStr == "" || strings.Contains(numberStr, ".") {
		c.JSON(http.StatusBadRequest, gin.H{"Number": "Invalid integer or empty query", "error": true})
		return
	}

	// Convert string to integer
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Number": "Alphabet or Special charaters are not allowed", "error": true})
		return
	}

	// Compute results
	digitSum := services.SumOfDigits(number)
	parity := services.GetParity(number)
	isPrime := services.IsPrime(number)
	isArmstrong := services.IsArmstrong(number)
	isPerfect := services.IsPerfectNumber(number)
	


	result := entity.Response2{
		Number: number,
		Prime: isPrime,
		Perfect: isPerfect,
		Digit_sum: digitSum,
		Fun_fact: "A good number",
	}

	if isArmstrong {
		result.Properties = append(result.Properties, "armstrong")
	}

	result.Properties = append(result.Properties, parity)

	// Return JSON response
	c.JSON(http.StatusOK, result)
}



