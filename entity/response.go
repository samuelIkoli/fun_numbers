package entity

type Response struct {
	Email string `json:"email"`
	Current_datetime string `json:"current_datetime"`
	Github_url string `json:"github_url"`
}

type Response2 struct {
	Number int `json:"number"`
	Prime bool `json:"is_prime"`
	Perfect bool `json:"is_perfect"`
	Properties []string `json:"properties"`
	Digit_sum int `json:"digit_sum"`
	Fun_fact string `json:"fun_fact"`
}