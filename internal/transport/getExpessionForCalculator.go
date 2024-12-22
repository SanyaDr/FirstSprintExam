package transport

import (
	"FirstSprintExam/pkg/calculation"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

// GetExpression Метод чтения выражения и отправка его в Calculation
func GetExpression(w http.ResponseWriter, r *http.Request) {
	var myReq Request
	// Проверяем что получен именно POST метод
	if r.Method != http.MethodPost {
		log.Println("ERROR: получен не POST метод!")
		returnError(w, "Method is not allowed", http.StatusUnprocessableEntity) // http.StatusMethodNotAllowed ?
		return
	}

	// Получаем тело запроса и сам expression
	body, err := ioutil.ReadAll(r.Body)

	log.Printf("body:\n %v", string(body))

	if err != nil {
		log.Printf("ERROR: Ошибка получения данных запроса! Текст ошибки: %v \n ", err)
		returnError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = json.Unmarshal(body, &myReq)
	//TODO Проверь код ошибки, точно ли badRequest?
	if err != nil {
		log.Printf("ERROR: Ошибка обработки JSON! Текст ошибки: %v \n ", err)
		returnError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Отправляем полученный expression в Calculator
	result, err := calculator.Calc(myReq.Expression)
	if err != nil {
		returnError(w, err.Error(), http.StatusUnprocessableEntity)
		log.Printf("ERROR: Ошибка при вычислении! Текст ошибки: %v \n ", err)
		return
	}

	// TODO переделать на HTTP ответ
	//fmt.Fprintf(w, "Answer: %v", result)
	returnAnswer(w, fmt.Sprint(result))
	log.Printf("Получен результат: %v", result)

}
