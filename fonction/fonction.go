package fonction

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	fmt.Println("Hello")
}
