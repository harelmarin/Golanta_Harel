package main

import (
	"Golanta/routeur"
	initTemplate "Golanta/templates"
)

func main() {

	initTemplate.InitTemplate()
	routeur.Serveur()

}
