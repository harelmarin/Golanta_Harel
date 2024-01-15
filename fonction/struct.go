package fonction

// Session stores all User info.
type Session struct {
	IsOpen bool
}

type Aventuriers struct {
	ID  int    `json:"id"`
	Nom string `json:"nom"`
	Age string `json:"age"`
}

var Aventurier Aventuriers
