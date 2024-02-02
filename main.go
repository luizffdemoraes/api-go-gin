package main

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Gui Lima", CPF: "00000000000", RG: "4700000000"},
		{Nome: "Ana Lima", CPF: "11111111111", RG: "4800000000"},
	}
	routes.HandlerRequests()
}
