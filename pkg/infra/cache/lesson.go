package cache

import (
	"encoding/json"
	"time"

	"github.com/morikuni/failure"

	"iGoStyle/pkg/domain"
	"iGoStyle/pkg/technical/errutil"
)

func NewLessonRepo(c *Cache) *LessonRepo {
	return &LessonRepo{c}
}

type LessonRepo struct {
	memory *Cache
}

func (repo *LessonRepo) SaveByTutorID(tutorIDKey string, lesson interface{}) error {
	_, ok := lesson.(*domain.Lesson)
	if !ok {
		panic("developer error")
	}

	v, _ := json.Marshal(lesson)
	err := repo.memory.Set(tutorIDKey, v)
	if err != nil {
		return failure.Translate(err, errutil.ErrServer)
	}
	return nil
}

func (repo LessonRepo) QueryByTutorID(id domain.TutorID) (*domain.Lesson, time.Duration, error) {
	v, keyRemainTime, err := repo.memory.Get(KeyService{}.TutorID(id))
	if err != nil {
		return nil, 0, failure.Wrap(err)
	}

	lesson := new(domain.Lesson)
	json.Unmarshal(v, lesson)
	return lesson, keyRemainTime, nil
}

func (repo *LessonRepo) SaveAllByTutorIDGroup(tutorIDGroupKey string, lessons interface{}) error {
	_, ok := lessons.([]*domain.Lesson)
	if !ok {
		panic("developer error")
	}

	v, _ := json.Marshal(lessons)
	err := repo.memory.Set(tutorIDGroupKey, v)
	if err != nil {
		return failure.Translate(err, errutil.ErrServer)
	}
	return nil
}

func (repo LessonRepo) QueryAllByTutorIDGroup(ids []domain.TutorID) ([]*domain.Lesson, time.Duration, error) {
	v, keyRemainTime, err := repo.memory.Get(KeyService{}.TutorIDGroup(ids))
	if err != nil {
		return nil, 0, failure.Wrap(err)
	}

	lessons := make([]*domain.Lesson, 0)
	json.Unmarshal(v, &lessons)
	return lessons, keyRemainTime, nil
}
