package logger

import (
	"bytes"
	"fmt"
	//"log/syslog"
	"strings"
	"testing"
)

type Writer struct {
	b *bytes.Buffer
}

func (w *Writer) Alert(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[alert] test[1234]: %s", m))
	return nil
}
func (w *Writer) Close() error {
	_, _ = w.b.WriteString(fmt.Sprintf("[close] test[1234]:"))
	return nil
}
func (w *Writer) Crit(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[crit] test[1234]: %s", m))
	return nil
}
func (w *Writer) Debug(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[debug] test[1234]: %s", m))
	return nil
}
func (w *Writer) Emerg(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[emerg] test[1234]: %s", m))
	return nil
}
func (w *Writer) Err(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[err] test[1234]: %s", m))
	return nil
}
func (w *Writer) Info(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[info] test[1234]: %s", m))
	return nil
}
func (w *Writer) Notice(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[notice] test[1234]: %s", m))
	return nil
}
func (w *Writer) Warning(m string) (err error) {
	_, _ = w.b.WriteString(fmt.Sprintf("[warning] test[1234]: %s", m))
	return nil
}

// Here just so that syslog struct satisfies our Writer interface
func (w *Writer) Write(b []byte) (int, error) {
	return 0, nil
}

func BenchmarkLoggerNew(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		var buff bytes.Buffer
		log := getLogger(&buff)
		if log == nil {
			panic("Logger not created")
		}
	}
}

func getLogger(b *bytes.Buffer) *Logger {

	config := LoggerConfig{
		Application: "test",
		Output:      &Writer{b},
	}

	return New(config)
}

func getBufferString(b *bytes.Buffer) (buffer string) {
	return b.String()
}

func assertLoggedMessage(b *bytes.Buffer, msg string, severity string, t *testing.T) {
	logged := getBufferString(b)

	if !strings.Contains(logged, msg) {
		t.Errorf("Log Message does not contain the logged message. Logged Message: %s\n", logged)
	}

	if !strings.Contains(logged, severity) {
		t.Errorf("Log Message does not contain the severity. Expected: %s\n Actual: %s\n", severity, logged)
	}
}

func TestLoggerNew(t *testing.T) {
	var b bytes.Buffer
	log := getLogger(&b)

	if log.Application != "test" {
		t.Errorf("Unexpected module: %s", log.Application)
	}
}

func TestCritical(t *testing.T) {
	var b bytes.Buffer
	log := getLogger(&b)
	msg := "Something went wrong"
	log.Critical(msg)
	severity := "[crit]"
	assertLoggedMessage(&b, msg, severity, t)
}

func TestError(t *testing.T) {
	var b bytes.Buffer
	log := getLogger(&b)
	msg := "Something went wrong"
	log.Error(msg)
	severity := "[err]"
	assertLoggedMessage(&b, msg, severity, t)
}

func TestWarning(t *testing.T) {
	var b bytes.Buffer
	log := getLogger(&b)
	msg := "Something went wrong"
	log.Warning(msg)
	severity := "[warning]"
	assertLoggedMessage(&b, msg, severity, t)
}

func TestNotice(t *testing.T) {
	var b bytes.Buffer
	log := getLogger(&b)
	msg := "Something went wrong"
	log.Notice(msg)
	severity := "[notice]"
	assertLoggedMessage(&b, msg, severity, t)
}

func TestInfo(t *testing.T) {
	var b bytes.Buffer
	log := getLogger(&b)
	msg := "Something went wrong"
	log.Info(msg)
	severity := "[info]"
	assertLoggedMessage(&b, msg, severity, t)
}

func TestDebug(t *testing.T) {
	var b bytes.Buffer
	log := getLogger(&b)
	msg := "Something went wrong"
	log.Debug(msg)
	severity := "[debug]"
	assertLoggedMessage(&b, msg, severity, t)
}
