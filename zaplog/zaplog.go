package zaplog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var InitLogger *Logger = nil

type Logger struct {
	logger *zap.Logger
	sugarLogger *zap.SugaredLogger
}

func GetInitLogger() *Logger {
	return InitLogger
}

func (log *Logger) Info(msg string, fields ...zapcore.Field) {
	InitLogger.logger.Info(msg, fields...)
}

func (*Logger) Infof(template string, args ...interface{}) {
	InitLogger.sugarLogger.Infof(template, args...)
}

func (s *Logger) Infow(msg string, keysAndValues ...interface{}) {
	InitLogger.sugarLogger.Infow(msg, keysAndValues...)
}

func (*Logger) Debugf(template string, args ...interface{}) {
	InitLogger.sugarLogger.Debugf(template, args...)
}

func (s *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	InitLogger.sugarLogger.Debugw(msg, keysAndValues...)
}

func (*Logger) Warnf(template string, args ...interface{}) {
	InitLogger.sugarLogger.Warnf(template, args...)
}

func (s *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	InitLogger.sugarLogger.Warnw(msg, keysAndValues...)
}

func (*Logger) Errorf(template string, args ...interface{}) {
	InitLogger.sugarLogger.Errorf(template, args...)
}

func (s *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	InitLogger.sugarLogger.Errorw(msg, keysAndValues...)
}

func (*Logger) Fatalf(template string, args ...interface{}) {
	InitLogger.sugarLogger.Fatalf(template, args...)
}

func (s *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	InitLogger.sugarLogger.Fatalw(msg, keysAndValues...)
}

func Init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		// CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		// EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeLevel:    zapcore.CapitalLevelEncoder,  // 大写编码器
		// EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 大写加颜色编辑器
		// EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      true,                                                // 开发模式，堆栈跟踪
		Encoding:         "console",                                              // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                       // 编码器配置
		// InitialFields:    map[string]interface{}{"serviceName": "test"},       // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout", "./test.log"},                    // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to init log: %v", err))
	}
	defer logger.Sync()
	// logger.Info("Succeed to init log")

	//logger.Info("Unable to get web url",
	//	zap.String("url", "http://www.baidu.com"),
	//	zap.Int("attempt", 3),
	//	zap.Duration("backoff", time.Second),
	//)

	sugarLogger := logger.Sugar()
	defer sugarLogger.Sync()

	//sugarLogger.Infow("Failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", "http://www.baidu.com",
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugarLogger.Infof("Failed to fetch URL: %s", "http://www.baidu.com")

	InitLogger = &Logger {
		logger: logger,
		sugarLogger: sugarLogger,
	}
}

func Test() {
	// logger0, _ := zap.NewDevelopment()
	logger0, _ := zap.NewProduction()
	defer logger0.Sync()

	logger0.Info("Unable to get web url",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "mesg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		// EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      true,                                                // 开发模式，堆栈跟踪
		Encoding:         "console",                                           // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                       // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "test"},       // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout", "./test.log"},                    // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to init log: %v", err))
	}
	defer logger.Sync()
	logger.Info("Succeed to init log", )

	logger.Info("Unable to get web url",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugarLogger := logger.Sugar()
	defer sugarLogger.Sync()

	sugarLogger.Infow("Failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "http://www.baidu.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugarLogger.Infof("Failed to fetch URL: %s", "http://www.baidu.com")
}
