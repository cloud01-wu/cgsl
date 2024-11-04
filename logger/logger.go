package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"

	util "github.com/cloud01-wu/cgsl/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// default log file folders
const (
	defaultLogFolder = (".")
)

var logLevelDirt = map[string]int{
	"debug":  -1,
	"info":   0,
	"warn":   1,
	"error":  2,
	"dpanic": 3,
	"panic":  4,
	"fatal":  5,
}

// Logger ...
type Logger struct {
	*zap.Logger
}

var (
	logFolder      string
	logRotateCount int
	logLevel       zapcore.Level
	once           sync.Once
	once1          sync.Once
	instance       *Logger
)

func CustomConsoleEncoder() zapcore.Encoder {
	consoleEncoderConfig := zap.NewProductionConfig().EncoderConfig
	consoleEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // ISO8601 UTC
	consoleEncoderConfig.EncodeLevel = customLevelEncoder
	consoleEncoderConfig.TimeKey = "time"
	consoleEncoderConfig.LevelKey = "level"
	consoleEncoderConfig.NameKey = "logger"
	consoleEncoderConfig.CallerKey = ""

	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	return consoleEncoder
}

func CustomJSONEncoder() zapcore.Encoder {
	jsonEncoderConfig := zap.NewProductionConfig().EncoderConfig
	jsonEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // ISO8601 UTC
	jsonEncoderConfig.TimeKey = "time"
	jsonEncoderConfig.LevelKey = "level"
	jsonEncoderConfig.NameKey = "logger"
	jsonEncoderConfig.CallerKey = ""

	jsonEncoder := zapcore.NewJSONEncoder(jsonEncoderConfig)
	return jsonEncoder
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", level.String()))
}

// SetEnvParam ...
func SetEnvParam(folderString string, rotateCount int, levelSetting string) {
	once1.Do(func() {
		logFolder = folderString
		logRotateCount = rotateCount
		val, ok := logLevelDirt[levelSetting]
		if !ok {
			val = -1
		}
		logLevel = zapcore.Level(val)
	})
}

// New ...
// logger.New().Info("This is an Info message", zap.String("customField", "123"))
func New() *Logger {
	once.Do(func() {
		// if SetEnvParam didn't call
		if strings.EqualFold(logFolder, "") {
			logFolder = defaultLogFolder
			logLevel = zapcore.Level(-1)
		}

		// check levels
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= logLevel
		})

		appName := util.GetAppName()
		if _, err := os.Stat(logFolder); os.IsNotExist(err) {
			// does not exist
			os.MkdirAll(logFolder, os.ModePerm)
		}

		fileRotateHook := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFolder + string(os.PathSeparator) + appName + ".log",
			MaxSize:    1,              // MB
			MaxBackups: logRotateCount, // max number of backup
			MaxAge:     30,             // day
			Compress:   false,
		})
		consoleHook := zapcore.Lock(os.Stdout)

		consoleEncoder := CustomConsoleEncoder()

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, fileRotateHook, lowPriority),
			zapcore.NewCore(consoleEncoder, consoleHook, lowPriority))

		instance = &Logger{zap.New(core, zap.AddCaller())}
	})

	return instance
}

func WriteHttpHandlerLog(logLevel zapcore.Level, statusCode int, message string, params map[string]interface{}) {
	callerName := util.CurrentCallerName()
	switch logLevel {
	case zapcore.DebugLevel:
		New().Debug(
			callerName,
			zap.Any("Params", params),
			zap.Int("Status", statusCode),
			zap.String("ErrorMessage", message))
	case zapcore.InfoLevel:
		New().Info(
			callerName,
			zap.Any("Params", params),
			zap.Int("Status", statusCode),
			zap.String("ErrorMessage", message))
	case zapcore.WarnLevel:
		New().Warn(
			callerName,
			zap.Any("Params", params),
			zap.Int("Status", statusCode),
			zap.String("ErrorMessage", message))
	case zapcore.ErrorLevel:
		New().Error(
			callerName,
			zap.Any("Params", params),
			zap.Int("Status", statusCode),
			zap.String("ErrorMessage", message))
	case zapcore.FatalLevel:
		New().Fatal(
			callerName,
			zap.Any("Params", params),
			zap.Int("Status", statusCode),
			zap.String("ErrorMessage", message))
	case zapcore.PanicLevel:
		New().Panic(
			callerName,
			zap.Any("Params", params),
			zap.Int("Status", statusCode),
			zap.String("ErrorMessage", message))
	case zapcore.DPanicLevel:
		New().DPanic(
			callerName,
			zap.Any("Params", params),
			zap.Int("Status", statusCode),
			zap.String("ErrorMessage", message))
	}
}
