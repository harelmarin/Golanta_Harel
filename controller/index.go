package controller

import (
	Golanta "Golanta/fonction"
	InitTemplate "Golanta/templates"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func DataIndex(w http.ResponseWriter, r *http.Request) {
	InitTemplate.Temp.ExecuteTemplate(w, "index", nil)
}

func TreatmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
	newCtn := Golanta.Aventuriers{
		ID:         Golanta.GetIdNewAdventurer(),
		Nom:        r.FormValue("Nom"),
		Age:        r.FormValue("age"),
		ImgTrainer: Golanta.GetImgTrainer(),
		ImgStarter: Golanta.GetImgStarter(r.FormValue("starter")),
		Date:       time.Now().Format("02/01/2006"),
	}
	Golanta.AddAventurier(newCtn)

	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func TreatmentDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("aventurier"))
	if err != nil {
		log.Fatal("log: TreatmentDelete() Atoi error!\n", err)
	}
	Golanta.DeleteAventurier(id)
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func DataModify(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("aventurier"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	aventurier, ok := Golanta.SelectAventurier(id)
	if !ok {
		fmt.Println("Error ")
		return
	}

	InitTemplate.Temp.ExecuteTemplate(w, "modify", aventurier)
}

func TreatmentModify(w http.ResponseWriter, r *http.Request) {
	idInForm, err := strconv.Atoi(r.URL.Query().Get("aventurier"))
	if err != nil {
		log.Fatal("log: TreatmentModify() Atoi error!\n", err)
	}
	newCtn := Golanta.Aventuriers{
		ID:         idInForm,
		Nom:        r.FormValue("Nom"),
		Age:        r.FormValue("age"),
		ImgTrainer: Golanta.GetImgTrainer(),
		ImgStarter: Golanta.GetImgStarter(r.FormValue("starter")),
		Date:       time.Now().Format("02/01/2006"),
	}

	Golanta.ModifyAventurier(newCtn)
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}
