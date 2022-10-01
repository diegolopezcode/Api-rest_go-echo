package models

type RequestSession struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Usuario struct {
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
	Username string `json:"username"`
	Password string `json:"password"`
}
