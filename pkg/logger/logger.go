package logger

import (
	"errors"
	"github.com/miladshalikar/cafe/pkg/richerror"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

const (
	defaultFilePath        = "logs/logs.json"
	defaultUseLocalTime    = false
	defaultFileMaxSizeInMB = 10
	defaultFileAgeInDays   = 30
)

type Config struct {
	FilePath         string
	UseLocalTime     bool
	FileMaxSizeInMB  int
	FileMaxAgeInDays int
}

var l *slog.Logger

func init() {
	fileWriter := &lumberjack.Logger{
		Filename:  defaultFilePath,
		LocalTime: defaultUseLocalTime,
		MaxSize:   defaultFileMaxSizeInMB,
		MaxAge:    defaultFileAgeInDays,
	}
	l = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{}),
	)
}

func L() *slog.Logger {
	return l
}

func New(cfg Config, opt *slog.HandlerOptions) *slog.Logger {
	fileWriter := &lumberjack.Logger{
		Filename:  cfg.FilePath,
		LocalTime: cfg.UseLocalTime,
		MaxSize:   cfg.FileMaxSizeInMB,
		MaxAge:    cfg.FileMaxAgeInDays,
	}

	logger := slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), opt),
	)

	return logger
}

func Log(err error) {
	var rErr richerror.RichError

	if errors.As(err, &rErr) {
		logger := L().With(
			slog.String("operation", rErr.Op()),
		)

		if rErr.WErr() != nil {
			logger = logger.With(slog.String("error", rErr.WErr().Error()))
		}

		if meta := rErr.Meta(); meta != nil {
			for k, v := range meta {
				logger = logger.With(slog.Any("meta."+k, v))
			}
		}

		logger.Error(rErr.Message())
		return
	}

	L().Error(err.Error())
}
