package cuslog

import (
	"testing"
)

func TestCuslog(t *testing.T) {
	l := New(WithLevel(InfoLevel),
		//cuslog.WithOutput(fd),
		WithFormatter(&JsonFormatter{IgnoreBasicFields: false}),
		WithApp("vuln", "vuln-center"),
	)
	l.Info("custom log with json formatter")
	l.InfoF(1, "custom log with json formatter")
}
