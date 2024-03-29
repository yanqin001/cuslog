package cuslog

import (
	"testing"
)

func TestCuslog(t *testing.T) {
	l := New(WithLevel(InfoLevel),
		//cuslog.WithOutput(fd),
		WithFormatter(&JsonFormatter{IgnoreBasicFields: false}),
		WithApp(1, "vuln-center"),
	)
	l.Info("custom log with json formatter")
	l.InfoF("custom log with json formatter")
}
