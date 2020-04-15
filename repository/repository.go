package repository

type WorkRepository interface {
	WorkExperiences(WorkExperienceRequest) (WorkExperienceResponse, error)
}

type LanguagesRepository interface {
	KnownLanguages(KnownLanguagesRequest) (KnownLanguagesResponse, error)
}

type ProjectsRepository interface {
	Projects(ProjectsRequest) (ProjectsResponse, error)
}

