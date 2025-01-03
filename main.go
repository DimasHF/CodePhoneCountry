package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed code_phone.json
var codePhoneJSON []byte

func main() {
	http.HandleFunc("/codephone", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(codePhoneJSON)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
