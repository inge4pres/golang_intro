package format

import (
	"strings"
	"testing"
)

func TestMajorMinorFormatHasChars(t *testing.T) {
	result := MajorMinorFormat(" Hello Go ")
	if !strings.Contains(result, "<<") || !strings.Contains(result, ">>") {
		t.Error("expected delimiters missing")
	}
}

func TestMessageInABox(t *testing.T) {
	result := MessageInABox("test")
	if !strings.Contains(result, emptyLine()) {
		t.Error("expected at least one empty line")
	}
}

func BenchmarkEmptyline(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = emptyLine()
	}
}
