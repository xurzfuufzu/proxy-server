package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"sync"
)

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var (
	storage   sync.Map
	idCounter int
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func DoProxy(w http.ResponseWriter, r *http.Request) {
	var proxyReq ProxyRequest

	if err := json.NewDecoder(r.Body).Decode(&proxyReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(proxyReq.Method, proxyReq.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	for key, value := range proxyReq.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	proxyResp := ProxyResponse{
		ID:      fmt.Sprintf("%d", generateID()),
		Status:  resp.StatusCode,
		Headers: make(map[string]string),
		Length:  len(body),
	}

	for key, values := range resp.Header {
		if len(values) > 0 {
			proxyResp.Headers[key] = values[0]
		}
	}

	responseBody, err := json.Marshal(proxyResp)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	storage.Store(proxyResp.ID, map[string]interface{}{
		"request":  proxyReq,
		"response": proxyResp,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	value, ok := storage.Load(id)
	if !ok {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func generateID() int {
	idCounter++
	return idCounter
}
