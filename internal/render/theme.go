package render

import (
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/glamour/styles"
)

// stringPtr returns a pointer to a string.
func stringPtr(s string) *string {
	return &s
}

// uintPtr returns a pointer to a uint.
func uintPtr(u uint) *uint {
	return &u
}

// boolPtr returns a pointer to a bool.
func boolPtr(b bool) *bool {
	return &b
}

// NewDarkStyleConfig creates an ansi.StyleConfig inspired by mdpreview's dark theme.
func NewDarkStyleConfig() ansi.StyleConfig {
	cfg := styles.DarkStyleConfig

	// Base document styling
	cfg.Document.Color = stringPtr("#ffffff")
	cfg.Document.Margin = uintPtr(2)

	// Headings
	cfg.Heading.Bold = boolPtr(true)
	cfg.Heading.Color = stringPtr("#ffffff")

	cfg.H1.Prefix = " "
	cfg.H1.Suffix = " "
	cfg.H1.Color = stringPtr("#ffffff")
	cfg.H1.BackgroundColor = stringPtr("#1f3a60")
	cfg.H1.Bold = boolPtr(true)

	cfg.H2.Prefix = "## "
	cfg.H2.Color = stringPtr("#8ab4ff")
	cfg.H2.Bold = boolPtr(true)

	cfg.H3.Prefix = "### "
	cfg.H3.Color = stringPtr("#ffffff")
	cfg.H3.Bold = boolPtr(true)

	// Links with mdpreview accent
	cfg.Link.Color = stringPtr("#8ab4ff")
	cfg.Link.Underline = boolPtr(true)
	cfg.LinkText.Color = stringPtr("#8ab4ff")
	cfg.LinkText.Bold = boolPtr(true)

	// Inline code
	cfg.Code.Color = stringPtr("#ffffff")
	cfg.Code.BackgroundColor = stringPtr("#2c313a")

	// Block quotes
	cfg.BlockQuote.Indent = uintPtr(1)
	cfg.BlockQuote.IndentToken = stringPtr("│ ")
	cfg.BlockQuote.Color = stringPtr("#b7b7b7")

	// Horizontal rule
	cfg.HorizontalRule.Color = stringPtr("#3a3a3a")

	// Code block background
	cfg.CodeBlock.Chroma.Background.BackgroundColor = stringPtr("#1a1a1a")

	return cfg
}

// NewLightStyleConfig creates an ansi.StyleConfig inspired by mdpreview's light theme.
func NewLightStyleConfig() ansi.StyleConfig {
	cfg := styles.LightStyleConfig

	// Base document styling
	cfg.Document.Color = stringPtr("#1f2937")
	cfg.Document.Margin = uintPtr(2)

	// Headings
	cfg.Heading.Bold = boolPtr(true)
	cfg.Heading.Color = stringPtr("#111827")

	cfg.H1.Prefix = " "
	cfg.H1.Suffix = " "
	cfg.H1.Color = stringPtr("#ffffff")
	cfg.H1.BackgroundColor = stringPtr("#1d4ed8")
	cfg.H1.Bold = boolPtr(true)

	cfg.H2.Prefix = "## "
	cfg.H2.Color = stringPtr("#1d4ed8")
	cfg.H2.Bold = boolPtr(true)

	cfg.H3.Prefix = "### "
	cfg.H3.Color = stringPtr("#111827")
	cfg.H3.Bold = boolPtr(true)

	// Links with mdpreview accent
	cfg.Link.Color = stringPtr("#1d4ed8")
	cfg.Link.Underline = boolPtr(true)
	cfg.LinkText.Color = stringPtr("#1d4ed8")
	cfg.LinkText.Bold = boolPtr(true)

	// Inline code
	cfg.Code.Color = stringPtr("#111827")
	cfg.Code.BackgroundColor = stringPtr("#e5e7eb")

	// Block quotes
	cfg.BlockQuote.Indent = uintPtr(1)
	cfg.BlockQuote.IndentToken = stringPtr("│ ")
	cfg.BlockQuote.Color = stringPtr("#6b7280")

	// Horizontal rule
	cfg.HorizontalRule.Color = stringPtr("#d1d5db")

	return cfg
}
