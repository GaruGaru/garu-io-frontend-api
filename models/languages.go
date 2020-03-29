package models

var (
	ExperienceExperienced = "Experienced"
	ExperienceExpert      = "Experienced"
	ExperienceProficient  = "Proficient"
	ExperienceBeginner    = "Beginner"
)

type Languages struct {
	Name       string `json:"name"`
	Experience string `json:"experience"`
	Level      int    `json:"level"`
}
