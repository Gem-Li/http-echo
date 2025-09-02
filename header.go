// Author: f210 (508699929@qq.com)
// Date: 2025/09/02
// Description header.go:

package main

import (
	"net/textproto"
	"sort"
	"strings"
)

func FormatHeader(h map[string][]string) map[string]string {
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, textproto.CanonicalMIMEHeaderKey(k))
	}
	sort.Strings(keys)

	headers := make(map[string]string, len(keys))
	for _, k := range keys {
		headers[k] = strings.Join(h[k], ", ")
	}

	return headers
}
