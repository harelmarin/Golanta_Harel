package fonction

// Session stores all User info.
type Session struct {
	IsOpen bool
}

type Aventuriers struct {
	ID         int    `json:"id"`
	Nom        string `json:"nom"`
	Age        string `json:"age"`
	ImgTrainer string `json:"imgtrainer"`
	ImgStarter string `json:"imgstarter"`
	Date       string `json:"date"`
}

var Aventurier Aventuriers
