package repository

import (
	"github.com/morikuni/failure"
	"golang.org/x/sync/singleflight"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/infra/cache"
	"AmazingTalker/pkg/infra/mysql"
)

func NewTutorRepo(cache *cache.TutorRepo, mysql *mysql.TutorRepo) *TutorRepo {
	return &TutorRepo{cache: cache, mysql: mysql}
}

type TutorRepo struct {
	cache         *cache.TutorRepo
	mysql         *mysql.TutorRepo
	raceProtector singleflight.Group
}

func (repo *TutorRepo) QueryByTutorSlug(tutorSlug string) (*domain.Tutor, error) {
	dbQuery := func() (result interface{}, err error) {
		return repo.mysql.QueryByTutorSlug(tutorSlug)
	}

	tutor, keyRemainTime, cacheQueryErr := repo.cache.QueryByTutorSlug(tutorSlug)
	if cacheQueryErr == nil {
		go policyWithCache(
			keyRemainTime,
			dbQuery,
			cache.KeyService{}.TutorSlug(tutorSlug),
			repo.cache.SaveByTutorSlug,
		)
		return tutor, nil
	}

	policyFn := newPolicyWithoutCache(
		cacheQueryErr,
		dbQuery,
		cache.KeyService{}.TutorSlug(tutorSlug),
		repo.cache.SaveByTutorSlug,
	)

	result, err, _ := repo.raceProtector.Do("QueryByTutorSlug", policyFn)
	return result.(*domain.Tutor), failure.Wrap(err)
}

func (repo *TutorRepo) QueryAllByLanguageID(id domain.LanguageID) (domain.TutorGroup, error) {
	dbQuery := func() (result interface{}, err error) {
		return repo.mysql.QueryAllByLanguageID(id)
	}

	tutors, keyRemainTime, cacheQueryErr := repo.cache.QueryAllByLanguageID(id)
	if cacheQueryErr == nil {
		go policyWithCache(
			keyRemainTime,
			dbQuery,
			cache.KeyService{}.LanguageID(id),
			repo.cache.SaveAllByLanguageID,
		)
		return tutors, nil
	}

	policyFn := newPolicyWithoutCache(
		cacheQueryErr,
		dbQuery,
		cache.KeyService{}.LanguageID(id),
		repo.cache.SaveAllByLanguageID,
	)

	result, err, _ := repo.raceProtector.Do("QueryAllByLanguageID", policyFn)
	return result.(domain.TutorGroup), failure.Wrap(err)
}
