package database

import (
	"context"
	"fmt"
	"time"

	"github.com/gary1030/learning-o11y/server/pkg/log"
	"gorm.io/gorm/logger"
)

type gormLogger struct {
	config logger.Config
}

// LogMode sets the log level for the logger.
func (gl *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newConfig := gl.config
	newConfig.LogLevel = level
	return &gormLogger{config: newConfig}
}

// Info logs informational messages.
func (gl *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if gl.config.LogLevel >= logger.Info {
		log.Info(fmt.Sprintf(msg, data...))
	}
}

// Warn logs warning messages.
func (gl *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if gl.config.LogLevel >= logger.Warn {
		log.Warn(fmt.Sprintf(msg, data...))
	}
}

// Error logs error messages.
func (gl *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if gl.config.LogLevel >= logger.Error {
		log.Fatal(fmt.Sprintf(msg, data...))
	}
}

// Trace logs detailed information about SQL queries.
func (gl *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.config.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		log.Fatal(fmt.Sprintf("SQL error: %s [%s] (%d rows) - took %s", err, sql, rows, elapsed))
	} else if elapsed > gl.config.SlowThreshold {
		log.Warn(fmt.Sprintf("Slow SQL: [%s] (%d rows) - took %s", sql, rows, elapsed))
	} else {
		log.Info(fmt.Sprintf("SQL: [%s] (%d rows) - took %s", sql, rows, elapsed))
	}
}
