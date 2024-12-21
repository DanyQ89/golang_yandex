package main

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type User struct {
	Expression string `json:"expression"`
}

type AnswerGood struct {
	Result string `json:"result"`
}

type AnswerBad struct {
	Error string `json:"error"`
}

func CalculationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		var meow_err AnswerBad
		w.WriteHeader(500)
		meow_err.Error = "Internal server error"
		JsonData, _ := json.Marshal(meow_err)

		w.Write(JsonData)
		return
	}

	res, err := Calc(user.Expression)
	if err != nil {
		var meow_err AnswerBad
		w.WriteHeader(422)
		meow_err.Error = "Expression is not valid"
		JsonData, _ := json.Marshal(meow_err)

		w.Write(JsonData)
		return
	}
	var answer AnswerGood
	answer.Result = fmt.Sprintf("%f", res)
	JsonData, _ := json.Marshal(answer)
	w.Write(JsonData)
}

func main() {
	http.HandleFunc("/api/v1/calculate", CalculationHandler)
	http.ListenAndServe(":8080", nil)

}
