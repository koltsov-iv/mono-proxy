package main

import (
	"net/http"
	"os"
)

func main() {
	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: nil,
	}
	server.Handler = &MonoHandler{
		apiKey:  os.Getenv("API_KEY"),
		apiUrl:  os.Getenv("API_URL"),
		account: os.Getenv("ACCOUNT"),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
