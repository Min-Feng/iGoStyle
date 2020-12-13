package cache

import (
	"strconv"
	"strings"

	"AmazingTalker/pkg/domain"
)

type KeyService struct{}

func (KeyService) TutorSlug(tutorSlug string) string {
	return "TutorSlug:" + tutorSlug
}

func (KeyService) LanguageID(id domain.LanguageID) string {
	return "LanguageID:" + strconv.Itoa(int(id))
}

func (KeyService) TutorID(id domain.TutorID) string {
	return "TutorID:" + strconv.Itoa(int(id))
}

func (KeyService) TutorIDGroup(ids []domain.TutorID) string {
	var b strings.Builder
	b.WriteString("TutorIDGroup:")
	for _, id := range ids {
		b.WriteString(strconv.Itoa(int(id)))
	}
	return b.String()
}
