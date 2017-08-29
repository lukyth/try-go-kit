package middleware

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/lukyth/try-go-kit/svc-string/pkg/service"
)

// InstrumentingMiddleware returns a service middleware that count the number of jobs
// processed, record the duration of requests after they’ve finished and track the number
// of in-flight operations.
func InstrumentingMiddleware(requestCount metrics.Counter, requestLatency, countResult metrics.Histogram) Middleware {
	return func(next service.StringService) service.StringService {
		return instrumentingMiddleware{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			countResult:    countResult,
			next:           next,
		}
	}
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.StringService
}

func (mw instrumentingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw instrumentingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.countResult.Observe(float64(n))
	}(time.Now())

	n = mw.next.Count(s)
	return
}
