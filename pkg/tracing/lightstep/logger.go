package lightstep

import (
	"github.com/sirupsen/logrus"
	"github.com/traefik/traefik/v2/pkg/log"
)

// lightstepLogger is an implementation of the Logger interface that delegates to traefik log.
type lightstepLogger struct {
	logger logrus.FieldLogger
}

func newLightstepLogger() *lightstepLogger {
	return &lightstepLogger{
		logger: log.WithoutContext().WithField(log.TracingProviderName, "lightstep"),
	}
}

func (l *lightstepLogger) Error(msg string) {
	l.logger.Errorf("Tracing lightstep error: %s", msg)
}

// Infof logs a message at debug priority.
func (l *lightstepLogger) Infof(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}
