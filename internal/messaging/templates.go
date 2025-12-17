package messaging

import "strings"

// RenderTemplate replaces variables in message templates
func RenderTemplate(template, name string) string {
	msg := strings.ReplaceAll(template, "{{name}}", name)

	// LinkedIn message limit safety (300 chars)
	if len(msg) > 300 {
		return msg[:300]
	}
	return msg
}
