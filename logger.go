package itools

import (
	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type LoggerOption struct {
	FilePath  string
	MaxSize   int
	MaxAge    int
	LocalTime bool
	Compress  bool
}

type LoggerEngineOption func(*LoggerOption)

func NewLogger(option ...LoggerEngineOption) *LoggerOption {
	l := &LoggerOption{}
	for _, o := range option {
		o(l)
	}
	return l
}

func WithLoggerFilePath(path string) LoggerEngineOption {
	return func(o *LoggerOption) {
		o.FilePath = path
	}
}

func WithLoggerMaxSize(size int) LoggerEngineOption {
	return func(o *LoggerOption) {
		o.MaxSize = size
	}
}

func WithLoggerMaxAge(age int) LoggerEngineOption {
	return func(o *LoggerOption) {
		o.MaxAge = age
	}
}

func WithLoggerLocalTime() LoggerEngineOption {
	return func(o *LoggerOption) {
		o.LocalTime = true
	}
}

func WithLoggerCompress() LoggerEngineOption {
	return func(o *LoggerOption) {
		o.Compress = true
	}
}

var (
	AppLog *zap.Logger
)

func (l *LoggerOption) Start() {
	encoder := zapEncoder()
	writeSyncer := zapWriteSyncer(l.FilePath, l.MaxSize, l.MaxAge, l.LocalTime, l.Compress)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	consoleDebug := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	var Codes []zapcore.Core
	Codes = append(Codes, core)
	Codes = append(Codes, zapcore.NewCore(consoleEncoder, consoleDebug, p))
	c := zapcore.NewTee(Codes...)
	AppLog = zap.New(c, zap.AddCaller())
}

func zapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer(fileName string, maxSize int, maxAge int, localTime, compress bool) zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   maxSize,
		MaxAge:    maxAge,
		LocalTime: localTime,
		Compress:  compress,
	}
	c := cron.New()
	_ = c.AddFunc("0 0 0 1/1 * ?", func() {
		_ = logger.Rotate()
	})

	c.Start()
	return zapcore.AddSync(logger)
}

var xLogger *ormLogger

type ormLogger struct{}

func (x *ormLogger) Write(p []byte) (n int, err error) {
	AppLog.Info("数据库操作", zap.String("数据库", string(p)))
	return len(p), nil
}

var EchoLog *EchoLogger

type EchoLogger struct {
}

func (el *EchoLogger) Write(p []byte) (n int, err error) {
	AppLog.Info("ECHO", zap.String("请求", string(p)))
	return len(p), nil
}
