package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CalculationRequest struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
}

func calculate(w http.ResponseWriter, r *http.Request) {
	var req CalculationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var res CalculationResponse
	switch r.URL.Path {
	case "/add":
		res.Result = add(req.Num1, req.Num2)
	case "/subtract":
		res.Result = subtract(req.Num1, req.Num2)
	case "/multiply":
		res.Result = multiply(req.Num1, req.Num2)
	case "/divide":
		result, err := divide(req.Num1, req.Num2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Result = result
	default:
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/add", calculate)
	http.HandleFunc("/subtract", calculate)
	http.HandleFunc("/multiply", calculate)
	http.HandleFunc("/divide", calculate)

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
