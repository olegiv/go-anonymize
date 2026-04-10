package anonymize

import "net/url"

// ExtractDomain returns only the hostname component of a referer URL,
// stripping scheme, path, query, and fragment. Empty input, parse errors,
// or URLs without a host all return an empty string.
func ExtractDomain(raw string) string {
	if raw == "" {
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	return u.Hostname()
}
