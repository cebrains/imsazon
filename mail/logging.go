package mail

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging service
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Send(email *Email) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Send",
			"to", email.To,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.Send(email)
}