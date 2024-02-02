package util

import (
	"context"
	"errors"
	"time"
)

func IsDeadlineExceeded(err error) bool {
	return errors.Is(err, context.DeadlineExceeded)
}

func UnixSecondsToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
