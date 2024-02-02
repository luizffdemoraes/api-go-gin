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

### 🛠 Tecnologias

- [GoLang 1.20](https://go.dev/)
- [Gin 1.9.1](https://github.com/gin-gonic/gin)