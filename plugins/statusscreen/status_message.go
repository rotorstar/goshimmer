package statusscreen

import (
	"github.com/iotaledger/hive.go/logger"
	"time"
)

type StatusMessage struct {
	Source   string
	LogLevel logger.LogLevel
	Message  string
	Time     time.Time
}
