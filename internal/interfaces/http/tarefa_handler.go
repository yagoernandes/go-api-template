package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yagoernandes/template-go-api/internal/application"
)

// TarefaHandler gerencia as requisições HTTP relacionadas a tarefas
type TarefaHandler struct {
	service *application.TarefaService
}

// NewTarefaHandler cria uma nova instância do handler de tarefas
func NewTarefaHandler(service *application.TarefaService) *TarefaHandler {
	return &TarefaHandler{
		service: service,
	}
}

// RegisterRoutes registra as rotas HTTP
func (h *TarefaHandler) RegisterRoutes() {
	http.HandleFunc("/", h.handleRoot)
	http.HandleFunc("/tarefas", h.handleTarefas)
	http.HandleFunc("/tarefas/", h.handleTarefaByID)
}

// handleRoot manipula a rota raiz
func (h *TarefaHandler) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo à API de tarefas!"))
}

// handleTarefas manipula operações CRUD em /tarefas
func (h *TarefaHandler) handleTarefas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listarTarefas(w, r)
	case http.MethodPost:
		h.criarTarefa(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

// handleTarefaByID manipula operações em uma tarefa específica
func (h *TarefaHandler) handleTarefaByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da URL (formato: /tarefas/123)
	path := strings.TrimPrefix(r.URL.Path, "/tarefas/")
	if path == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID da tarefa é obrigatório"))
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID da tarefa deve ser um número"))
		return
	}

	switch r.Method {
	case http.MethodPut:
		h.atualizarTarefa(w, r, id)
	case http.MethodDelete:
		h.excluirTarefa(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

// listarTarefas retorna todas as tarefas
func (h *TarefaHandler) listarTarefas(w http.ResponseWriter, r *http.Request) {
	tarefas, err := h.service.ListarTarefas()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro interno do servidor"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefas)
}

// criarTarefa cria uma nova tarefa
func (h *TarefaHandler) criarTarefa(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Descricao string `json:"descricao"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}

	tarefa, err := h.service.CriarTarefa(request.Descricao)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tarefa)
}

// atualizarTarefa atualiza uma tarefa existente
func (h *TarefaHandler) atualizarTarefa(w http.ResponseWriter, r *http.Request, id int) {
	var request struct {
		Descricao string `json:"descricao"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}

	tarefa, err := h.service.AtualizarTarefa(id, request.Descricao)
	if err != nil {
		if err.Error() == "tarefa não encontrada" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefa)
}

// excluirTarefa remove uma tarefa
func (h *TarefaHandler) excluirTarefa(w http.ResponseWriter, r *http.Request, id int) {
	err := h.service.ExcluirTarefa(id)
	if err != nil {
		if err.Error() == "tarefa não encontrada" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
