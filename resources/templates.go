// services/plantillas.go
package templates_sistem

import (
	"log"
	"text/template"
)

var Plantillas *template.Template

func InitTemplates() {
	Plantillas = template.Must(template.ParseGlob("plantillas/**/*.html"))
	log.Println("Plantillas cargadas correctamente")
}
