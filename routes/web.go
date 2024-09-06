package web

import (
	services_user "main/services"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", services_user.Index)
	http.HandleFunc("/crear", services_user.Create)
	http.HandleFunc("/insertar", services_user.Store)
	http.HandleFunc("/delete", services_user.Delete)
	http.HandleFunc("/edit", services_user.Edit)
	http.HandleFunc("/actualizar", services_user.Update)
}
