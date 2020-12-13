package repository

import (
	"time"

	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"

	"AmazingTalker/pkg/infra/cache"
	"AmazingTalker/pkg/technical/errutil"
)

func newPolicyWithoutCache(
	cacheQueryErr error,
	dbQuery func() (result interface{}, err error),
	cacheKey string,
	cacheSave cache.SaverFunc,
) func() (interface{}, error) {

	return func() (interface{}, error) {
		return policyWithoutCache(cacheQueryErr, dbQuery, cacheKey, cacheSave)
	}
}

func policyWithoutCache(
	cacheQueryErr error,
	dbQuery func() (result interface{}, err error),
	cacheKey string,
	cacheSave cache.SaverFunc,
) (interface{}, error) {

	if !failure.Is(cacheQueryErr, errutil.ErrNotFound) {
		Err := failure.Wrap(cacheQueryErr)
		log.Error().Msgf("%v", Err)
	}

	result, dbQueryErr := dbQuery()
	if dbQueryErr != nil {
		return nil, failure.Wrap(dbQueryErr)
	}

	if saveErr := cacheSave(cacheKey, result); saveErr != nil {
		Err := failure.Wrap(saveErr)
		log.Error().Msgf("%v", Err)
		if log.Debug().Enabled() {
			log.Error().Msgf("policyWithoutCache: ErrorStack=\n%+v", Err)
		}
	}

	return result, nil
}

func policyWithCache(
	keyRemainTime time.Duration,
	dbQuery func() (result interface{}, err error),
	cacheKey string,
	cacheSave cache.SaverFunc,
) {

	if keyRemainTime <= 10*time.Minute {
		return
	}

	result, dbQueryErr := dbQuery()
	if dbQueryErr != nil {
		Err := failure.Wrap(dbQueryErr)
		log.Error().Msgf("%v", Err)
		if log.Debug().Enabled() {
			log.Error().Msgf("policyWithCache: ErrorStack=\n%+v", Err)
		}
		return
	}

	if saveErr := cacheSave(cacheKey, result); saveErr != nil {
		Err := failure.Wrap(saveErr)
		log.Error().Msgf("%v", Err)
		if log.Debug().Enabled() {
			log.Error().Msgf("policyWithCache: ErrorStack=\n%+v", Err)
		}
		return
	}
}
