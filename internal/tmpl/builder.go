package tmpl

import (
	_ "embed"
	"strings"
	"text/template"
)

//go:embed denied.html
var accessDenied string

//go:embed unavailable.html
var serviceUnavailable string

// Encode the text to prevent XSS (replace \n with <br> and escape HTML)
func encodeText(text string) string {
	lines := strings.Split(text, "\n")
	encodedLines := make([]string, len(lines))
	for i, m := range lines {
		encodedLines[i] = template.HTMLEscapeString(m)
	}
	return strings.Join(encodedLines, "<br>")
}

func Build401Page(msg string) []byte {
	accessDeniedWithDetails := strings.ReplaceAll(accessDenied, "{{details}}", encodeText(msg))
	return []byte(accessDeniedWithDetails)
}
func Build503Page(msg string) []byte {
	accessDeniedWithDetails := strings.ReplaceAll(serviceUnavailable, "{{details}}", encodeText(msg))
	return []byte(accessDeniedWithDetails)
}
