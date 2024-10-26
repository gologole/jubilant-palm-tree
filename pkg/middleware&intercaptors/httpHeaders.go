package main

import (
	"errors"
	"fmt"
	"net/http"
)

/*
Проверяет переданный заголовок с помощью переданной функции,возвращает 401 ,если verificator возвращает false или пустую строку
*/
func CheckHeaderMiddleware(next http.Handler, headerName string, verificator func(string) (bool, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if access, err := verificator(r.Header.Get(headerName)); err != nil || !access {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "401 Unauthorized")
		}

		next.ServeHTTP(w, r)
	})
}

/*пример verificator*/
func CheckAcessToken(headerValue string) (bool, error) {
	if headerValue == "Bearer valid_token" {
		return true, nil
	}
	if headerValue == "" {
		return false, errors.New("AuthHeader value is empty")
	}
	return false, nil
}
