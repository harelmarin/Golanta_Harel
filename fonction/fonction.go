package fonction

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	_, b, _, _ = runtime.Caller(0)
	path       = filepath.Dir(b) + "\\"
)
var jsonfile = path + "../content\\data.json"

func RetrieveAventurier() ([]Aventuriers, error) {
	var aventurier []Aventuriers

	data, err := os.ReadFile(jsonfile)

	if len(data) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(data, &aventurier)
	if err != nil {
		return nil, err
	}
	return aventurier, nil
}

// overwrites data.json.
func ChangeAventurier(aventuriers []Aventuriers) {
	data, errJSON := json.Marshal(aventuriers)
	if errJSON != nil {
		log.Fatal("log: addArticle()\t JSON Marshall error!\n", errJSON)
	}
	errWrite := os.WriteFile(jsonfile, data, 0666)
	if errWrite != nil {
		log.Fatal("log: addArticle()\t WriteFile error!\n", errWrite)
	}
}

// returns first unused id in data.json.
func GetIdNewAdventurer() int {
	aventurier, err := RetrieveAventurier()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	var id int
	var idFound bool
	for id = 1; !idFound; id++ {
		idFound = true
		for _, Aventurier := range aventurier {
			if Aventurier.ID == id {
				idFound = false
			}
		}
	}
	id--
	return id
}

// Add an adventurer in Json
func AddAventurier(newCtn Aventuriers) {
	aventurier, err := RetrieveAventurier()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	aventurier = append(aventurier, newCtn)
	ChangeAventurier(aventurier)

}

func DeleteAventurier(id int) {
	aventurier, err := RetrieveAventurier()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	for i, Aventurier := range aventurier {
		if Aventurier.ID == id {
			aventurier = append(aventurier[:i], aventurier[i+1:]...)
		}
	}
	ChangeAventurier(aventurier)
}

func ModifyAventurier(updatedAventurier Aventuriers) {
	aventurier, err := RetrieveAventurier()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	for i, Aventurier := range aventurier {
		if Aventurier.ID == updatedAventurier.ID {
			aventurier[i] = updatedAventurier
		}
	}
	ChangeAventurier(aventurier)
}

func SelectAventurier(id int) (Aventuriers, bool) {
	var aventurier Aventuriers
	aventuriers, err := RetrieveAventurier()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	var ok bool
	for _, singleAventurier := range aventuriers {
		if singleAventurier.ID == id {
			ok = true
			aventurier = singleAventurier
		}
	}
	return aventurier, ok
}

func GetImgTrainer() string {
	ImgPath := []string{
		"/static/img/Aventurier2.png",
		"/static/img/Aventurier3.png",
		"/static/img/Aventurier4.png",
		"/static/img/Aventurier5.png",
		"/static/img/Aventurier6.png",
		"/static/img/Aventurier7.png",
		"/static/img/Aventurier8.png",
	}
	rand.NewSource(time.Now().UnixNano())
	randomindex := rand.Intn(len(ImgPath))

	return ImgPath[randomindex]
}

func GetImgStarter(starter string) string {
	var PathImg string
	switch starter {
	case "Bulbizarre":
		PathImg = "/static/img/Bulbisar.png"

	case "Salamèche":
		PathImg = "/static/img/Salamèche.png"

	case "Carapuce":
		PathImg = "/static/img/Carapuce.png"
	}
	return PathImg
}
