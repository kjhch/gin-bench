package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

const (
	DefaultLogger LoggerName = "default"
	MetricsLogger LoggerName = "metrics"
)

type LoggerName string

type LoggerConf struct {
	//zap.Config `mapstructure:",squash"`
	Encoding string
	Level    string
	Outputs  []OutputConf
}

type OutputConf struct {
	File       string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type LoggerFactory struct {
	loggerMap map[LoggerName]*zap.Logger
}

func NewLoggerFactory() *LoggerFactory {
	loggerFactory := &LoggerFactory{
		loggerMap: make(map[LoggerName]*zap.Logger),
	}
	for k, v := range appConf.Logger {
		loggerFactory.loggerMap[LoggerName(k)] = buildLogger(v)
	}
	loggerFactory.GetDefaultLogger().Info("logger factory init...")
	return loggerFactory
}
func (lf *LoggerFactory) GetLogger(name LoggerName) *zap.Logger {
	return lf.loggerMap[name]
}
func (lf *LoggerFactory) GetDefaultLogger() *zap.Logger {
	return lf.GetLogger(DefaultLogger)
}

func buildLogger(c LoggerConf) *zap.Logger {
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		panic(err)
	}
	return zap.New(zapcore.NewCore(getEncoder(c.Encoding), getWriteSyncer(c.Outputs), level), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

func getEncoder(encoding string) zapcore.Encoder {
	ecfg := zap.NewProductionEncoderConfig()
	ecfg.EncodeLevel = zapcore.CapitalLevelEncoder
	ecfg.EncodeTime = zapcore.ISO8601TimeEncoder
	ecfg.FunctionKey = "function"
	ecfg.NameKey = "name"
	if encoding == "json" {
		return zapcore.NewJSONEncoder(ecfg)
	}
	return zapcore.NewConsoleEncoder(ecfg)
}

func getWriteSyncer(confs []OutputConf) zapcore.WriteSyncer {
	writeSyncers := make([]zapcore.WriteSyncer, 0)
	for _, c := range confs {
		switch strings.ToLower(c.File) {
		case "stdout", "console":
			writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
		default:
			lumberjackLogger := &lumberjack.Logger{
				Filename:   c.File,
				MaxSize:    c.MaxSize,
				MaxAge:     c.MaxAge,
				MaxBackups: c.MaxBackups,
				LocalTime:  true,
				Compress:   true,
			}
			writeSyncers = append(writeSyncers, zapcore.AddSync(lumberjackLogger))
		}
	}

	return zapcore.NewMultiWriteSyncer(writeSyncers...)
}
