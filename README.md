# Go User Control: API Poderosa para Vagas de Emprego em Go

Bem-vindo ao **Gopportunities**, um projeto que demonstra a construção de uma **API RESTful completa e robusta** em **Go (Golang)**. Esta API é projetada para gerenciar vagas de emprego (Job Openings), oferecendo funcionalidades completas de **CRUD (Criar, Ler, Atualizar, Deletar)** e uma documentação interativa detalhada.

## Tecnologias Utilizadas

Este projeto faz uso de um conjunto poderoso de tecnologias para garantir escalabilidade, performance e facilidade de desenvolvimento:

*   **Go (Golang):** A linguagem de programação principal, escolhida por sua eficiência, concorrência e tipagem forte.
*   **Gin Gonic (Gin):** Um **framework web** leve e de alto desempenho para Go, utilizado para a construção da API. O Gin simplifica o roteamento, a manipulação de requisições e a geração de respostas JSON.
*   **GORM:** Um **ORM (Object-Relational Mapper)** abrangente para Go, que facilita a interação com o banco de dados.
*   **SQLite:** Um banco de dados leve e embarcado, utilizado para a persistência de dados localmente no projeto. Ideal para desenvolvimento e testes.
*   **Swagger (via `gin-swagger` e `swag`):** Ferramentas essenciais para a **documentação automática e interativa da API**. Com comentários no código, o Swagger gera uma interface web que permite explorar e testar os endpoints.
*   **Makefile:** Utilizado para **automatizar tarefas** comuns do projeto, como rodar a aplicação, compilar, executar testes e gerar a documentação.
*   **Logger Personalizado:** Uma implementação de logger configurável para o monitoramento e depuração da aplicação.
*   **Postman:** Ferramenta utilizada para testar os endpoints da API.
*   **Go Modules:** Sistema de gerenciamento de dependências do Go, garantindo a organização e a reprodutibilidade do projeto.
*   **VS Code Go Extension:** Extensão oficial que oferece suporte robusto ao desenvolvimento Go, incluindo auto-completar, formatação e importação automática de pacotes.

## Funcionalidades Principais da API

A API Gopportunities oferece os seguintes endpoints para gerenciamento de vagas:

*   **Criação de Vagas:** Permite o registro de novas vagas de emprego, incluindo campos como cargo, empresa, localização, se é remoto e salário. Implementa validação robusta dos dados de entrada.
*   **Listagem de Vagas:** Retorna uma lista de todas as vagas de emprego cadastradas.
*   **Exibição de Vaga Específica:** Permite buscar os detalhes de uma única vaga pelo seu ID.
*   **Atualização de Vagas:** Suporta a modificação de campos de vagas existentes, com validação de parâmetros.
*   **Deleção de Vagas:** Remove vagas específicas da base de dados.

## Estrutura do Projeto

O projeto é organizado de forma modular, seguindo as melhores práticas de Go para grandes aplicações:

*   **`main.go`**: Ponto de entrada da aplicação, responsável por inicializar as configurações e o roteador.
*   **`router/`**: Contém a lógica de roteamento do Gin, definindo os endpoints e agrupando-os por versão (`/api/v1`).
*   **`handlers/`**: Abriga as funções de manipulação de requisições (handlers) para cada endpoint da API. Essas funções são exportadas e encapsulam a lógica de negócio.
    *   **`request/` e `response/`**: Subdiretórios dentro de `handlers` que definem as estruturas de dados (structs) esperadas nas requisições e respostas, incluindo validações e formatação JSON.
*   **`schemas/`**: Define os modelos de dados (structs) das entidades do banco de dados, como `Opening`, utilizados pelo GORM.
*   **`config/`**: Gerencia as configurações globais da aplicação, incluindo a inicialização do banco de dados (SQLite) e a instância do logger personalizado.

## Como Rodar o Projeto

Siga os passos abaixo para configurar e executar a API Gopportunities:

1.  **Clone o Repositório:**
    ```bash
    git clone https://github.com/Arthur404dev/Gopportunities.git
    cd Gopportunities
    ```
2.  **Inicialize os Módulos Go:**
    ```bash
    go mod init github.com/Arthur404dev/Gopportunities
    ```
    *(Caso já exista, certifique-se de que o `go.mod` está configurado corretamente com o nome do módulo.)*
3.  **Baixe as Dependências:**
    ```bash
    go mod tidy
    ```
    *(Este comando irá baixar automaticamente todas as dependências do projeto, incluindo Gin, GORM e o driver SQLite, e gerará o `go.sum`.)*
4.  **Gere a Documentação Swagger:**
    ```bash
    make docs
    ```
    *(Certifique-se de ter a ferramenta `swag` instalada globalmente (`go install github.com/swaggo/swag/cmd/swag@latest`) e seu `GOPATH/bin` adicionado ao `PATH` do sistema.)*
5.  **Execute a Aplicação:**
    ```bash
    make run
    ```
    A API será iniciada na porta `8080`.

## Acessando a Documentação da API

Após iniciar a aplicação, você pode acessar a documentação interativa do Swagger em seu navegador:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---