package database

import (
	"fmt"
	"gogin/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USER   = "root"
	PASS   = "root"
	HOSTDB = "172.20.0.2"
	DBNAME = "go5db"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	stringConexao := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOSTDB, DBNAME)
	DB, err = gorm.Open(mysql.Open(stringConexao))
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados!")
	}
	DB.AutoMigrate(&models.Aluno{})
}
