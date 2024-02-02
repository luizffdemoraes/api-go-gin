package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.POST("/alunos/:id", controllers.BuscaAlunoPorId)
	r.Run()
}
