package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/almaraz333/proxy-api/handlers"
	"github.com/joho/godotenv"
)

func main() {
    mux := http.NewServeMux()

	godotenv.Load()

	PORT := os.Getenv("PORT")

    mux.HandleFunc("/", handlers.RootHandler)

    fmt.Printf("Server Started on port %v\n", PORT)

    err := http.ListenAndServe(PORT, mux)

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("Server Closed")
    } else if err != nil {
        fmt.Printf("Server Error: %v", err.Error())
        os.Exit(1)
    }
}