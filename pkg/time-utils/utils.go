package timeutils

import "time"

func TimeNow() *time.Time {
	now := time.Now().Truncate(time.Microsecond)
	return &now
}
