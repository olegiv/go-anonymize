package anonymize

import (
	"testing"
	"time"
)

func TestAnonymize(t *testing.T) {
	ip := "192.168.1.100"
	ua := "Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36"
	referer := "https://www.example.com/landing?utm=email"
	ts := time.Date(2026, 4, 10, 14, 23, 45, 123456789, time.UTC)

	got := Anonymize(ip, ua, referer, ts)

	want := Result{
		IP:             "192.168.1.0",
		Browser:        "Chrome Mobile",
		BrowserVersion: "120",
		OS:             "Android",
		Mobile:         true,
		Timestamp:      time.Date(2026, 4, 10, 14, 23, 0, 0, time.UTC),
		Referer:        "www.example.com",
	}

	if got.IP != want.IP {
		t.Errorf("IP = %q, want %q", got.IP, want.IP)
	}
	if got.Browser != want.Browser {
		t.Errorf("Browser = %q, want %q", got.Browser, want.Browser)
	}
	if got.BrowserVersion != want.BrowserVersion {
		t.Errorf("BrowserVersion = %q, want %q", got.BrowserVersion, want.BrowserVersion)
	}
	if got.OS != want.OS {
		t.Errorf("OS = %q, want %q", got.OS, want.OS)
	}
	if got.Mobile != want.Mobile {
		t.Errorf("Mobile = %v, want %v", got.Mobile, want.Mobile)
	}
	if !got.Timestamp.Equal(want.Timestamp) {
		t.Errorf("Timestamp = %v, want %v", got.Timestamp, want.Timestamp)
	}
	if got.Referer != want.Referer {
		t.Errorf("Referer = %q, want %q", got.Referer, want.Referer)
	}
}
