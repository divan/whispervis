package widgets

import (
	"fmt"
	"testing"
)

func TestLoader(t *testing.T) {
	l := NewLoader(100)

	for i := 1; i < 100; i++ {
		l.Inc()
		got := l.Progress()
		if int(got) != i {
			t.Fatalf("Expect progress to be %d%%, but got %.f%%", i, got)
		}

		gotText := l.text()
		expectedText := fmt.Sprintf("Loading %d%%...", i)
		if gotText != expectedText {
			t.Fatalf("Expect text to be %s, but got %s", expectedText, gotText)
		}
	}

	// 100%
	l.Inc()
	if l.Progress() != 100.0 {
		t.Fatalf("Expect progress to be 100%%, but got %.f%%", l.Progress())
	}
	if l.text() != "Completed" {
		t.Fatalf("Expect text to be \"Completed\", but got %s", l.text())
	}
}
