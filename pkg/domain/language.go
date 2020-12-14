package domain

type LanguageLookupFormRepo interface {
	LanguageLookupForm() (LanguageLookupForm, error)
}

type LanguageLookupForm struct {
	MapSlugToID map[LanguageSlug]LanguageID
}

func (t LanguageLookupForm) SlugToID(slug LanguageSlug) LanguageID {
	return t.MapSlugToID[slug]
}

type LanguageID int
type LanguageSlug string
