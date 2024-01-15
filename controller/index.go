package controller

import (
	Golanta "Golanta/fonction"
	InitTemplate "Golanta/templates"
	"log"
	"net/http"
	"strconv"
)

func DataIndex(w http.ResponseWriter, r *http.Request) {
	InitTemplate.Temp.ExecuteTemplate(w, "index", nil)
}

func TreatmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
		return
	}
	newCtn := Golanta.Aventuriers{
		ID:  Golanta.GetIdNewAdventurer(),
		Nom: r.FormValue("Nom"),
		Age: r.FormValue("age"),
	}
	Golanta.AddAventurier(newCtn)

	http.Redirect(w, r, "/profil", http.StatusPermanentRedirect)
}

func TreatmentDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("aventurier"))
	if err != nil {
		log.Fatal("log: deleteArticleTreatment() Atoi error!\n", err)
	}
	Golanta.DeleteAventurier(id)
	http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
}
