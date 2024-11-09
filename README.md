# NuvemHub

**NuvemHub** é uma aplicação web desenvolvida em Go para gerenciar o cadastro de usuários, cursos, aulas e controle de presença na Escola da Nuvem. O objetivo do projeto é oferecer uma solução prática e escalável para gerenciar informações e presença de alunos e professores.

## Recursos do Projeto

- Cadastro de usuários (alunos e professores).
- CRUD básico com GORM (Create, Read, Update, Delete).
- Endpoints RESTful para manipulação de dados.
- Planejamento para controle de presença com QR Code.

## Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Framework de ORM:** [GORM](https://gorm.io/)
- **Framework para gerenciamento de rotas** [Gin-Gonic](https://gin-gonic.com/)
- **Biblioteca para QR Code:** [Go-qrcode](https://pkg.go.dev/github.com/skip2/go-qrcode)
- **Banco de Dados:** PostgreSQL (ou outro DB relacional configurado)


## Estrutura do Projeto

```plaintext
nuvemhub/
├── cmd/
│   └── main.go        # Arquivo principal da aplicação
├── internal/
│   ├── models/        # Modelos de dados com GORM
│   ├── handlers/      # Handlers para endpoints HTTP
│   └── services/      # Regras de negócios e lógica
├── migrations/        # Scripts de migração do banco de dados
├── README.md          # Documentação do projeto
└── go.mod             # Gerenciamento de dependências