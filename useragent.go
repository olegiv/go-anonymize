package anonymize

import (
	"strings"
	"sync"

	"github.com/ua-parser/uap-go/uaparser"
)

var (
	uaParser     *uaparser.Parser
	uaParserOnce sync.Once
)

func loadParser() *uaparser.Parser {
	uaParserOnce.Do(func() {
		p, err := uaparser.New()
		if err != nil {
			return
		}
		uaParser = p
	})
	return uaParser
}

// ParseUA extracts coarse-grained browser, browser major version, operating
// system, and a mobile flag from a raw User-Agent header. The major version
// is the portion before the first dot only. The mobile flag is true only
// when the OS family is iOS or Android; tablets on other OSes are treated
// as non-mobile. Empty input or parser failure yields zero values.
func ParseUA(raw string) (browser, version, os string, mobile bool) {
	if raw == "" {
		return "", "", "", false
	}
	p := loadParser()
	if p == nil {
		return "", "", "", false
	}

	client := p.Parse(raw)
	if client == nil {
		return "", "", "", false
	}

	if client.UserAgent != nil {
		browser = client.UserAgent.Family
		version = client.UserAgent.Major
		if version == "" && browser != "" {
			if idx := strings.IndexByte(client.UserAgent.Minor, '.'); idx >= 0 {
				version = client.UserAgent.Minor[:idx]
			}
		}
	}
	if client.Os != nil {
		os = client.Os.Family
		mobile = os == "iOS" || os == "Android"
	}
	return browser, version, os, mobile
}
