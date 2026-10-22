package ui

import (
	"strings"
	"testing"
)

func TestScrollbar_GenerateTrack(t *testing.T) {
	s := NewScrollbar(true)

	// Short content: all track
	track := s.GenerateTrack(10, 5, 0)
	if len(track) != 10 {
		t.Fatalf("expected track length 10, got %d", len(track))
	}
	for _, ch := range track {
		if !strings.Contains(ch, "│") {
			t.Errorf("expected track char │, got %s", ch)
		}
	}

	// Long content at top: thumb at top
	trackTop := s.GenerateTrack(10, 100, 0)
	if len(trackTop) != 10 {
		t.Fatalf("expected track length 10, got %d", len(trackTop))
	}
	if !strings.Contains(trackTop[0], "█") {
		t.Errorf("expected thumb █ at index 0 for top offset")
	}

	// Long content at bottom: thumb at bottom
	trackBot := s.GenerateTrack(10, 100, 90)
	if !strings.Contains(trackBot[9], "█") {
		t.Errorf("expected thumb █ at bottom for bottom offset")
	}
}

func TestScrollbar_CombineWithViewport(t *testing.T) {
	s := NewScrollbar(true)
	view := "Line 1\nLine 2\nLine 3"
	combined := s.CombineWithViewport(view, 20, 3, 10, 0)

	lines := strings.Split(combined, "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}

	for _, line := range lines {
		if !strings.Contains(line, "█") && !strings.Contains(line, "│") {
			t.Errorf("expected line to contain scrollbar character: %q", line)
		}
	}
}

func TestStatusBar_View(t *testing.T) {
	sb := NewStatusBar("test.md", true)
	view := sb.View(80, 0.45, 45, 100)

	if !strings.Contains(view, "test.md") {
		t.Errorf("expected status bar to contain title 'test.md'")
	}
	if !strings.Contains(view, "45%") {
		t.Errorf("expected status bar to contain percentage '45%%'")
	}
	if !strings.Contains(view, "Ctrl+Q Quit") {
		t.Errorf("expected status bar to contain shortcut hint 'Ctrl+Q Quit', got: %s", view)
	}
}
