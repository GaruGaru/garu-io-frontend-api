package repository


type LocalLanguagesRepository struct {
}

func (l LocalLanguagesRepository) KnownLanguages(KnownLanguagesRequest) (KnownLanguagesResponse, error) {
	return KnownLanguagesResponse{Results: knownLanguages}, nil
}

type LocalWorkExperiencesRepository struct {
}

func (l LocalWorkExperiencesRepository) WorkExperiences(WorkExperienceRequest) (WorkExperienceResponse, error) {
	return WorkExperienceResponse{Results: workExperiences}, nil
}
