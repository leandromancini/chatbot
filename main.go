package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Handler para el endpoint /ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message": "pong"}`)
	})

	// Handler para cualquier ruta no encontrada
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error": "Endpoint not found"}`)
	})

	// Puerto definido por Netlify en la variable de entorno PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor predeterminado
	}

	fmt.Printf("Server listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
