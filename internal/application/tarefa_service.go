package application

import (
	"errors"

	"github.com/yagoernandes/template-go-api/internal/domain"
)

// TarefaService implementa a lógica de negócio para tarefas
type TarefaService struct {
	repository domain.TarefaRepository
}

// NewTarefaService cria uma nova instância do serviço de tarefas
func NewTarefaService(repository domain.TarefaRepository) *TarefaService {
	return &TarefaService{
		repository: repository,
	}
}

// ListarTarefas retorna todas as tarefas
func (s *TarefaService) ListarTarefas() ([]*domain.Tarefa, error) {
	return s.repository.FindAll()
}

// CriarTarefa cria uma nova tarefa
func (s *TarefaService) CriarTarefa(descricao string) (*domain.Tarefa, error) {
	if descricao == "" {
		return nil, errors.New("descrição da tarefa não pode estar vazia")
	}

	tarefa := domain.NewTarefa(descricao)
	tarefa.ID = s.repository.GetNextID()

	if !tarefa.IsValid() {
		return nil, errors.New("dados da tarefa são inválidos")
	}

	err := s.repository.Save(tarefa)
	if err != nil {
		return nil, err
	}

	return tarefa, nil
}

// AtualizarTarefa atualiza uma tarefa existente
func (s *TarefaService) AtualizarTarefa(id int, descricao string) (*domain.Tarefa, error) {
	if descricao == "" {
		return nil, errors.New("descrição da tarefa não pode estar vazia")
	}

	tarefa, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if tarefa == nil {
		return nil, errors.New("tarefa não encontrada")
	}

	tarefa.Descricao = descricao

	if !tarefa.IsValid() {
		return nil, errors.New("dados da tarefa são inválidos")
	}

	err = s.repository.Update(tarefa)
	if err != nil {
		return nil, err
	}

	return tarefa, nil
}

// ExcluirTarefa remove uma tarefa
func (s *TarefaService) ExcluirTarefa(id int) error {
	tarefa, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}
	if tarefa == nil {
		return errors.New("tarefa não encontrada")
	}

	return s.repository.Delete(id)
}
