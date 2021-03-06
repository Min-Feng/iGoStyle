package repository

import (
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/singleflight"

	"iGoStyle/pkg/domain"
	"iGoStyle/pkg/infra/cache"
	"iGoStyle/pkg/infra/mysql"
)

func NewLessonRepo(cache *cache.LessonRepo, mysql *mysql.LessonRepo) *LessonRepo {
	return &LessonRepo{cache: cache, mysql: mysql}
}

type LessonRepo struct {
	cache         *cache.LessonRepo
	mysql         *mysql.LessonRepo
	raceProtector singleflight.Group
}

func (repo *LessonRepo) QueryByTutorID(id domain.TutorID) (*domain.Lesson, error) {
	dbQuery := func() (result interface{}, err error) {
		return repo.mysql.QueryByTutorID(id)
	}

	lesson, keyRemainTime, cacheQueryErr := repo.cache.QueryByTutorID(id)
	if cacheQueryErr == nil {
		go policyWithCache(
			keyRemainTime,
			dbQuery,
			cache.KeyService{}.TutorID(id),
			repo.cache.SaveByTutorID,
		)
		return lesson, nil
	}

	policyFn := newPolicyWithoutCache(
		cacheQueryErr,
		dbQuery,
		cache.KeyService{}.TutorID(id),
		repo.cache.SaveByTutorID,
	)

	result, err, _ := repo.raceProtector.Do("QueryByTutorID", policyFn)
	return result.(*domain.Lesson), failure.Wrap(err)
}

func (repo *LessonRepo) QueryAllByTutorIDGroup(ids []domain.TutorID) ([]*domain.Lesson, error) {
	dbQuery := func() (result interface{}, err error) {
		return repo.mysql.QueryAllByTutorIDGroup(ids)
	}

	lessons, keyRemainTime, cacheQueryErr := repo.cache.QueryAllByTutorIDGroup(ids)
	if cacheQueryErr == nil {
		go policyWithCache(
			keyRemainTime,
			dbQuery,
			cache.KeyService{}.TutorIDGroup(ids),
			repo.cache.SaveAllByTutorIDGroup,
		)
		return lessons, nil
	}

	policyFn := newPolicyWithoutCache(
		cacheQueryErr,
		dbQuery,
		cache.KeyService{}.TutorIDGroup(ids),
		repo.cache.SaveAllByTutorIDGroup,
	)

	result, err, shared := repo.raceProtector.Do("QueryAllByTutorIDGroup", policyFn)
	log.Debug().Bool("shared", shared).Send()
	return result.([]*domain.Lesson), failure.Wrap(err)
}
