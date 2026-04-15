package logger

import (
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOption struct {
	filePath  string
	maxSize   int
	maxAge    int
	localTime bool
	compress  bool
	console bool
	formatter string
	aLog *zap.Logger
}

type LoggerEngineOption func(*LoggerOption)


func NewLogger(option ...LoggerEngineOption) *LoggerOption {
	l := &LoggerOption{
		filePath: "app.log", // 默认日志文件
		maxSize: 1024, // 默认大小 1 G
		maxAge: 30, // 默认保留 30 天
		localTime: false,
		compress: true,
		console: false,
		formatter: "json", // json(JSON) | console(CONSOLE)
	}
	for _, o := range option {
		o(l)
	}

	l.start()
	return l
}

func  WithLoggerFilePath(path string) LoggerEngineOption {
	return func(o *LoggerOption) {
		o.filePath = path
	}
}

func  WithLoggerMaxSize(size int) LoggerEngineOption {
	return func(o *LoggerOption) {
		o.maxSize = size
	}
}

func WithLoggerMaxAge(age int) LoggerEngineOption {
	return func(o *LoggerOption) {
		o.maxAge = age
	}
}

func WithLoggerLocalTime() LoggerEngineOption {
	return func(o *LoggerOption) {
		o.localTime = true
	}
}

func WithLoggerCompress() LoggerEngineOption {
	return func(o *LoggerOption) {
		o.compress = true
	}
}

func WithLoggerConsole(consoleOutput bool) LoggerEngineOption {
	return func (o *LoggerOption)  {
		o.console = consoleOutput
	}
}

func WithLoggerFormatter(formatter string) LoggerEngineOption {
	return func(o *LoggerOption) {
		o.formatter = formatter
	}
}

func (l *LoggerOption) Close() error {
	if l.aLog != nil {
		return l.aLog.Sync()
	}
	return nil
}
   
func (l *LoggerOption) start() {
	if l.aLog != nil {
		return
	}

	// file output
	writeSyncer := zapWriteSyncer(l.filePath, l.maxSize, l.maxAge, l.localTime, l.compress)

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	var fileEncoder zapcore.Encoder
	switch strings.ToLower(l.formatter) {
	case "json":
		// 文件 core（使用 JSON 更适合生产，文件日志推荐 JSON)
		encoder.TimeKey = "time"
		fileEncoder = zapcore.NewJSONEncoder(encoder)
	case "console": 
		// 文件 core（使用更直观的控制台输出的形式写入日志文件)
		fileEncoder = zapcore.NewConsoleEncoder(encoder)
	default: 
		fileEncoder = zapcore.NewJSONEncoder(encoder)
	}
	
	fileCore := zapcore.NewCore(fileEncoder, writeSyncer, zapcore.DebugLevel)

	if l.console {
		// 控制台输出（开发友好，使用 Console Encoder）
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleDebug := zapcore.Lock(os.Stdout)
		consoleCore := zapcore.NewCore(consoleEncoder, consoleDebug, zapcore.DebugLevel)

		core := zapcore.NewTee(fileCore, consoleCore)
			l.aLog = zap.New(core, 
			zap.AddCaller(), // 显示调用文件和行号
			zap.AddCallerSkip(1), // 跳过包装层（可选，根据需要调整）
		)
	} else {
		core := zapcore.NewTee(fileCore)
			l.aLog = zap.New(core, 
			zap.AddCaller(), 
			zap.AddCallerSkip(1),
		)
	}
}

func zapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer(fileName string, maxSize int, maxAge int, localTime, compress bool) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   maxSize,
		MaxAge:    maxAge,
		LocalTime: localTime,
		Compress:  compress,
	})
}

func (l *LoggerOption) Debug(msg string, filed ...zap.Field) {
	l.aLog.Debug(msg, filed...)
}

func (l *LoggerOption) Info(msg string, fileds ...zap.Field) {
	l.aLog.Info(msg, fileds...)
}

func (l *LoggerOption) Warn(msg string, filed ...zap.Field) {
	l.aLog.Warn(msg, filed...)
}

func (l *LoggerOption) Error(msg string, filed ...zap.Field) {
	l.aLog.Error(msg, filed...)
}

func (l *LoggerOption) DPanic(msg string, filed ...zap.Field) {
	l.aLog.DPanic(msg, filed...)
}

func (l *LoggerOption) Panic(msg string, filed ...zap.Field) {
	l.aLog.Panic(msg, filed...)
}

func (l *LoggerOption) Fatal(msg string, filed ...zap.Field) {
	l.aLog.Fatal(msg, filed...)
}


func (l *LoggerOption) SugarDebug(args ...interface{}) {
	l.aLog.Sugar().Debug(args...)
}

func (l *LoggerOption) SugarDebugf(template string, args ...interface{}) {
	l.aLog.Sugar().Debugf(template, args...)
}

func (l *LoggerOption) SugarDebugw(template string, keysAndValues ...interface{}) {
	l.aLog.Sugar().Debugw(template, keysAndValues...)
}

func (l *LoggerOption) SugarInfo(args ...interface{}) {
	l.aLog.Sugar().Info(args...)
}

func (l *LoggerOption) SugarInfof(template string, args ...interface{}) {
	l.aLog.Sugar().Infof(template, args...)
}

func (l *LoggerOption) SugarInfow(template string, keysAndValues ...interface{}) {
	l.aLog.Sugar().Infow(template, keysAndValues...)
}

func (l *LoggerOption) SugarWarn(args ...interface{}) {
	l.aLog.Sugar().Warn(args...)
}

func (l *LoggerOption) SugarWarnf(template string, args ...interface{}) {
	l.aLog.Sugar().Warnf(template, args...)
}

func (l *LoggerOption) SugarWarnw(template string, keysAndValues ...interface{}) {
	l.aLog.Sugar().Warnw(template, keysAndValues...)
}

func (l *LoggerOption) SugarError(args ...interface{}) {
	l.aLog.Sugar().Error(args...)
}

func (l *LoggerOption) SugarErrorf(template string, args ...interface{}) {
	l.aLog.Sugar().Errorf(template, args...)
}

func (l *LoggerOption) SugarErrorw(template string, keysAndValues ...interface{}) {
	l.aLog.Sugar().Errorw(template, keysAndValues...)
}

func (l *LoggerOption) SugarDPanic(args ...interface{}) {
	l.aLog.Sugar().DPanic(args...)
}

func (l *LoggerOption) SugarDPanicf(template string, args ...interface{}) {
	l.aLog.Sugar().DPanicf(template, args...)
}

func (l *LoggerOption) SugarDPanicw(template string, keysAndValues ...interface{}) {
	l.aLog.Sugar().DPanicw(template, keysAndValues...)
}

func (l *LoggerOption) SugarPanic(args ...interface{}) {
	l.aLog.Sugar().Panic(args...)
}

func (l *LoggerOption) SugarPanicf(template string, args ...interface{}) {
	l.aLog.Sugar().Panicf(template, args...)
}

func (l *LoggerOption) SugarPanicw(template string, keysAndValues ...interface{}) {
	l.aLog.Sugar().Panicw(template, keysAndValues...)
}


func (l *LoggerOption) SugarFatal(args ...interface{}) {
	l.aLog.Sugar().Fatal(args...)
}

func (l *LoggerOption) SugarFatalf(template string, args ...interface{}) {
	l.aLog.Sugar().Fatalf(template, args...)
}

func (l *LoggerOption) SugarFatalw(template string, keysAndValues ...interface{}) {
	l.aLog.Sugar().Fatalw(template, keysAndValues...)
}
