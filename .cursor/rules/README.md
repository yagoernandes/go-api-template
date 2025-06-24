# Regras do Projeto - Go API Template

Este diretório contém as regras de desenvolvimento para o projeto Go API Template, seguindo a arquitetura Domain-Driven Design (DDD).

## Estrutura das Regras

### 📋 Regras Principais

1. **`ddd-architecture.mdc`** - Arquitetura e separação de camadas
   - **Tipo**: Always Applied
   - **Escopo**: Todos os arquivos Go
   - **Função**: Define a estrutura DDD obrigatória

2. **`go-standards.mdc`** - Padrões da linguagem Go
   - **Tipo**: Always Applied
   - **Escopo**: Todos os arquivos Go
   - **Função**: Convenções de código e organização

3. **`naming-conventions.mdc`** - Convenções de nomenclatura
   - **Tipo**: Always Applied
   - **Escopo**: Domain e Application layers
   - **Função**: Regras de nomeação bilíngue (PT/EN)

### 🎯 Regras Específicas por Camada

4. **`http-handlers.mdc`** - Padrões para handlers HTTP
   - **Tipo**: Auto Attached
   - **Escopo**: `internal/interfaces/http/**/*.go`
   - **Função**: Convenções REST e tratamento de requisições

5. **`repository-pattern.mdc`** - Padrão Repository
   - **Tipo**: Auto Attached
   - **Escopo**: Arquivos de repositório
   - **Função**: Interfaces e implementações de persistência

6. **`service-layer.mdc`** - Camada de serviços
   - **Tipo**: Auto Attached
   - **Escopo**: `internal/application/**/*.go`
   - **Função**: Lógica de aplicação e coordenação

## Como as Regras São Aplicadas

### 🔄 Always Applied (Sempre Ativas)
- `ddd-architecture.mdc`
- `go-standards.mdc`
- `naming-conventions.mdc`

Estas regras estão sempre no contexto do modelo de IA e são aplicadas a qualquer código Go.

### 📎 Auto Attached (Anexadas Automaticamente)
- `http-handlers.mdc` → Quando editando handlers HTTP
- `repository-pattern.mdc` → Quando trabalhando com repositórios
- `service-layer.mdc` → Quando editando serviços de aplicação

Estas regras são incluídas automaticamente quando você edita arquivos nos diretórios correspondentes.

## Benefícios das Regras

### ✅ Consistência
- Todos os desenvolvedores seguem os mesmos padrões
- Código previsível e organizado
- Arquitetura mantida ao longo do tempo

### ✅ Qualidade
- Princípios SOLID aplicados automaticamente
- Separação de responsabilidades clara
- Facilita testes e manutenção

### ✅ Produtividade
- Reduz tempo de decisão sobre estrutura
- Acelera code reviews
- Facilita onboarding de novos desenvolvedores

## Evolução das Regras

### 📝 Para Adicionar Novas Regras:
1. Crie um arquivo `.mdc` neste diretório
2. Defina o tipo apropriado (Always, Auto Attached, etc.)
3. Especifique o escopo com padrões glob
4. Documente exemplos práticos

### 🔄 Para Modificar Regras Existentes:
1. Edite o arquivo `.mdc` correspondente
2. Mantenha consistência com regras relacionadas
3. Atualize exemplos se necessário
4. Teste com código existente

## Uso no Cursor

As regras são automaticamente carregadas pelo Cursor quando você trabalha no projeto. Para referência manual, use:
- `@ddd-architecture` - Para questões de arquitetura
- `@http-handlers` - Para padrões de API
- `@repository-pattern` - Para persistência de dados

---

**Nota**: Estas regras são versionadas junto com o código e evoluem com o projeto. Sempre mantenha-as atualizadas conforme as necessidades do domínio de negócio. 