package messaging

import "strings"

// RenderTemplate personalizes a message template
// It safely replaces variables and enforces LinkedIn limits
func RenderTemplate(template, name string) string {

	// Replace placeholder with actual profile name
	msg := strings.ReplaceAll(template, "{{name}}", name)

	// -------------------------------------------------
	// MESSAGE LENGTH SAFETY
	// -------------------------------------------------
	// LinkedIn messages are limited in length
	// Trim message to avoid rejection or errors
	if len(msg) > 300 {
		return msg[:300]
	}

	return msg
}
