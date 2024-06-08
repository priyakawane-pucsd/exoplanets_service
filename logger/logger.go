package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

var (
	zapLogger *zap.Logger
)

func init() {

	zapConfig := zap.NewProductionConfig()

	zapLogger, _ = zapConfig.Build(zap.AddCallerSkip(1))

}

func Info(format string, a ...any) {
	zapLogger.Info(fmt.Sprintf(format, a...))
}

func Error(ctx context.Context, format string, a ...any) {
	zapLogger.Error(fmt.Sprintf(format, a...))
}

func Warn(ctx context.Context, format string, a ...any) {
	zapLogger.Warn(fmt.Sprintf(format, a...))
}

func Debug(ctx context.Context, format string, a ...any) {
	zapLogger.Debug(fmt.Sprintf(format, a...))
}

func Panic(ctx context.Context, format string, a ...any) {
	zapLogger.Panic(fmt.Sprintf(format, a...))
}

func Fatal(ctx context.Context, format string, a ...any) {
	zapLogger.Fatal(fmt.Sprintf(format, a...))
}
