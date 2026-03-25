package main

type Card struct {
	CardNumber string
	CardHolder string
	PinHash    string
	Balance    float64
	Status     string
}

type TransactionRequest struct {
	CardNumber string  `json:"cardNumber"`
	Pin        string  `json:"pin"`
	Type       string  `json:"type"`
	Amount     float64 `json:"amount"`
}

type TransactionResponse struct {
	Status   string  `json:"status"`
	RespCode string  `json:"respCode"`
	Message  string  `json:"message,omitempty"`
	Balance  float64 `json:"balance,omitempty"`
}

type Transaction struct {
	TransactionID string
	CardNumber    string
	Type          string
	Amount        float64
	Status        string
	Timestamp     string
}
