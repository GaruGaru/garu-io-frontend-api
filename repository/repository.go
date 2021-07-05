package repository

import "context"

type WorkRepository interface {
	WorkExperiences(WorkExperienceRequest) (WorkExperienceResponse, error)
}

type LanguagesRepository interface {
	KnownLanguages(KnownLanguagesRequest) (KnownLanguagesResponse, error)
}

type ProjectsRepository interface {
	Projects(context.Context, ProjectsRequest) (ProjectsResponse, error)
}
