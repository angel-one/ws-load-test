package log

import (
	"context"
	"github.com/angel-one/go-utils/constants"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type Level string

func (l Level) zeroLogLevel() zerolog.Level {
	switch l {
	case constants.TraceLevel:
		return zerolog.TraceLevel
	case constants.DebugLevel:
		return zerolog.DebugLevel
	case constants.InfoLevel:
		return zerolog.InfoLevel
	case constants.WarnLevel:
		return zerolog.WarnLevel
	case constants.ErrorLevel:
		return zerolog.ErrorLevel
	case constants.FatalLevel:
		return zerolog.FatalLevel
	case constants.PanicLevel:
		return zerolog.PanicLevel
	default:
		return zerolog.DebugLevel
	}
}

// InitLogger is used to initialize logger
func InitLogger(level Level) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(level.zeroLogLevel())
	log.Logger = log.With().Caller().Logger()
}

func Trace(ctx context.Context) *zerolog.Event {
	return withID(ctx, log.Trace())
}

// Debug is the for debug log
func Debug(ctx context.Context) *zerolog.Event {
	return withID(ctx, log.Debug())
}

// Info is the for info log
func Info(ctx context.Context) *zerolog.Event {
	return withID(ctx, log.Info())
}

// Warn is the for warn log
func Warn(ctx context.Context) *zerolog.Event {
	return withID(ctx, log.Warn())
}

// Error is the for error log
func Error(ctx context.Context) *zerolog.Event {
	return withID(ctx, log.Error())
}

// Panic is the for panic log
func Panic(ctx context.Context) *zerolog.Event {
	return withID(ctx, log.Panic())
}

// Fatal is the for fatal log
func Fatal(ctx context.Context) *zerolog.Event {
	return withID(ctx, log.Fatal())
}

func withID(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		return event
	}
	return event.Interface(constants.IDLogParam, ctx.Value(constants.IDLogParam))
}
