
# SP001LOGIN - API Documentation

## Visão Geral

Este repositório contém a implementação de uma API para autenticação e gerenciamento de uma lista negra (blacklist). A API é composta por dois serviços principais: **Login** e **Blacklist**. Cada serviço é projetado seguindo princípios de modularidade e separação de responsabilidades, utilizando pacotes internos para repositórios, serviços, e filas.

---

## Serviço de Login

O serviço de **Login** autentica usuários e retorna um token válido para futuras requisições. Ele também verifica se o usuário está em uma blacklist e registra atividades em uma fila SQS.

### Principais Funcionalidades

- **Autenticação de Usuários**
  - Verifica credenciais do usuário no banco de dados.
  - Valida a senha usando hash criptográfico.
- **Geração de Tokens**
  - Cria tokens JWT usando o algoritmo Paseto.
  - Registra o histórico de tokens no banco de dados.
- **Integração com Filas SQS**
  - Envia informações do usuário autenticado para uma fila AWS SQS.

### Fluxo do Serviço

1. Verifica as credenciais do usuário.
2. Valida se o usuário está na blacklist.
3. Gera um token JWT se a autenticação for bem-sucedida.
4. Envia informações do usuário para uma fila SQS.

### Interface do Serviço

#### `LoginService`

```go
LoginService(ctx context.Context, data LoginRequest) (LoginResponse, error)
```

- **Parâmetros**:
  - `ctx`: Contexto para controle de tempo e cancelamento.
  - `data`: Objeto contendo login e senha do usuário.
- **Retorno**:
  - `LoginResponse`: Dados do usuário autenticado e token gerado.
  - `error`: Erro, se houver.

---

## Serviço de Blacklist

O serviço de **Blacklist** gerencia a criação e atualização de entradas na lista negra. Ele é usado para restringir o acesso de determinados usuários ao sistema.

### Principais Funcionalidades

- **Criar Entradas na Blacklist**
  - Adiciona usuários à lista negra com base em informações fornecidas.
- **Atualizar Entradas na Blacklist**
  - Modifica o status de entradas existentes.

### Fluxo do Serviço

1. Recebe os dados do usuário ou ID para processamento.
2. Executa a operação de criação ou atualização no banco de dados.

### Interface do Serviço

#### `CreateBlackList`

```go
CreateBlackList(ctx context.Context, data BlackListRequestCreate) (string, error)
```

- **Parâmetros**:
  - `ctx`: Contexto para controle de tempo e cancelamento.
  - `data`: Objeto contendo informações para criar a entrada na blacklist.
- **Retorno**:
  - `string`: Mensagem de sucesso.
  - `error`: Erro, se houver.

#### `UpdateBlackList`

```go
UpdateBlackList(ctx context.Context, id int64) (string, error)
```

- **Parâmetros**:
  - `ctx`: Contexto para controle de tempo e cancelamento.
  - `id`: Identificador único da entrada na blacklist.
- **Retorno**:
  - `string`: Mensagem de sucesso.
  - `error`: Erro, se houver.

---

## Exemplos de Uso

### Serviço de Login

#### Requisição de Login

```json
POST /login
{
  "login": "user@example.com",
  "password": "password123",
  "accessKey": 1234
}
```

#### Resposta de Sucesso

```json
{
  "id": "1",
  "name": "User Name",
  "email": "user@example.com",
  "role": "admin",
  "nickname": "user_nick",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "firstAccess": false
}
```

---

### Serviço de Blacklist

#### Criar Entrada na Blacklist

```json
POST /black-list/create
{
  "userId": 1,
  "reason": "Fraudulent activity"
}
```

#### Resposta de Sucesso

```json
{
  "message": "Blacklist created successfully"
}
```

#### Atualizar Entrada na Blacklist

```json
PUT /black-list/update/1
```

#### Resposta de Sucesso

```json
{
  "message": "Blacklist updated successfully"
}
```

---

## Configurações do Projeto

### Arquivo `.env`

Certifique-se de configurar variáveis de ambiente necessárias, como:

```bash
SERVER_NAME=your_server_name
SERVER_PORT=:your_port
ENVIRONMENT=your_environment

AWS_ACCESS_KEY=your_access_key
AWS_QUEUE_URL=your_queue_url
AWS_REGION=your_region
AWS_SECRET_KEY=your_secret_key

DB_DATABASE=your_database
DB_DRIVER=your_driver
DB_HOST=your_host
DB_PASSWORD=your_password
DB_PORT=your_database_port
DB_SSL_MODE=?sslmode=your_mode
DB_USER=your_user

SIGNATURE_STRING=your_secret_token
```

## Como Executar o Projeto

Este projeto utiliza `make` para simplificar tarefas de configuração e execução. Certifique-se de ter as dependências necessárias instaladas, como `docker`, `make`, e o Go instalado na sua máquina.

### Passo a Passo

1. **Configurar o Banco de Dados PostgreSQL**

   Para inicializar e configurar o banco de dados PostgreSQL com Docker:
   ```bash
   make postgres-setup
   ```

   Para iniciar ou parar o container do PostgreSQL posteriormente:
   ```bash
   make start-postgres
   make stop-postgres
   ```

   Para criar o banco de dados configurado no `.env`:
   ```bash
   make createdb
   ```

2. **Rodar Migrações no Banco de Dados**

   Para aplicar as migrações no banco de dados:
   ```bash
   make migration-up
   ```

   Para desfazer as migrações aplicadas:
   ```bash
   make migration-down
   ```

3. **Gerar Arquivos Necessários**

   Para gerar arquivos do SQLC:
   ```bash
   make sqlc
   ```

   Para gerar a documentação Swagger:
   ```bash
   make swag
   ```

4. **Executar o Projeto**

   Para executar o projeto localmente:
   ```bash
   make run
   ```

