package cache

import (
	"encoding/json"
	"time"

	"github.com/morikuni/failure"

	"iGoStyle/pkg/domain"
	"iGoStyle/pkg/technical/errutil"
)

func NewTutorRepo(c *Cache) *TutorRepo {
	return &TutorRepo{memory: c}
}

type TutorRepo struct {
	memory *Cache
}

func (repo *TutorRepo) SaveByTutorSlug(tutorSlugKey string, tutor interface{}) error {
	_, ok := tutor.(*domain.Tutor)
	if !ok {
		panic("developer error")
	}

	v, _ := json.Marshal(tutor)
	err := repo.memory.Set(tutorSlugKey, v)
	if err != nil {
		return failure.Translate(err, errutil.ErrServer)
	}
	return nil
}

func (repo TutorRepo) QueryByTutorSlug(tutorSlug string) (*domain.Tutor, time.Duration, error) {
	v, keyRemainTime, err := repo.memory.Get(KeyService{}.TutorSlug(tutorSlug))
	if err != nil {
		return nil, 0, failure.Wrap(err)
	}

	tutor := new(domain.Tutor)
	json.Unmarshal(v, tutor)
	return tutor, keyRemainTime, nil
}

func (repo *TutorRepo) SaveAllByLanguageID(languageIDKey string, tutors interface{}) error {
	_, ok := tutors.(domain.TutorGroup)
	if !ok {
		panic("developer error")
	}

	v, _ := json.Marshal(tutors)
	err := repo.memory.Set(languageIDKey, v)
	if err != nil {
		return failure.Translate(err, errutil.ErrServer)
	}
	return nil
}

func (repo TutorRepo) QueryAllByLanguageID(id domain.LanguageID) (domain.TutorGroup, time.Duration, error) {
	v, keyRemainTime, err := repo.memory.Get(KeyService{}.LanguageID(id))
	if err != nil {
		return nil, 0, failure.Wrap(err)
	}

	tutors := make(domain.TutorGroup)
	json.Unmarshal(v, &tutors)
	return tutors, keyRemainTime, nil
}
