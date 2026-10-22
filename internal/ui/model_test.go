package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestModel_InitAndResize(t *testing.T) {
	md := "# Title\n\nSome test content\nLine 2\nLine 3"
	m := NewModel(md, "test.md", true)

	// Before resize
	viewBefore := m.View()
	if viewBefore != "Loading..." {
		t.Errorf("expected loading view before resize, got %q", viewBefore)
	}

	// Trigger WindowSizeMsg
	newModel, cmd := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	if cmd != nil {
		t.Errorf("expected nil cmd from window resize, got %v", cmd)
	}

	m = newModel.(Model)
	if !m.ready {
		t.Errorf("expected model to be ready after window resize")
	}

	viewAfter := m.View()
	if !strings.Contains(viewAfter, "test.md") {
		t.Errorf("expected view to contain status bar with 'test.md'")
	}
	if !strings.Contains(viewAfter, "Title") {
		t.Errorf("expected view to contain rendered Title")
	}
}

func TestModel_ExitDialogFlow(t *testing.T) {
	md := "# Title\n\nContent"
	m := NewModel(md, "test.md", true)

	// Initialize with size
	newM, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = newM.(Model)

	// Press ctrl+q to open dialog
	newM, _ = m.Update(tea.KeyMsg{Type: tea.KeyCtrlQ})
	m = newM.(Model)

	if !m.showDialog {
		t.Errorf("expected showDialog to be true after ctrl+q")
	}

	viewWithDialog := m.View()
	if !strings.Contains(viewWithDialog, "Are you sure you want to quit?") {
		t.Errorf("expected view to show exit dialog modal")
	}

	// In dialog, default is ButtonYes, so Enter should confirm and quit
	newM, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = newM.(Model)
	if cmd == nil {
		t.Errorf("expected tea.Quit cmd when pressing Enter on default ButtonYes")
	}

	// Reset quitting and showDialog to test Esc
	m.quitting = false
	m.showDialog = true

	// In dialog, press Esc to cancel
	newM, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	m = newM.(Model)

	if m.showDialog {
		t.Errorf("expected showDialog to be false after Esc")
	}

	// Press q to open dialog again
	newM, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	m = newM.(Model)
	if !m.showDialog {
		t.Errorf("expected showDialog to be true after 'q'")
	}

	// Select Yes with 's' to quit
	newM, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}})
	m = newM.(Model)

	if cmd == nil {
		t.Errorf("expected tea.Quit cmd when confirming exit with 's'")
	}
}

func TestModel_Scrolling(t *testing.T) {
	var sb strings.Builder
	for i := 0; i < 50; i++ {
		sb.WriteString("Line of text to test scrolling\n")
	}
	m := NewModel(sb.String(), "long.md", true)

	newM, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 10})
	m = newM.(Model)

	// Initial offset is 0
	if m.viewport.YOffset != 0 {
		t.Errorf("expected initial YOffset to be 0, got %d", m.viewport.YOffset)
	}

	// Navigate down with down arrow
	newM, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = newM.(Model)

	if m.viewport.YOffset < 1 {
		t.Errorf("expected YOffset to increase after KeyDown, got %d", m.viewport.YOffset)
	}

	// Go to bottom with 'G'
	newM, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'G'}})
	m = newM.(Model)

	if m.viewport.YOffset == 0 {
		t.Errorf("expected YOffset to be at bottom after 'G'")
	}

	// Go to top with 'g'
	newM, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'g'}})
	m = newM.(Model)

	if m.viewport.YOffset != 0 {
		t.Errorf("expected YOffset to be 0 after 'g', got %d", m.viewport.YOffset)
	}
}
