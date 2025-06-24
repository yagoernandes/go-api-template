package repository

import (
	"errors"

	"github.com/yagoernandes/template-go-api/internal/domain"
)

// InMemoryTarefaRepository implementa TarefaRepository usando armazenamento em memória
type InMemoryTarefaRepository struct {
	tarefas []*domain.Tarefa
	nextID  int
}

// NewInMemoryTarefaRepository cria uma nova instância do repositório em memória
func NewInMemoryTarefaRepository() domain.TarefaRepository {
	return &InMemoryTarefaRepository{
		tarefas: make([]*domain.Tarefa, 0),
		nextID:  1,
	}
}

// FindAll retorna todas as tarefas
func (r *InMemoryTarefaRepository) FindAll() ([]*domain.Tarefa, error) {
	return r.tarefas, nil
}

// FindByID encontra uma tarefa pelo ID
func (r *InMemoryTarefaRepository) FindByID(id int) (*domain.Tarefa, error) {
	for _, tarefa := range r.tarefas {
		if tarefa.ID == id {
			return tarefa, nil
		}
	}
	return nil, nil // Não encontrado
}

// Save salva uma nova tarefa
func (r *InMemoryTarefaRepository) Save(tarefa *domain.Tarefa) error {
	if tarefa == nil {
		return errors.New("tarefa não pode ser nula")
	}

	r.tarefas = append(r.tarefas, tarefa)
	return nil
}

// Update atualiza uma tarefa existente
func (r *InMemoryTarefaRepository) Update(tarefa *domain.Tarefa) error {
	if tarefa == nil {
		return errors.New("tarefa não pode ser nula")
	}

	for i, t := range r.tarefas {
		if t.ID == tarefa.ID {
			r.tarefas[i] = tarefa
			return nil
		}
	}

	return errors.New("tarefa não encontrada para atualização")
}

// Delete remove uma tarefa pelo ID
func (r *InMemoryTarefaRepository) Delete(id int) error {
	for i, tarefa := range r.tarefas {
		if tarefa.ID == id {
			r.tarefas = append(r.tarefas[:i], r.tarefas[i+1:]...)
			return nil
		}
	}
	return errors.New("tarefa não encontrada para exclusão")
}

// GetNextID retorna o próximo ID disponível
func (r *InMemoryTarefaRepository) GetNextID() int {
	id := r.nextID
	r.nextID++
	return id
}
