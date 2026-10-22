package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestExitDialogNavigation(t *testing.T) {
	d := NewExitDialog(true)

	// Default should be ButtonYes
	if d.Selected != ButtonYes {
		t.Errorf("expected default to be ButtonYes, got %v", d.Selected)
	}

	// Press enter on default (ButtonYes) -> should confirm
	action := d.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if action != ActionConfirm {
		t.Errorf("expected ActionConfirm on default ButtonYes enter, got %v", action)
	}

	// Press left arrow (moves to ButtonNo on the left)
	action = d.Update(tea.KeyMsg{Type: tea.KeyLeft})
	if action != ActionNone {
		t.Errorf("expected ActionNone, got %v", action)
	}
	if d.Selected != ButtonNo {
		t.Errorf("expected selection to switch to ButtonNo, got %v", d.Selected)
	}

	// Press enter on ButtonNo -> should cancel
	action = d.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if action != ActionCancel {
		t.Errorf("expected ActionCancel on ButtonNo enter, got %v", action)
	}

	// Press right arrow (moves back to ButtonYes on the right)
	action = d.Update(tea.KeyMsg{Type: tea.KeyRight})
	if action != ActionNone {
		t.Errorf("expected ActionNone, got %v", action)
	}
	if d.Selected != ButtonYes {
		t.Errorf("expected selection to switch to ButtonYes, got %v", d.Selected)
	}

	// Tab toggles to ButtonNo
	d.Update(tea.KeyMsg{Type: tea.KeyTab})
	if d.Selected != ButtonNo {
		t.Errorf("expected selection to switch to ButtonNo after Tab, got %v", d.Selected)
	}

	// Esc should cancel
	action = d.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if action != ActionCancel {
		t.Errorf("expected ActionCancel on Esc, got %v", action)
	}
}

func TestExitDialogView(t *testing.T) {
	d := NewExitDialog(true)
	view := d.View(80, 24)

	if !strings.Contains(view, "Are you sure you want to quit?") {
		t.Errorf("expected view to contain question text")
	}

	// Verify order: No appears before Yes
	idxNo := strings.Index(view, "No")
	idxYes := strings.Index(view, "Yes")
	if idxNo == -1 || idxYes == -1 {
		t.Fatalf("expected both button labels in view")
	}
	if idxNo > idxYes {
		t.Errorf("expected 'No' to appear before 'Yes' in view, got idxNo=%d, idxYes=%d", idxNo, idxYes)
	}
}
