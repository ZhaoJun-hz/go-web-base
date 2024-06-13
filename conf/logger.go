package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	separator := string(filepath.Separator)
	rootDir, _ := os.Getwd()
	logFilePath := rootDir + separator + "log" + separator + time.Now().Format(time.DateOnly) + ".log"
	l := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    viper.GetInt("log.maxSize"), // megabytes
		MaxBackups: viper.GetInt("log.maxBackups"),
		MaxAge:     viper.GetInt("log.maxAge"), //days
		Compress:   false,                      // disabled by default
	}

	return zapcore.AddSync(l)
}
