package main

import (
	"log"
	conection "main/db"
	templates_sistem "main/resources"
	web "main/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	conection.ContectionDB()
	templates_sistem.InitTemplates()
	web.Routes()
	log.Println("Server Runner")
	http.ListenAndServe(":8080", nil)
}
