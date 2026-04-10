package anonymize

import "testing"

func TestExtractDomain(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"full url with path and query", "https://www.example.com/some/path?q=1", "www.example.com"},
		{"http scheme", "http://example.com", "example.com"},
		{"with port", "https://example.com:8443/api", "example.com"},
		{"ipv6 host", "https://[2001:db8::1]:8080/", "2001:db8::1"},
		{"bare domain no scheme", "example.com", ""},
		{"empty", "", ""},
		{"malformed brackets", "http://[::", ""},
		{"scheme only", "https://", ""},
		{"fragment and query", "https://example.com/#section?x=1", "example.com"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ExtractDomain(tc.in)
			if got != tc.want {
				t.Errorf("ExtractDomain(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
