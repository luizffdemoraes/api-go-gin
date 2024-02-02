<h1 align="center">
Go e Gin: criando API rest com simplicidade
</h1>

<h2 align="center">
Instalando e criando a primeira rota com Gin
</h2>

* Iniciamos o projeto do zero, criando uma pasta e instalando os módulos necessários para criar uma API com Gin;

* Instalamos o gin em nosso projeto com o comando go get -u github.com/gin-gonic/gin;

* Criamos um endpoint, que recebe uma requisição GET, retornando um Json exibindo recursos de um aluno.

Obtendo gim
Com suporte ao módulo Go, basta adicionar o seguinte import

```go get -u github.com/gin-gonic/gin```

<h2 align="center">
Modularizando o código, modelos
</h2>

* Modularizamos nosso código criando pacotes de controles e rotas, tornando nosso código mais fácil de manter e editar;

* Aprendemos como pegar informações passadas por parâmetros e retornar uma mensagem personalizada;

* Criamos a struct de aluno, que vamos disponibilizar como recurso da nossa API.

<h2 align="center">
Struct, banco de dados e ORM
</h2>

* Instalamos o Gorm para não escrever código sql, facilitando a comunicação da aplicação com o banco de dados;
* Conectamos a API com banco de dados e realizamos uma migração com base na struct de aluno;
* Alteramos o controle para exibir os alunos registrados no banco de dados!

Comandos para obter o ip para configurar o PostgreSQL.:

```docker-compose exec postgres sh```
```hostname -i``` 

Instalando GORM.:

```go get -u gorm.io/gorm```

Instalando o driver do postgres.:

```go get gorm.io/driver/postgres```

<h2 align="center">
Implementando rotas HTTP
</h2>

* Alteramos o controller para exibir alunos do banco de dados;
* Criamos um endpoint para exibir alunos por ID;
* Alteramos o comportamento da API para exibir uma mensagem quando o ID do aluno não for encontrado;

<h2 align="center">
Deletando, editando e buscando alunos
</h2>

* Adicionamos um endpoint com método Delete para deletar um aluno e removê-lo do banco de dados;
* Adicionamos um endpoint com método Patch para atualizar o cadastro de um aluno;
* Criamos um endpoint para buscar alunos pelo número do CPF;

### 🛠 Tecnologias

- [GoLang 1.20](https://go.dev/)
- [Gin 1.9.1](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/index.html)