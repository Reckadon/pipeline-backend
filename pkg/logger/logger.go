package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/instill-ai/pipeline-backend/config"
)

var logger *zap.Logger
var once sync.Once
var core zapcore.Core

// GetZapLogger returns an instance of zap logger
func GetZapLogger() (*zap.Logger, error) {
	var err error
	once.Do(func() {
		// info level enabler
		infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level == zapcore.InfoLevel
		})

		// error and fatal level enabler
		errorFatalLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level == zapcore.WarnLevel || level == zapcore.ErrorLevel || level == zapcore.FatalLevel
		})

		// write syncers
		stdoutSyncer := zapcore.Lock(os.Stdout)
		stderrSyncer := zapcore.Lock(os.Stderr)

		// tee core
		if config.Config.Server.Debug {
			core = zapcore.NewTee(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
					stdoutSyncer,
					infoLevel,
				),
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
					stderrSyncer,
					errorFatalLevel,
				),
			)
		} else {
			core = zapcore.NewTee(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
					stdoutSyncer,
					infoLevel,
				),
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
					stderrSyncer,
					errorFatalLevel,
				),
			)
		}

		// finally construct the logger with the tee core
		logger = zap.New(core)
	})

	return logger, err
}