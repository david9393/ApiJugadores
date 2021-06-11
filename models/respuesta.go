package models

type Respuesta struct {
	TotalPages int     `json:"totalPages"`
	Items      []Items `json:"items"`
}
type Items struct {
	CommonName string `json:"commonName"`
	Position   string `json:"position"`
	Nation     Nation `json:"nation"`
	Club       Club   `json:"club"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
}

type Nation struct {
	Name string `json:"name"`
}
type Club struct {
	Name string `json:"name"`
}
