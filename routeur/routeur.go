package routeur

import (
	"Golanta/controller"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func Serveur() {
	http.HandleFunc("/index", controller.DataIndex)
	http.HandleFunc("/treatment/index", controller.TreatmentIndex)
	http.HandleFunc("/profil", controller.DataProfil)
	http.HandleFunc("/deleteaventurier", controller.TreatmentDelete)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	runServer()
}

func runServer() {
	port := "localhost:8080"
	url := "http://" + port + "/index"
	go http.ListenAndServe(port, nil)
	fmt.Println("Server is running...")
	time.Sleep(time.Second * 3)
	cmd := exec.Command("explorer", url)
	cmd.Run()
	fmt.Println("If the navigator didn't open on its own, just go to ", url, " on your browser.")
	isRunning := true
	for isRunning {
		fmt.Println("If you want to end the server, type 'stop' here :")
		var command string
		fmt.Scanln(&command)
		if command == "stop" {
			isRunning = false
		}
	}
}
