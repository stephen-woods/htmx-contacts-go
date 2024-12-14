package main

import (
	"net/http"
)

func Method(methods []string, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		contains := false

		for _, m := range methods {
			if m == req.Method {
				contains = true
				break
			}
		}

		if !contains {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(res, req)
	})
}
