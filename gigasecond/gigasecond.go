package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(1e9 * time.Second)
	return t
}
