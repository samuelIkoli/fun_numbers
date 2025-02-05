# fun_numbers
An API that takes a number and returns interesting mathematical properties about it, along with a fun fact.

## 📁 Project Configuration

The project is divided into:

- Handlers: found in `/handlers` folder. The functions that get executed when the endpoints are called is defined here.

- Routes: found in `/routes` directory. URL endpoints and their corresponding method/action.

- Entites: found in `/entity` directory. To ensure conformity of output.

- Services: found in `/services` directory. Little helper functions for larger handlers are organised here.

## Getting Started: Running the Server

### 🔧 Tech Stack

- GoLang
- Gin
- Air

### 📝 Requirements

This project requires go 1.23.3, Gin backend framework for go and OPTIONALLY Air module for live reload

### 💻 Running Locally

1. Clone this repository by running:
   ```bash
   git clone https://github.com/samuelIkoli/fun_numbers
   cd HNG12_BE_0
   ```
2. Install the dependencies:
   ```bash
   go mod tidy
   ```
3. Start the server in dev mode (if you have air installed):
   ```bash
   air
   ```
   or (if you do not want to use air live reload)
   ```bash
   go run main.go
   ```

### 💻 Testing

Online API testing tools such as **Postman** and **Thunderclient** can be used to test the endpoints or just easily make a get request from your browser to the endpoints in the documentation. Or better still, use your terminal and Curl 😈 .

## 📖 Documentation

## Endpoints


- **URL**: `/api/classify-number/?number={number}` where you replace {number} with an actual integer
- **Method**: `GET`
- **Body**: No body but a QUERY in form of an integer should be passed
- **Response**:
  ```json
    {
        "number": 371,
        "is_prime": false,
        "is_perfect": false,
        "properties": ["armstrong", "odd"],
        "digit_sum": 11,  // sum of its digits
        "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371" //gotten from the numbers API
    }
  ```



## 🔗 Link(s)

- [You can interact with the project here](https://fun-numbers.onrender.com/)
- [Backlink] (https://hng.tech/hire/golang-developers)
Built by SAMUEL IKOLI
