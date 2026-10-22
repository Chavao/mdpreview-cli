package render

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
)

func TestRenderer_Render(t *testing.T) {
	md := `# Main Title

This is a **bold** paragraph with some ` + "`inline code`" + ` and a [link](https://example.com).

> [!NOTE]
> This is a special note alert!

` + "```go\npackage main\n\nfunc main() {}\n```\n"

	// Test Dark mode
	darkRenderer := NewRenderer(true)
	darkOut, err := darkRenderer.Render(md, 80)
	if err != nil {
		t.Fatalf("dark render failed: %v", err)
	}

	strippedDark := ansi.Strip(darkOut)
	if !strings.Contains(strippedDark, "NOTE") {
		t.Errorf("expected dark output to contain NOTE alert header")
	}
	if !strings.Contains(strippedDark, "special note alert") {
		t.Errorf("expected dark output to contain alert body")
	}
	if !strings.Contains(strippedDark, "Main Title") {
		t.Errorf("expected dark output to contain Main Title")
	}

	// Test Light mode
	lightRenderer := NewRenderer(false)
	lightOut, err := lightRenderer.Render(md, 80)
	if err != nil {
		t.Fatalf("light render failed: %v", err)
	}
	strippedLight := ansi.Strip(lightOut)
	if !strings.Contains(strippedLight, "NOTE") {
		t.Errorf("expected light output to contain NOTE alert header")
	}

	// Test Empty input
	emptyOut, err := darkRenderer.Render("", 80)
	if err != nil {
		t.Fatalf("empty render failed: %v", err)
	}
	strippedEmpty := ansi.Strip(emptyOut)
	if !strings.Contains(strippedEmpty, "No markdown content") {
		t.Errorf("expected empty message, got %q", strippedEmpty)
	}
}
