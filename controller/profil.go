package controller

import (
	Golanta "Golanta/fonction"
	InitTemplate "Golanta/templates"
	"net/http"
)

func DataProfil(w http.ResponseWriter, r *http.Request) {
	aventuriers, err := Golanta.RetrieveAventurier()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des aventuriers", http.StatusInternalServerError)
		return
	}
	data := struct {
		Aventuriers []Golanta.Aventuriers
	}{
		Aventuriers: aventuriers,
	}

	InitTemplate.Temp.ExecuteTemplate(w, "profil", data)
}
