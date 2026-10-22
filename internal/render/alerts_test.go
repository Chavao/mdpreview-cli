package render

import (
	"strings"
	"testing"
)

func TestExtractAlerts(t *testing.T) {
	md := `# Main Title

> [!NOTE]
> This is a note with important info.
> And a second line.

Some regular text here.

> [!WARNING]
> Watch out!

A regular blockquote:
> Just a quote
> Nothing special

` + "```markdown\n> [!TIP]\n> Code tip\n```\n"

	extractedMd, alerts := ExtractAlerts(md)

	if len(alerts) != 2 {
		t.Fatalf("expected 2 alerts, got %d", len(alerts))
	}

	if alerts[0].Alert.Type != AlertNote {
		t.Errorf("expected first alert to be NOTE, got %s", alerts[0].Alert.Type)
	}
	expectedBody0 := "This is a note with important info.\nAnd a second line."
	if alerts[0].Alert.Body != expectedBody0 {
		t.Errorf("expected body %q, got %q", expectedBody0, alerts[0].Alert.Body)
	}

	if alerts[1].Alert.Type != AlertWarning {
		t.Errorf("expected second alert to be WARNING, got %s", alerts[1].Alert.Type)
	}
	if alerts[1].Alert.Body != "Watch out!" {
		t.Errorf("expected body 'Watch out!', got %q", alerts[1].Alert.Body)
	}

	// Verify placeholder exists in extractedMd
	if !strings.Contains(extractedMd, alerts[0].Placeholder) {
		t.Errorf("expected extractedMd to contain placeholder %s", alerts[0].Placeholder)
	}

	// Verify code block tip was NOT extracted
	if strings.Contains(extractedMd, "MDPREVIEWALERTTIP") {
		t.Errorf("alert inside code block should not have been extracted")
	}

	// Verify regular quote was preserved
	if !strings.Contains(extractedMd, "> Just a quote") {
		t.Errorf("regular quote should have been preserved")
	}
}

func TestRenderAlert(t *testing.T) {
	alert := AlertBlock{
		Type: AlertImportant,
		Body: "Key takeaway message",
	}

	rendered := RenderAlert(alert, alert.Body, true, 80)
	if !strings.Contains(rendered, "IMPORTANT") {
		t.Errorf("rendered alert should contain title IMPORTANT")
	}
	if !strings.Contains(rendered, "Key takeaway message") {
		t.Errorf("rendered alert should contain body")
	}
	if !strings.Contains(rendered, "▌") {
		t.Errorf("rendered alert should have border marker ▌")
	}
}
