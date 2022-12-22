package logger

import (
	"os"

	"github.com/satioO/order-mgmt/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	InitLogger()
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
}

type appLogger struct {
	level       string
	devMode     bool
	encoding    string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

// NewAppLogger App Logger constructor
func NewAppLogger(cfg *config.Config) *appLogger {
	return &appLogger{level: cfg.Logger.LogLevel, devMode: cfg.Logger.DevMode, encoding: cfg.Logger.Encoder}
}

// For mapping config logger to email_service logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *appLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *appLogger) InitLogger() {
	logLevel := l.getLoggerLevel()

	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if l.devMode {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.NameKey = "[SERVICE]"
	encoderCfg.TimeKey = "[TIME]"
	encoderCfg.LevelKey = "[LEVEL]"
	encoderCfg.CallerKey = "[LINE]"
	encoderCfg.MessageKey = "[MESSAGE]"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

	if l.encoding == "console" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = logger
	l.sugarLogger = logger.Sugar()
}

// WithName add logger microservice name
func (l *appLogger) WithName(name string) {
	l.logger = l.logger.Named(name)
	l.sugarLogger = l.sugarLogger.Named(name)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *appLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *appLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// Info uses fmt.Sprint to construct and log a message
func (l *appLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *appLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *appLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *appLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}
