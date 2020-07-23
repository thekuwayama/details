package format

import "testing"

func TestHeaderWithEmpty(t *testing.T) {
	s := ""
	h := Header(&s)
	expected := "<details>\n\n```"

	if h != expected {
		t.Errorf("Got %v expected %v", h, expected)
	}
}

func TestHeaderWithNil(t *testing.T) {
	h := Header(nil)
	expected := "<details>\n\n```"

	if h != expected {
		t.Errorf("Got %v expected %v", h, expected)
	}
}

func TestHeaderWithPlaceholder(t *testing.T) {
	s := "placeholder"
	h := Header(&s)
	expected := "<details><summary>placeholder</summary>\n\n```"

	if h != expected {
		t.Errorf("Got %v expected %v", h, expected)
	}
}
