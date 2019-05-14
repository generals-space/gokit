package lorem_logging

import (
	"time"

	"github.com/go-kit/kit/log"
)

type ServiceMiddleware func(Service) Service

// LoggingMiddleware 日志中间件, 用于Service对象的装饰器.
// 实际上相当于loggingMiddleware的构造函数(即常规的New方法).
func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return loggingMiddleware{next, logger}
	}
}

type loggingMiddleware struct {
	Service
	logger log.Logger
}

// Implement Service Interface for LoggingMiddleware
func (mw loggingMiddleware) Lorem(requestType string, min, max int) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"type", requestType,
			"min", min,
			"max", max,
			"result", output,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Service.Lorem(requestType, min, max)
	return
}
