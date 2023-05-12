package routes

import (
	"gogin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCpf)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/", controllers.PaginaIndex)
	r.NoRoute(controllers.PaginaError)
	r.Run(":5000")
}
