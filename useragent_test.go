package anonymize

import "testing"

func TestParseUA(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		wantBrowser string
		wantVersion string
		wantOS      string
		wantMobile  bool
	}{
		{
			name:        "empty",
			in:          "",
			wantBrowser: "",
			wantVersion: "",
			wantOS:      "",
			wantMobile:  false,
		},
		{
			name:        "chrome on windows",
			in:          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			wantBrowser: "Chrome",
			wantVersion: "120",
			wantOS:      "Windows",
			wantMobile:  false,
		},
		{
			name:        "safari on ios",
			in:          "Mozilla/5.0 (iPhone; CPU iPhone OS 17_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Mobile/15E148 Safari/604.1",
			wantBrowser: "Mobile Safari",
			wantVersion: "17",
			wantOS:      "iOS",
			wantMobile:  true,
		},
		{
			name:        "firefox on linux",
			in:          "Mozilla/5.0 (X11; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0",
			wantBrowser: "Firefox",
			wantVersion: "121",
			wantOS:      "Linux",
			wantMobile:  false,
		},
		{
			name:        "googlebot",
			in:          "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			wantBrowser: "Googlebot",
			wantVersion: "2",
			wantOS:      "Other",
			wantMobile:  false,
		},
		{
			name:        "chrome on android",
			in:          "Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			wantBrowser: "Chrome Mobile",
			wantVersion: "120",
			wantOS:      "Android",
			wantMobile:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			browser, version, os, mobile := ParseUA(tc.in)
			if browser != tc.wantBrowser {
				t.Errorf("browser = %q, want %q", browser, tc.wantBrowser)
			}
			if version != tc.wantVersion {
				t.Errorf("version = %q, want %q", version, tc.wantVersion)
			}
			if os != tc.wantOS {
				t.Errorf("os = %q, want %q", os, tc.wantOS)
			}
			if mobile != tc.wantMobile {
				t.Errorf("mobile = %v, want %v", mobile, tc.wantMobile)
			}
		})
	}
}
