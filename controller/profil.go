package controller

import (
	Golanta "Golanta/fonction"
	InitTemplate "Golanta/templates"
	"net/http"
)

func DataProfil(w http.ResponseWriter, r *http.Request) {
	data := Golanta.RetrieveAventurier
	InitTemplate.Temp.ExecuteTemplate(w, "profil", data)
}
