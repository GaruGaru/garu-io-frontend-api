package models

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Uri         string `json:"uri"`
}
