package anonymize

import "net"

// MaskIP returns an anonymized form of an IP address suitable for analytics
// storage. IPv4 addresses have the final octet zeroed (192.168.1.100 becomes
// 192.168.1.0). IPv6 addresses have their lower 80 bits zeroed, preserving
// only the /48 routing prefix. Invalid or empty input yields an empty string.
func MaskIP(raw string) string {
	ip := net.ParseIP(raw)
	if ip == nil {
		return ""
	}

	if v4 := ip.To4(); v4 != nil {
		masked := make(net.IP, net.IPv4len)
		copy(masked, v4)
		masked[3] = 0
		return masked.String()
	}

	masked := make(net.IP, net.IPv6len)
	copy(masked, ip.To16())
	for i := 6; i < net.IPv6len; i++ {
		masked[i] = 0
	}
	return masked.String()
}
