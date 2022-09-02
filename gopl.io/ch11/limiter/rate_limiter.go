package limiter

import (
	"context"
	"sort"

	"golang.org/x/time/rate"
)

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *CustomLimiter {
	byLimitComparator := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}

	sort.Slice(limiters, byLimitComparator)
	return &CustomLimiter{limiters: limiters}
}

type CustomLimiter struct {
	limiters []RateLimiter
}

func (ml *CustomLimiter) Wait(ctx context.Context) error {
	for _, limiter := range ml.limiters {
		if err := limiter.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ml *CustomLimiter) Limit() rate.Limit {
	return ml.limiters[0].Limit()
}
