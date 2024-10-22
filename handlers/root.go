package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    client := &http.Client{}

	SBUX_API_URL := os.Getenv("API_URL")

    req, err := http.NewRequest(r.Method, SBUX_API_URL+r.RequestURI, r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Println(req.URL, method)

    for name, values := range r.Header {
        for _, value := range values {
            req.Header.Add(name, value)
        }
    }

    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }

    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}