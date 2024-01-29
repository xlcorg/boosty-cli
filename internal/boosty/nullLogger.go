package boosty

type Logger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

var _ Logger = (*nullLogger)(nil)

type nullLogger struct {
}

func (l *nullLogger) Errorf(format string, v ...interface{}) {
}

func (l *nullLogger) Warnf(format string, v ...interface{}) {
}

func (l *nullLogger) Debugf(format string, v ...interface{}) {
}
