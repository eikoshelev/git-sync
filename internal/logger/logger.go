package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger() *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapConfig.DisableStacktrace = true
	zapConfig.DisableCaller = true
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("error setup zap logger: %s", err.Error())
	}
	return logger
}
