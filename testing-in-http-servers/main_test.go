package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPizzasHandler(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		input      *Pizzas //주소값으로 받게 한 이유는? 객체 수정을 하기 위해서
		want       string
		statusCode int
	}{
		{
			name:       "without pizzas",
			method:     http.MethodGet,
			input:      &Pizzas{},
			want:       "Error: No pizzas found",
			statusCode: http.StatusNotFound,
		},
		{
			name:   "with pizzas",
			method: http.MethodGet,
			input: &Pizzas{
				Pizza{
					ID:    1,
					Name:  "Foo",
					Price: 10,
				},
			},
			want:       `[{"id":1,"name":"Foo","price":10}]`,
			statusCode: http.StatusOK,
		},
		{
			name:       "with bad method",
			method:     http.MethodPost,
			input:      &Pizzas{},
			want:       "Method not allowed",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/orders", nil)
			responseRecorder := httptest.NewRecorder()

			//왜 &tc.input을 하지 않았나? - tc.input 변수가 주소를 받고 있기 떄문에
			pizzasHandler{tc.input}.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
