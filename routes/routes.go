package routes

import (
	"gogin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCpf)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.Run(":5000")
}
