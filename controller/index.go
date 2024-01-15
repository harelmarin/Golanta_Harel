package controller

import (
	InitTemplate "Golanta/templates"
	"net/http"
)

func DataIndex(w http.ResponseWriter, r *http.Request) {
	InitTemplate.Temp.ExecuteTemplate(w, "index", nil)
}

func TreatmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
		return
	}
}
