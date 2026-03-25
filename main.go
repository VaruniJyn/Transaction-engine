package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var cards = make(map[string]Card)
var transactions []Transaction

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	card, exists := cards[req.CardNumber]
	if !exists {
		logTransaction(req, "FAILED")
		json.NewEncoder(w).Encode(TransactionResponse{
			Status:   "FAILED",
			RespCode: "05",
			Message:  "Invalid card",
		})
		return
	}

	if card.Status != "ACTIVE" {
		logTransaction(req, "FAILED")
		json.NewEncoder(w).Encode(TransactionResponse{
			Status:   "FAILED",
			RespCode: "05",
			Message:  "Card blocked",
		})
		return
	}

	if hashPIN(req.Pin) != card.PinHash {
		logTransaction(req, "FAILED")
		json.NewEncoder(w).Encode(TransactionResponse{
			Status:   "FAILED",
			RespCode: "06",
			Message:  "Invalid PIN",
		})
		return
	}

	if req.Amount <= 0 {
		logTransaction(req, "FAILED")
		json.NewEncoder(w).Encode(TransactionResponse{
			Status:   "FAILED",
			RespCode: "07",
			Message:  "Invalid amount",
		})
		return
	}

	if req.Type == "withdraw" {
		if card.Balance < req.Amount {
			logTransaction(req, "FAILED")
			json.NewEncoder(w).Encode(TransactionResponse{
				Status:   "FAILED",
				RespCode: "99",
				Message:  "Insufficient balance",
			})
			return
		}
		card.Balance -= req.Amount
	} else if req.Type == "topup" {
		card.Balance += req.Amount
	} else {
		logTransaction(req, "FAILED")
		json.NewEncoder(w).Encode(TransactionResponse{
			Status:   "FAILED",
			RespCode: "08",
			Message:  "Invalid transaction type",
		})
		return
	}

	cards[req.CardNumber] = card

	logTransaction(req, "SUCCESS")

	json.NewEncoder(w).Encode(TransactionResponse{
		Status:   "SUCCESS",
		RespCode: "00",
		Balance:  card.Balance,
	})
}

func logTransaction(req TransactionRequest, status string) {
	txn := Transaction{
		TransactionID: generateID(),
		CardNumber:    req.CardNumber,
		Type:          req.Type,
		Amount:        req.Amount,
		Status:        status,
		Timestamp:     time.Now().String(),
	}
	transactions = append(transactions, txn)
}

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	cardNumber := r.URL.Path[len("/api/card/balance/"):]

	card, exists := cards[cardNumber]
	if !exists {
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "Card not found",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "SUCCESS",
		"balance": card.Balance,
	})
}

func getTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	cardNumber := r.URL.Path[len("/api/card/transactions/"):]

	var result []Transaction

	for _, txn := range transactions {
		if txn.CardNumber == cardNumber {
			result = append(result, txn)
		}
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	fmt.Println("Server starting...")

	cards["4123456789012345"] = Card{
		CardNumber: "4123456789012345",
		CardHolder: "John Doe",
		PinHash:    hashPIN("1234"),
		Balance:    1000,
		Status:     "ACTIVE",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		card := cards["4123456789012345"]
		w.Write([]byte("Card Holder: " + card.CardHolder))
	})

	http.HandleFunc("/api/card/balance/", getBalanceHandler)
	http.HandleFunc("/api/card/transactions/", getTransactionsHandler)
	http.HandleFunc("/api/transaction", transactionHandler)

	http.ListenAndServe(":8080", nil)
}
