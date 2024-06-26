package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "OOOps", http.StatusBadRequest)
	}
	h.l.Printf("Hello %s", d)

	fmt.Fprintf(rw, "Hello %s", d)
}
