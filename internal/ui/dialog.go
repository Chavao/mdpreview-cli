package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// DialogButton represents the button selection in the confirmation dialog.
type DialogButton int

const (
	ButtonYes DialogButton = iota
	ButtonNo
)

// DialogAction represents the outcome of a key press in the dialog.
type DialogAction int

const (
	ActionNone DialogAction = iota
	ActionConfirm
	ActionCancel
)

// ExitDialog represents the exit confirmation modal dialog.
type ExitDialog struct {
	Selected DialogButton
	isDark   bool
}

// NewExitDialog creates a new exit dialog with default selection on ButtonYes.
func NewExitDialog(isDark bool) *ExitDialog {
	return &ExitDialog{
		Selected: ButtonYes,
		isDark:   isDark,
	}
}

// Update handles key events when the dialog is active.
func (d *ExitDialog) Update(msg tea.KeyMsg) DialogAction {
	switch msg.String() {
	case "left", "h":
		d.Selected = ButtonNo
		return ActionNone

	case "right", "l":
		d.Selected = ButtonYes
		return ActionNone

	case "tab", "shift+tab":
		if d.Selected == ButtonYes {
			d.Selected = ButtonNo
		} else {
			d.Selected = ButtonYes
		}
		return ActionNone

	case "enter":
		if d.Selected == ButtonYes {
			return ActionConfirm
		}
		return ActionCancel

	case "esc":
		return ActionCancel

	case "s", "S", "y", "Y":
		d.Selected = ButtonYes
		return ActionConfirm

	case "n", "N":
		d.Selected = ButtonNo
		return ActionCancel
	}

	return ActionNone
}

// View renders the modal box centered within the given terminal dimensions.
func (d *ExitDialog) View(termWidth, termHeight int) string {
	accentColor := lipgloss.Color("#1d4ed8")
	borderColor := lipgloss.Color("#cbd5e1")
	boxBg := lipgloss.Color("#ffffff")
	textColor := lipgloss.Color("#111827")
	dimBtnBg := lipgloss.Color("#f3f4f6")
	dimBtnText := lipgloss.Color("#374151")

	if d.isDark {
		accentColor = lipgloss.Color("#8ab4ff")
		borderColor = lipgloss.Color("#333333")
		boxBg = lipgloss.Color("#181818")
		textColor = lipgloss.Color("#ffffff")
		dimBtnBg = lipgloss.Color("#262626")
		dimBtnText = lipgloss.Color("#a3a3a3")
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(textColor).
		MarginBottom(1).
		Align(lipgloss.Center)

	activeBtnStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ffffff")).
		Background(accentColor).
		Padding(0, 2).
		Margin(0, 1)

	if d.isDark {
		activeBtnStyle = activeBtnStyle.Foreground(lipgloss.Color("#000000"))
	}

	inactiveBtnStyle := lipgloss.NewStyle().
		Foreground(dimBtnText).
		Background(dimBtnBg).
		Padding(0, 2).
		Margin(0, 1)

	var yesBtn, noBtn string
	if d.Selected == ButtonYes {
		yesBtn = activeBtnStyle.Render("[ Yes ]")
		noBtn = inactiveBtnStyle.Render("  No  ")
	} else {
		yesBtn = inactiveBtnStyle.Render("  Yes  ")
		noBtn = activeBtnStyle.Render("[ No ]")
	}

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, noBtn, yesBtn)

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render("Are you sure you want to quit?"),
		buttons,
	)

	dialogBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Background(boxBg).
		Padding(1, 4).
		Render(content)

	if termWidth <= 0 || termHeight <= 0 {
		return dialogBox
	}

	// Center dialog over background
	return lipgloss.Place(
		termWidth,
		termHeight,
		lipgloss.Center,
		lipgloss.Center,
		dialogBox,
		lipgloss.WithWhitespaceChars(" "),
	)
}
