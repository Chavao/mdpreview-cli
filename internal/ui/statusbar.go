package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// StatusBar renders the bottom status bar for the markdown previewer.
type StatusBar struct {
	Title  string
	isDark bool
}

// NewStatusBar creates a new status bar.
func NewStatusBar(title string, isDark bool) *StatusBar {
	if title == "" {
		title = "[stdin]"
	}
	return &StatusBar{
		Title:  title,
		isDark: isDark,
	}
}

// View renders the status bar for the given terminal width and scroll percent.
func (s *StatusBar) View(width int, scrollPercent float64, currentLine, totalLines int) string {
	if width <= 0 {
		return ""
	}

	badgeBg := lipgloss.Color("#1d4ed8")
	badgeFg := lipgloss.Color("#ffffff")
	barBg := lipgloss.Color("#e5e7eb")
	barFg := lipgloss.Color("#374151")
	hintFg := lipgloss.Color("#6b7280")

	if s.isDark {
		badgeBg = lipgloss.Color("#2563eb")
		badgeFg = lipgloss.Color("#ffffff")
		barBg = lipgloss.Color("#1a1a1a")
		barFg = lipgloss.Color("#e5e7eb")
		hintFg = lipgloss.Color("#888888")
	}

	badgeStyle := lipgloss.NewStyle().
		Bold(true).
		Background(badgeBg).
		Foreground(badgeFg).
		Padding(0, 1)

	badge := badgeStyle.Render(s.Title)

	// Format progress
	var progressText string
	switch {
	case totalLines <= 0 || scrollPercent <= 0.0:
		progressText = "TOP"
	case scrollPercent >= 1.0:
		progressText = "BOT"
	default:
		progressText = fmt.Sprintf("%2.0f%%", scrollPercent*100)
	}

	posText := ""
	if totalLines > 0 {
		posText = fmt.Sprintf(" %s (%d/%d)", progressText, currentLine, totalLines)
	} else {
		posText = fmt.Sprintf(" %s", progressText)
	}

	hintText := "↑/↓ Scroll • Mouse • Ctrl+Q Quit"

	posStyle := lipgloss.NewStyle().
		Foreground(barFg).
		Background(barBg).
		Bold(true)

	hintStyle := lipgloss.NewStyle().
		Foreground(hintFg).
		Background(barBg)

	rightContent := lipgloss.JoinHorizontal(
		lipgloss.Center,
		hintStyle.Render(hintText),
		posStyle.Render(posText),
	)

	leftLen := lipgloss.Width(badge)
	rightLen := lipgloss.Width(rightContent)
	gapLen := width - leftLen - rightLen - 1
	if gapLen < 0 {
		gapLen = 0
	}

	gap := lipgloss.NewStyle().
		Background(barBg).
		Render(fmt.Sprintf("%*s", gapLen, ""))

	barStyle := lipgloss.NewStyle().
		Background(barBg).
		Width(width)

	return barStyle.Render(lipgloss.JoinHorizontal(lipgloss.Center, badge, gap, rightContent))
}
