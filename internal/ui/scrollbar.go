package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Scrollbar renders a vertical scrollbar for the viewport.
type Scrollbar struct {
	isDark bool
}

// NewScrollbar creates a new scrollbar renderer.
func NewScrollbar(isDark bool) *Scrollbar {
	return &Scrollbar{isDark: isDark}
}

// GenerateTrack generates a slice of string characters for each row in height.
func (s *Scrollbar) GenerateTrack(height, totalLines, yOffset int) []string {
	if height <= 0 {
		return nil
	}

	track := make([]string, height)

	thumbColor := lipgloss.Color("#1d4ed8")
	trackColor := lipgloss.Color("#d1d5db")
	if s.isDark {
		thumbColor = lipgloss.Color("#8ab4ff")
		trackColor = lipgloss.Color("#333333")
	}

	thumbStyle := lipgloss.NewStyle().Foreground(thumbColor)
	trackStyle := lipgloss.NewStyle().Foreground(trackColor)

	thumbChar := "█"
	trackChar := "│"

	if totalLines <= height {
		for i := 0; i < height; i++ {
			track[i] = trackStyle.Render(trackChar)
		}
		return track
	}

	// Calculate thumb height
	thumbSize := (height * height) / totalLines
	if thumbSize < 1 {
		thumbSize = 1
	}
	if thumbSize > height {
		thumbSize = height
	}

	maxOffset := totalLines - height
	if maxOffset < 1 {
		maxOffset = 1
	}

	if yOffset < 0 {
		yOffset = 0
	}
	if yOffset > maxOffset {
		yOffset = maxOffset
	}

	thumbTop := (yOffset * (height - thumbSize)) / maxOffset
	if thumbTop < 0 {
		thumbTop = 0
	}
	if thumbTop+thumbSize > height {
		thumbTop = height - thumbSize
	}

	for i := 0; i < height; i++ {
		if i >= thumbTop && i < thumbTop+thumbSize {
			track[i] = thumbStyle.Render(thumbChar)
		} else {
			track[i] = trackStyle.Render(trackChar)
		}
	}

	return track
}

// CombineWithViewport attaches the vertical scrollbar to the viewport lines.
func (s *Scrollbar) CombineWithViewport(viewportView string, width, height, totalLines, yOffset int) string {
	track := s.GenerateTrack(height, totalLines, yOffset)
	if len(track) == 0 {
		return viewportView
	}

	lines := strings.Split(viewportView, "\n")
	var b strings.Builder

	contentWidth := width - 1
	if contentWidth < 1 {
		contentWidth = 1
	}

	for i := 0; i < height; i++ {
		var line string
		if i < len(lines) {
			line = lines[i]
		}

		// Pad or truncate line to contentWidth so scrollbar stays at right edge
		lineWidth := lipgloss.Width(line)
		if lineWidth < contentWidth {
			line = line + strings.Repeat(" ", contentWidth-lineWidth)
		}

		b.WriteString(line)
		b.WriteString(track[i])
		if i < height-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
