package main

import (
	"log"
	"net/http"

	"github.com/yagoernandes/template-go-api/internal/application"
	"github.com/yagoernandes/template-go-api/internal/infrastructure/repository"
	httpHandler "github.com/yagoernandes/template-go-api/internal/interfaces/http"
)

func main() {
	// Inicializa as dependências seguindo a inversão de dependência
	tarefaRepository := repository.NewInMemoryTarefaRepository()
	tarefaService := application.NewTarefaService(tarefaRepository)
	tarefaHandler := httpHandler.NewTarefaHandler(tarefaService)

	// Registra as rotas
	tarefaHandler.RegisterRoutes()

	// Inicia o servidor
	log.Println("Servidor iniciado na porta 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
