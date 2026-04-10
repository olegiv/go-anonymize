// Package anonymize provides pure functions for anonymizing analytics data:
// masking client IP addresses, reducing User-Agent strings to coarse
// browser/OS fields, rounding timestamps to the minute, and reducing HTTP
// Referer URLs to a bare hostname.
package anonymize

import "time"

// Result is the anonymized representation of a single analytics event.
type Result struct {
	IP             string    `json:"ip"`
	Browser        string    `json:"browser"`
	BrowserVersion string    `json:"browser_version"`
	OS             string    `json:"os"`
	Mobile         bool      `json:"mobile"`
	Timestamp      time.Time `json:"timestamp"`
	Referer        string    `json:"referer"`
}

// Anonymize applies all of the package's anonymization steps to a single
// event and returns the combined Result. It is safe to call concurrently.
func Anonymize(ip, ua, referer string, ts time.Time) Result {
	browser, version, os, mobile := ParseUA(ua)
	return Result{
		IP:             MaskIP(ip),
		Browser:        browser,
		BrowserVersion: version,
		OS:             os,
		Mobile:         mobile,
		Timestamp:      RoundTimestamp(ts),
		Referer:        ExtractDomain(referer),
	}
}
