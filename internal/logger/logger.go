package itaxlogger

import (
	"github.com/thanakrit-promsiri/gotipitakaduplicate/internal/viperconf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//https://developerpod.com/2019/05/27/adding-uber-go-zap-logger-to-golang-project/

// var sugarLogger *zap.SugaredLogger
var logger *zap.Logger
var logconfig viperconf.LoggerConfiguration

// InitLogger for initail Logger
func InitLogger(logpath string, logconf viperconf.LoggerConfiguration) *zap.Logger {

	logconfig = logconf

	writerSyncer := getLogWriter(logpath)

	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())

	//zapLogger, err := config.Build()

	//sugarLogger = logger.Sugar()
	//sugarLogger, _ = zap.NewProduction()

	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logpath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{

		Filename: logpath,
		MaxSize:  logconfig.Maxsize,
		//MaxBackups: logconfig.Maxbackups,
		//MaxAge:   logconfig.Maxage,
		Compress:  logconfig.Compress,
		LocalTime: true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
