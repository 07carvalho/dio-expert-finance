package http

import (
	"github.com/07carvalho/dio-expert-finance/adapter/http/transaction"
	"github.com/07carvalho/dio-expert-finance/adapter/http/actuator"
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)
	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8888", nil)
}
