package main

import (
	"net/http"
	"time"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", getTransactions)

	http.ListenAndServe(":8888", nil)
}

type Transaction struct {
	Title string `json:"title"`
	Amount float32 `json:"amount"`
	Type int `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-13T20:14:53"
	var salaryReceived, _ = time.Parse(layout, "2020-08-23T23:44:03")

	var transactions = Transactions{
		Transaction{
			Title: "Salário",
			Amount: 2800.0,
			Type: 0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}