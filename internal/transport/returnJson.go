package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Методы для возврата ответа в виде JSON
/*
Формат ответа:
{
    "result": "результат выражения"
}

Формат ошибки:
{
    "error": "Expression is not valid"
}
*/

type answerResponse struct {
	Result string `json:"result"`
}
type errorResponse struct {
	Error string `json:"error"`
}

// Вывод на экран ответ в виде JSON
func returnAnswer(w http.ResponseWriter, text string) {
	resp := answerResponse{Result: text}
	rData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(rData))
}

// Вывод на экран ошибки в формате JSON
func returnError(w http.ResponseWriter, text string, statusCode int) {
	resp := errorResponse{Error: text}
	rData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Error(w, string(rData), statusCode)
}
