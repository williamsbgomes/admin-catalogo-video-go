package timeutils

import (
	"testing"
	"time"
)

func TestTimeNow_TruncatesToMicrosecond(t *testing.T) {
	now := TimeNow()
	if now.Nanosecond()%1000 != 0 {
		t.Errorf(
			"expected nanoseconds to be a multiple "+
				"of 1000 (microsecond precision), got %d",
			now.Nanosecond(),
		)
	}
}

func TestTimeNow_ReturnsCurrentTime(t *testing.T) {
	before := time.Now().Add(-time.Millisecond)
	got := TimeNow()
	after := time.Now().Add(time.Millisecond)

	if got.Before(before) || got.After(after) {
		t.Errorf("TimeNow() = %v; want between %v and %v", got, before, after)
	}
}
