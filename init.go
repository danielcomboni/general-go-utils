/**
 * Author: Daniel Comboni
 */

package general_goutils

import (
	"fmt"
	"os"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Initialize(shouldLogToFile ...bool) {
	fmt.Println(runtime.GOOS)

	if len(shouldLogToFile) > 0 && shouldLogToFile[0] {

		// initialize the rotator
		logFile := "logs/app-%Y-%m-%d-%H.log"
		rotator, err := rotatelogs.New(
			logFile,
			rotatelogs.WithMaxAge(60*24*time.Hour),
			rotatelogs.WithRotationTime(time.Hour))
		if err != nil {
			panic(err)
		}

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		fileEncoder := zapcore.NewJSONEncoder(config)
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		writer := zapcore.AddSync(rotator)
		defaultLogLevel := zapcore.DebugLevel
		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		Logger.Info("Now logging in a rotated file")

	} else {

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(config)

		defaultLogLevel := zapcore.DebugLevel

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)

		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		println()
		println("----------------------------")
		Logger.Info("Logging to console only")
		println()
	}

}
