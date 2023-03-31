package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLog *zap.SugaredLogger

func InitLogger() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	SugarLog = logger.Sugar()

	SugarLog.Info("Logger is initialized")
}
