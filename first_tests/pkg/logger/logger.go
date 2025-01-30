package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

// Уровни логирования
const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// ANSI-коды для цветов
var levelColors = map[int]string{
	LevelDebug: "\033[36m", // Голубой
	LevelInfo:  "\033[32m", // Зеленый
	LevelWarn:  "\033[33m", // Желтый
	LevelError: "\033[31m", // Красный
	LevelFatal: "\033[35m", // Фиолетовый
}

// Логгер
type Logger struct {
	slog slog.Logger
	level       int
	output      io.Writer
	enableColor bool
	mu          sync.Mutex
	exitFn      func(int)
}

// Конструктор логгера
func NewLogger(level int, output io.Writer) *Logger {
	logger := &Logger{
		level:       level,
		output:      output,
		enableColor: isTerminal(output), // Определяем терминал
		exitFn:      os.Exit,            // Используем os.Exit
	}
	
	return logger
}

func (l *Logger) Error(msg string, args ...any)  {
	// handle error
	l.slog.Error(fmt.Sprintf("%s  %s","my error ",msg),args)
}
// Проверяет, является ли `output` терминалом
func isTerminal(output io.Writer) bool {
	if file, ok := output.(*os.File); ok {
		return term.IsTerminal(int(file.Fd()))
	}
	return false
}

// Логирование с метаданными
func (l *Logger) log(level int, message string, meta map[string]interface{}) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[level]

	// Добавляем цвет
	colorStart, colorEnd := "", ""
	if l.enableColor {
		colorStart, colorEnd = levelColors[level], "\033[0m"
	}

	// Формируем строку с метаданными
	metaStr := ""
	if meta != nil {
		for k, v := range meta {
			metaStr += fmt.Sprintf(" %s=%v", k, v)
		}
	}

	fmt.Fprintf(l.output, "%s[%s] [%s] : %s%s%s\n", colorStart, timestamp, levelStr, message, metaStr, colorEnd)

	if level == LevelFatal {
		l.exitFn(1)
	}
}

// Методы логирования
func (l *Logger) Debug(msg string, meta map[string]interface{}) { l.log(LevelDebug, msg, meta) }
func (l *Logger) Info(msg string, meta map[string]interface{})  { l.log(LevelInfo, msg, meta) }
func (l *Logger) Warn(msg string, meta map[string]interface{})  { l.log(LevelWarn, msg, meta) }
// func (l *Logger) Error(msg string, meta map[string]interface{}) { l.log(LevelError, msg, meta) }
func (l *Logger) Fatal(msg string, meta map[string]interface{}) { l.log(LevelFatal, msg, meta) }
