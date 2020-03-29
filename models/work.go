package models

type Company struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Logo    string `json:"logo"`
}

type WorkExperience struct {
	Role        string  `json:"role"`
	StartYear   int     `json:"start_year"`
	EndYear     int     `json:"end_year"`
	Description string  `json:"description"`
	Company     Company `json:"company"`
}
