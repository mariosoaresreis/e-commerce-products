package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"products/internal/logger"
)

func validateEnvironmentVariables() {
	envProps := []string{
		"ADDRESS",
		"PORT",
		"USER",
		"PASS",
		"URL",
		"DB_PORT",
		"DB",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(
				fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	validateEnvironmentVariables()
	router := mux.NewRouter()
	handlers := ClienteHandlers{service.NewClienteService(domain.NewClienteRepositoryDb(getDbClient()))}
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/clientes", handlers.getAllClientes)
	router.HandleFunc("/clientes/{cliente_id[0-9]+}", handlers.getCliente)
	endereco := os.Getenv("ENDERECO")
	porta := os.Getenv("PORTA")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", endereco, porta), router))
}
