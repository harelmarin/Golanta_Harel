package routeur

import (
	"Golanta/controller"
	"net/http"
	"os"
)

func Serveur() {
	http.HandleFunc("/index", controller.DataIndex)
	http.HandleFunc("/treatment", controller.TreatmentIndex)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
