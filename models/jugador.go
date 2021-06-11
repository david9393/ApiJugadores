package models

type Players struct {
	Players []Player `json:"players"`
}

type Player struct {
	Name         string `json:"name"`
	Position     string `json:"position"`
	Nacionalidad string `json:"nacionalidad"`
	Club         string `json:"club"`
}
