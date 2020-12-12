package cache

import (
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/morikuni/failure"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/technical/errutil"
)

func NewTutorRepo(expiredTime time.Duration) *TutorRepo {
	cfg := bigcache.DefaultConfig(expiredTime)
	cache, _ := bigcache.NewBigCache(cfg)
	return &TutorRepo{memory: cache}
}

type TutorRepo struct {
	memory *bigcache.BigCache
}

func (repo *TutorRepo) SaveTutor(tutorSlug string, tutor domain.Tutor) error {
	v, _ := json.Marshal(tutor)
	err := repo.memory.Set(tutorSlug, v)
	return failure.Translate(err, errutil.ErrServer)
}

func (repo TutorRepo) QueryByTutorSlug(tutorSlug string) (domain.Tutor, error) {
	v, err := repo.memory.Get(tutorSlug)
	if err != nil {
		return domain.Tutor{}, failure.Wrap(handleGetErr(err))
	}

	tutor := new(domain.Tutor)
	json.Unmarshal(v, tutor)
	return *tutor, nil
}

func (repo *TutorRepo) SaveAllByLanguageID(id domain.LanguageID, tutors domain.TutorGroup) error {
	v, _ := json.Marshal(tutors)
	err := repo.memory.Set(keyService{}.LanguageID(id), v)
	return failure.Translate(err, errutil.ErrServer)
}

func (repo TutorRepo) QueryAllByLanguageID(id domain.LanguageID) (domain.TutorGroup, error) {
	v, err := repo.memory.Get(keyService{}.LanguageID(id))
	if err != nil {
		return nil, failure.Wrap(handleGetErr(err))
	}

	tutors := make(domain.TutorGroup)
	json.Unmarshal(v, &tutors)
	return tutors, nil
}
