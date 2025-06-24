package domain

// TarefaRepository define o contrato para operações de persistência de tarefas
type TarefaRepository interface {
	FindAll() ([]*Tarefa, error)
	FindByID(id int) (*Tarefa, error)
	Save(tarefa *Tarefa) error
	Update(tarefa *Tarefa) error
	Delete(id int) error
	GetNextID() int
}
