package cache

import (
	"errors"
	"strconv"
	"strings"

	"github.com/allegro/bigcache/v3"
	"github.com/morikuni/failure"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/technical/errutil"
)

type keyService struct{}

func (keyService) LanguageID(id domain.LanguageID) string {
	return strconv.Itoa(int(id))
}

func (keyService) TutorID(id domain.TutorID) string {
	return strconv.Itoa(int(id))
}

func (keyService) TutorIDGroup(ids []domain.TutorID) string {
	var b strings.Builder
	for _, id := range ids {
		b.WriteString(strconv.Itoa(int(id)))
	}
	return b.String()
}

func handleGetErr(err error) error {
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return failure.Translate(err, errutil.ErrNotFound)
	}
	return failure.Translate(err, errutil.ErrServer)
}
