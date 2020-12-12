package cache

import (
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/morikuni/failure"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/technical/errutil"
)

func NewLessonRepo(expiredTime time.Duration) *LessonRepo {
	config := bigcache.DefaultConfig(expiredTime)
	cache, _ := bigcache.NewBigCache(config)
	return &LessonRepo{cache}
}

type LessonRepo struct {
	memory *bigcache.BigCache
}

func (repo *LessonRepo) SaveLessonByTutorID(id domain.TutorID, lesson domain.Lesson) error {
	v, _ := json.Marshal(lesson)
	err := repo.memory.Set(keyService{}.TutorID(id), v)
	return failure.Translate(err, errutil.ErrServer)
}

func (repo LessonRepo) QueryByTutorID(id domain.TutorID) (domain.Lesson, error) {
	v, err := repo.memory.Get(keyService{}.TutorID(id))
	if err != nil {
		return domain.Lesson{}, failure.Wrap(handleGetErr(err))
	}

	lesson := new(domain.Lesson)
	json.Unmarshal(v, lesson)
	return *lesson, nil
}

func (repo *LessonRepo) SaveAllByTutorIDGroup(ids []domain.TutorID, lessons []*domain.Lesson) error {
	v, _ := json.Marshal(lessons)
	err := repo.memory.Set(keyService{}.TutorIDGroup(ids), v)
	return failure.Translate(err, errutil.ErrServer)
}

func (repo LessonRepo) QueryAllByTutorIDGroup(ids []domain.TutorID) ([]*domain.Lesson, error) {
	v, err := repo.memory.Get(keyService{}.TutorIDGroup(ids))
	if err != nil {
		return nil, failure.Wrap(handleGetErr(err))
	}

	lessons := make([]*domain.Lesson, 0)
	json.Unmarshal(v, &lessons)
	return lessons, nil
}
