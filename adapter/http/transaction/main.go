package transaction

import (
	"github.com/07carvalho/dio-expert-finance/model/transaction"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-13T20:14:53"
	var salaryReceived, _ = time.Parse(layout, "2020-08-23T23:44:03")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title: "Salário",
			Amount: 2800.0,
			Type: 0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var body, _ = ioutil.ReadAll(r.Body)
	var res = transaction.Transactions{}
	_ = json.Unmarshal(body, &res)
}
