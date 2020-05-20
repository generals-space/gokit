package lorem_metrics

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type ServiceMiddleware func(Service) Service

// MetricsMiddleware 装饰器中间件, 用于Service对象的装饰器
// 实际上相当于metricsMiddleware的构造函数(即常规的New方法)
func MetricsMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) ServiceMiddleware {
	return func(next Service) Service {
		return metricsMiddleware{
			next,
			requestCount,
			requestLatency,
		}
	}
}

type metricsMiddleware struct {
	Service
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// Implement service functions and add label method for our metrics
func (mw metricsMiddleware) Lorem(requestType string, min, max int) (output string, err error) {
	defer func(begin time.Time) {
		vals := []string{"method", "Lorem"}
		mw.requestCount.With(vals...).Add(1)
		mw.requestLatency.With(vals...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.Service.Lorem(requestType, min, max)
	return
}
