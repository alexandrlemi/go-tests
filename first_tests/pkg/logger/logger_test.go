package logger

import (
	"bytes"
	"io"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

// Тест уровней логирования
func TestLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(LevelWarn, &buf)
	logger.Debug("Debug message", nil)
	logger.Info("Info message", nil)
	logger.Warn("Warn message", nil)
	logger.Error("Error message", nil)

	// Проверяем паттерн
	expectedPattern := `\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \[(WARN|ERROR)\] : .*`
	matched, err := regexp.MatchString(expectedPattern, buf.String())
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("Expected log output to match pattern %q, got %q", expectedPattern, buf.String())
	}

	// Проверяем наличие WARN и ERROR
	if !strings.Contains(buf.String(), "[WARN]") || !strings.Contains(buf.String(), "[ERROR]") {
		t.Errorf("Expected log output to contain WARN and ERROR messages, got %q", buf.String())
	}
}

// Тест цветного логирования
func TestLoggerColor(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(LevelInfo, &buf)

	// Проверка вывода без цвета
	logger.Info("Test message", nil)
	if !strings.Contains(buf.String(), "[INFO]") {
		t.Errorf("Expected log output to contain 'INFO', got %q", buf.String())
	}

	// Проверка вывода с цветом
	buf.Reset()
	colorBuf := &mockTerminal{Writer: &buf}
	logger.output = colorBuf
	logger.enableColor = true

	logger.Info("Test message", nil)

	// Проверяем наличие ANSI-кодов
	if !strings.Contains(buf.String(), "\033[") {
		t.Errorf("Expected colored output, got no ANSI codes: %q", buf.String())
	}
}

// Мок терминала для проверки ANSI-кодов
type mockTerminal struct {
	io.Writer
}

func (m *mockTerminal) Fd() uintptr {
	return 1
}

// Тест завершения при FATAL
func TestLoggerFatal(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(LevelFatal, &buf)
	done := make(chan bool)
	logger.exitFn = func(int) { done <- true }

	go func() {
		logger.Fatal("Fatal message", nil)
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("Expected program to exit on Fatal, but it did not")
	}
}

// Тест многопоточности
func TestLoggerConcurrency(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(LevelInfo, &buf)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger.Info("Concurrent message", map[string]interface{}{"goroutine": i})
		}(i)
	}
	wg.Wait()

	// Проверяем количество строк
	lines := bytes.Split(buf.Bytes(), []byte("\n"))
	if len(lines)-1 != 100 {
		t.Errorf("Expected 100 log lines, got %d", len(lines)-1)
	}
}

func TestLoggerError(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(LevelInfo, &buf)
	logger.Error("test")
}
