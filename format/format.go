package format

import "fmt"

func Header(summary *string) string {
	if summary == nil || len(*summary) == 0 {
		return "<details>\n\n```"
	}

	return fmt.Sprintf("<details><summary>%s</summary>\n\n```", *summary)
}

func Footer() string {
	return "```\n\n</details>"
}
