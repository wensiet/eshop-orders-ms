package appError

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"log/slog"
)

var (
	UnableToProceedOrder = errors.New("unable to proceed order, there is not enough products")
	InvalidStatus        = errors.New("invalid order's status")
	OrderNotFound        = errors.New("order not found")
)

var errorsMap map[error]bool

func init() {
	errorsMap = map[error]bool{
		UnableToProceedOrder: true,
		InvalidStatus:        true,
		OrderNotFound:        true,
	}
}

func LogIfNotApp(err error, logger *slog.Logger) {
	if !errorsMap[err] {
		sentry.CaptureException(err)
		logger.Error(err.Error())
	}
}
