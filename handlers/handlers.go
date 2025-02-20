package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
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
	resp, err:= http.Get("http://numbersapi.com/"+ strconv.Itoa(number) + "/math")
	if err != nil {
		// If there's an error, send a response with the error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch data from API: %v", err),
		})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// If there's an error reading the response body
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to read API response: %v", err),
		})
		return
	}

	// Send the data from the external API as the response
	
	fact := string(body)
	fmt.Println("the fact of the matter is", fact)


	result := entity.Response2{
		Number: number,
		Prime: isPrime,
		Perfect: isPerfect,
		Digit_sum: digitSum,
		Fun_fact: fact,
	}

	if isArmstrong {
		result.Properties = append(result.Properties, "armstrong")
	}

	result.Properties = append(result.Properties, parity)

	// Return JSON response
	c.JSON(http.StatusOK, result)
}

func Get_symbols(ctx *gin.Context){

	var payload entity.MonitorPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		fmt.Println("⚠️ No payload supplied, proceeding without it...")
	}

	symbols := []string{
        "GBPUSD", "EURJPY",
        "EURUSD", "EURCHF",
        "USDCHF", "EURGBP",
        "USDCAD", "AUDCAD",
	}

	var results strings.Builder
	results.WriteString("📈 **Forex Exchange Rates**\n--------------------------\n")

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, symbol := range symbols {
		wg.Add(1)
		go func(sym string) {
			defer wg.Done()
			base := symbol[:3]  // First 3 letters
			quote := symbol[3:] // Last 3 letters
			var exchangeRate = services.TwelveDemo(base, quote)
			mu.Lock()
			defer mu.Unlock()
			results.WriteString(fmt.Sprintf("%s/%s → 💹 Rate: %s\n", base, quote, exchangeRate)) 
		}(symbol)
	}
	wg.Wait()
	fmt.Println(results.String())
	telex_data := gin.H{
		"message": strings.Join(strings.Split(results.String(), "\n"), "\n"),
		"username": "Samex Forex Update",
        "event_name": "Forex Update",
        "status": "success",
	}
	telex_url := os.Getenv("TELEX_WEBHOOK")
	telresponse, err := services.PostToReturnURL(telex_url, telex_data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(telresponse)

	ctx.JSON(200, telex_data)
}

func Webhook(ctx *gin.Context){
	ctx.JSON(200, entity.Integrationson)
}

//id = 01950b90-b1bf-75b7-b9e6-e831fdd18b5f