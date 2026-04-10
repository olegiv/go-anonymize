package anonymize

import "testing"

func TestMaskIP(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"ipv4 typical", "192.168.1.100", "192.168.1.0"},
		{"ipv4 already zero last octet", "10.0.0.0", "10.0.0.0"},
		{"ipv4 all ones", "255.255.255.255", "255.255.255.0"},
		{"ipv6 global", "2001:db8::1", "2001:db8::"},
		{"ipv6 full form", "2001:db8:1234:5678:abcd:ef01:2345:6789", "2001:db8:1234::"},
		{"ipv4 in ipv6", "::ffff:192.168.1.100", "192.168.1.0"},
		{"invalid", "not-an-ip", ""},
		{"empty", "", ""},
		{"ipv4 with port should fail", "192.168.1.100:8080", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaskIP(tc.in)
			if got != tc.want {
				t.Errorf("MaskIP(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
