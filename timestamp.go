package anonymize

import "time"

// RoundTimestamp truncates t to the nearest preceding minute, discarding any
// sub-minute precision. This removes fingerprinting risk from high-resolution
// timestamps while keeping analytics data usable at minute granularity.
func RoundTimestamp(t time.Time) time.Time {
	return t.Truncate(time.Minute)
}
