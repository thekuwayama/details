package format

import "fmt"

// Header returns the string contains details opening tag.
func Header(summary *string) string {
	if summary == nil || len(*summary) == 0 {
		return "<details>\n\n```"
	}

	return fmt.Sprintf("<details><summary>%s</summary>\n\n```", *summary)
}

// Footer returns the string contains details closing tag.
func Footer() string {
	return "```\n\n</details>"
}
