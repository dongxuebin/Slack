package webscan

import (
	"fmt"
	"slack-wails/lib/clients"
	"slack-wails/lib/util"
	"strings"
	"testing"
)

func TestInfoscan(t *testing.T) {
	InitFingprintDB(util.HomeDir() + "/slack/config/webfinger.yaml")
	title := ""
	server := ""
	ct := ""
	banner := ""
	resp, body, _ := clients.NewRequest("GET", "https://oa.shanghai-cjsw.com/", nil, nil, 10, clients.DefaultClient())
	if match := util.RegTitle.FindSubmatch(body); len(match) > 1 {
		title = util.Str2UTF8(string(match[1]))
	}
	for k, v := range resp.Header {
		if k == "Server" {
			server = strings.Join(v, ";")
		}
		if k == "Content-Type" {
			ct = strings.Join(v, ";")
		}
	}
	headers, _, _ := DumpResponseHeadersAndRaw(resp)

	result := FingerScan(&TargetINFO{
		HeadeString:   string(headers),
		ContentType:   ct,
		Cert:          GetTLSString("https", "oa.shanghai-cjsw.com:443"),
		BodyString:    string(body),
		Path:          "/",
		Title:         title,
		Server:        server,
		ContentLength: len(body),
		Port:          443,
		IconHash:      FaviconHash("https://oa.shanghai-cjsw.com/", clients.DefaultClient()),
		StatusCode:    resp.StatusCode,
		Banner:        banner,
	}, FingerprintDB)
	fmt.Printf("result: %v\n", result)
}