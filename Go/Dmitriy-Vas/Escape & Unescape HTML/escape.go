package Escape_Unescape

import (
	"html"
	"strings"
)

func Escape(s string) string {
	return strings.NewReplacer(
		`&`, "&amp;",
		`'`, "&#39;",
		`<`, "&lt;",
		`>`, "&gt;",
		`"`, "&#34;",
		).Replace(s)
}

func Unescape(s string) string {
	return html.UnescapeString(s)
}
