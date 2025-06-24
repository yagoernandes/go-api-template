package domain

// Tarefa representa uma tarefa no domínio da aplicação
type Tarefa struct {
	ID        int    `json:"id"`
	Descricao string `json:"descricao"`
}

// NewTarefa cria uma nova instância de Tarefa
func NewTarefa(descricao string) *Tarefa {
	return &Tarefa{
		Descricao: descricao,
	}
}

// IsValid valida se a tarefa possui dados válidos
func (t *Tarefa) IsValid() bool {
	return t.Descricao != ""
}
