package transportTest

import (
	"FirstSprintExam/internal/transport"
	"bytes"
	"net/http"
	"testing"
)

func GetBodydata(expression string) []byte {
	myReq := `{"expression": "` + expression + `"}`
	return []byte(myReq)
}

// TODO ТЕСТ ХУЙНЯ ПЕРЕДЕЛЫВАЙ
func TestApplication(t *testing.T) {
	target := "/api/v1/calculate"
	tests := []struct {
		expression     string
		wantExpression string
		method         string
		wantBody       string
		err            error
		wantStatusCode int
	}{
		{"2+2*2", "2+2*2", "POST", "", nil, 200},
		{"2*   2  * 2 + 1", "2*2*2+1", "POST", "", nil, 200},
	}
	for ind, tc := range tests {
		t.Run(string(rune(ind)), func(t *testing.T) {
			bodyData := GetBodydata(tc.expression) // Тело запроса в виде строки -> example `{"expression": "2+2*2"}`
			b := bytes.NewReader(bodyData)
			//req, err := http.NewRequest(tc.method, target, b) // Создаем запрос

			resp, err := http.Post(target, "application/json", b) // Вызываем POST
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != tc.wantStatusCode {
				t.Error("WA: StatusCode-err. want:", tc.wantStatusCode, "; got:", resp.StatusCode, ";")
			}
			var gotBody []byte
			_, err = resp.Body.Read(gotBody)
			if err != nil {
				t.Fatal(err)
			}
			if string(gotBody) != tc.wantBody {
				t.Error("WA: body. want:", tc.wantBody, "; got:", string(gotBody))
			}

		})
	}
}

func TestReadPostRequest(t *testing.T) {
	type args struct {
		smth string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transport.ReadPostRequest(tt.args.smth); got != tt.want {
				t.Errorf("ReadPostRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
