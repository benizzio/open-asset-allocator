// Package urlutil provides URL query inspection helpers that preserve parameter presence,
// including empty values and queries with malformed value encoding.
// Authored by: GPT-6.1 Sol
package urlutil

import (
	"net/url"
	"strings"
)

// HasQueryParameter reports parameter presence independently of whether its value is empty.
// For example, HasQueryParameter(url.Values{"pageCursor": {""}}, "pageCursor") returns true.
// Co-authored by: GPT-6.1 Sol and OpenCode
func HasQueryParameter(query url.Values, name string) bool {
	var _, exists = query[name]
	return exists
}

// RawQueryHasParameter checks decoded query keys even when a malformed value prevents normal parsing.
// Invalidly encoded keys are ignored. The raw query must not include a leading question mark.
// For example, RawQueryHasParameter("page%43ursor=%ZZ", "pageCursor") returns true.
// Co-authored by: GPT-6.1 Sol and OpenCode
func RawQueryHasParameter(rawQuery string, name string) bool {
	for _, pair := range strings.Split(rawQuery, "&") {
		var rawName = strings.SplitN(pair, "=", 2)[0]
		var decodedName, err = url.QueryUnescape(rawName)
		if err == nil && decodedName == name {
			return true
		}
	}
	return false
}
