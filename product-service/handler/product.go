package handler

import (
	"log"
	"net/http"
)

// dependencies
type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

func (h *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		h.logger.Println("GET Product")
		Greetings(rw, r)
		return
	}
}

func Greetings(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Greetings from product service ssss\n"))
}
