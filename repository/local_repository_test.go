package repository

import "testing"

func TestLocalLanguagesRepository_KnownLanguages(t *testing.T) {
	repo := LocalLanguagesRepository{}
	response, err := repo.KnownLanguages(KnownLanguagesRequest{})

	if err != nil {
		t.Fatal(err)
	}

	if len(response.Results) == 0 {
		t.Fatalf("response is empty")
	}

	for _, lang := range response.Results {
		if len(lang.Name) == 0 {
			t.Error("language name is empty")
		}
		if lang.Level < 0 {
			t.Error("language knowledge level is negative")
		}
		if len(lang.Experience) == 0 {
			t.Error("language experience is empty")
		}
	}
}
