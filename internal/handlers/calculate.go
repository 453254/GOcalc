package handlers

import (
	"encoding/json"
	"net/http"

	"calculator/internal/calculator"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}

	result, err := calculator.Calculate(req.Expression)
	if err != nil {
		if err == calculator.ErrInvalidExpression {
			sendError(w, "Expression is not valid", http.StatusUnprocessableEntity)
			return
		}
		sendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	sendResponse(w, CalculateResponse{Result: result})
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(CalculateResponse{Error: message})
}

func sendResponse(w http.ResponseWriter, response CalculateResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
