# API de Tarefas - Arquitetura DDD

Esta aplicação foi reestruturada seguindo os princípios do Domain-Driven Design (DDD) e as convenções padrão de estrutura de projetos Go.

## Estrutura do Projeto

```
template-go-api/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação  
├── internal/
│   ├── domain/                  # Camada de Domínio
│   │   ├── tarefa.go           # Entidade Tarefa
│   │   └── repository.go       # Interface do repositório
│   ├── application/             # Camada de Aplicação
│   │   └── tarefa_service.go   # Serviços de negócio
│   ├── infrastructure/         # Camada de Infraestrutura
│   │   └── repository/
│   │       └── in_memory_tarefa_repository.go # Implementação do repositório
│   └── interfaces/             # Camada de Interface
│       └── http/
│           └── tarefa_handler.go # Handlers HTTP
├── api/                        # (diretório existente)
├── go.mod
└── README.md
```

## Camadas da Arquitetura

### Domain (Domínio) 
- **Entidades**: Objetos de negócio com identidade única (`Tarefa`)
- **Interfaces**: Contratos para repositórios e serviços externos
- **Regras de Negócio**: Validações e lógicas fundamentais do domínio

### Application (Aplicação)
- **Serviços de Aplicação**: Orquestram casos de uso (`TarefaService`)
- **Lógica de Aplicação**: Coordenação entre domínio e infraestrutura

### Infrastructure (Infraestrutura)
- **Repositórios**: Implementações concretas para persistência
- **Adaptadores**: Integrações com sistemas externos

### Interfaces (Interfaces)
- **Controllers/Handlers**: Entrada da aplicação (HTTP, CLI, etc.)
- **Apresentação**: Formatação de dados para saída

## Como Executar

### Compilar e executar:
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

### Ou executar diretamente:
```bash
go run cmd/api/main.go
```

A aplicação será iniciada na porta 8081.

## API Endpoints

### Listar todas as tarefas
```bash
curl -X GET http://localhost:8081/tarefas
```

### Criar uma nova tarefa
```bash
curl -X POST http://localhost:8081/tarefas \
  -H "Content-Type: application/json" \
  -d '{"descricao": "Minha nova tarefa"}'
```

### Atualizar uma tarefa
```bash
curl -X PUT http://localhost:8081/tarefas/1 \
  -H "Content-Type: application/json" \
  -d '{"descricao": "Tarefa atualizada"}'
```

### Excluir uma tarefa
```bash
curl -X DELETE http://localhost:8081/tarefas/1
```

### Verificar se a API está funcionando
```bash
curl http://localhost:8081/
```

## Princípios DDD Aplicados

1. **Separação por Camadas**: Cada camada tem responsabilidade específica
2. **Inversão de Dependência**: Domínio não depende de infraestrutura
3. **Injeção de Dependência**: Dependências são injetadas no main
4. **Interfaces**: Uso de contratos para desacoplamento
5. **Entidades Rica**: Tarefa possui comportamentos e validações

## Benefícios da Arquitetura

- **Testabilidade**: Cada camada pode ser testada independentemente
- **Manutenibilidade**: Código organizado e responsabilidades claras  
- **Flexibilidade**: Fácil troca de implementações (ex: trocar repositório em memória por banco de dados)
- **Escalabilidade**: Estrutura preparada para crescimento da aplicação 