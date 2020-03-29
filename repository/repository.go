package repository

type WorkRepository interface {
	WorkExperiences(WorkExperienceRequest) (WorkExperienceResponse, error)
}

type LanguagesRepository interface {
	KnownLanguages(KnownLanguagesRequest) (KnownLanguagesResponse, error)
}
