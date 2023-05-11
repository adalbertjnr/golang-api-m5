package main

import (
	"bytes"
	"encoding/json"
	"gogin/controllers"
	"gogin/database"
	"gogin/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "AlunoTest", Cpf: "12345678901", Rg: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestCheckSaudacao(t *testing.T) {
	r := SetupRotasTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/adal", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")
}

func TestListAll(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos", controllers.ExibeAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestListByCpf(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestListById(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoId)
	pathBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathBusca, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)
	assert.Equal(t, "AlunoTest", alunoMock.Nome)
}

func TestDeletaId(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	r := SetupRotasTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathBusca, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestEditAluno(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "AlunoTest", Cpf: "12345677701", Rg: "123455589"}
	valorJson, _ := json.Marshal(aluno)
	pathEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathEditar, bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var alunoAtualizado models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoAtualizado)
	assert.Equal(t, "AlunoTest", alunoAtualizado.Nome)
	assert.Equal(t, "12345677701", alunoAtualizado.Cpf)
	assert.Equal(t, "123455589", alunoAtualizado.Rg)
}
