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

func BenchmarkEmptyline(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = emptyLine()
	}
}
