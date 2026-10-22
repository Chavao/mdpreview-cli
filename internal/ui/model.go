package ui

import (
	"strings"

	"github.com/Chavao/mdpreview-cli/internal/render"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

// Model represents the main Bubble Tea model for the Markdown previewer.
type Model struct {
	viewport        viewport.Model
	renderer        *render.Renderer
	dialog          *ExitDialog
	statusbar       *StatusBar
	scrollbar       *Scrollbar
	title           string
	rawMarkdown     string
	renderedContent string
	width           int
	height          int
	ready           bool
	isDark          bool
	showDialog      bool
	quitting        bool
}

// NewModel creates a new previewer Model.
func NewModel(rawMarkdown, title string, isDark bool) Model {
	if title == "" {
		title = "[stdin]"
	}

	return Model{
		renderer:    render.NewRenderer(isDark),
		dialog:      NewExitDialog(isDark),
		statusbar:   NewStatusBar(title, isDark),
		scrollbar:   NewScrollbar(isDark),
		title:       title,
		rawMarkdown: rawMarkdown,
		isDark:      isDark,
		showDialog:  false,
	}
}

// Init initializes the Bubbletea program.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles incoming messages and events.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// When dialog is open, all keys go to the dialog
		if m.showDialog {
			action := m.dialog.Update(msg)
			switch action {
			case ActionConfirm:
				m.quitting = true
				return m, tea.Quit
			case ActionCancel:
				m.showDialog = false
				return m, nil
			default:
				return m, nil
			}
		}

		// Normal previewer key handling
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "ctrl+q", "q":
			m.dialog.Selected = ButtonYes
			m.showDialog = true
			return m, nil

		case "g", "home":
			m.viewport.GotoTop()
			return m, nil

		case "G", "end":
			m.viewport.GotoBottom()
			return m, nil

		case "ctrl+d":
			m.viewport.HalfViewDown()
			return m, nil

		case "ctrl+u":
			m.viewport.HalfViewUp()
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		contentWidth := m.width - 3
		if contentWidth < 20 {
			contentWidth = 20
		}

		rendered, err := m.renderer.Render(m.rawMarkdown, contentWidth)
		if err == nil {
			m.renderedContent = rendered
		}

		viewportHeight := m.height - 1
		if viewportHeight < 1 {
			viewportHeight = 1
		}

		if !m.ready {
			m.viewport = viewport.New(m.width-1, viewportHeight)
			m.viewport.YPosition = 0
			m.viewport.SetContent(m.renderedContent)
			m.ready = true
		} else {
			m.viewport.Width = m.width - 1
			m.viewport.Height = viewportHeight
			m.viewport.SetContent(m.renderedContent)
		}

	case tea.MouseMsg:
		if m.showDialog {
			return m, nil
		}
	}

	// Update viewport with scroll events (keyboard and mouse)
	if !m.showDialog {
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// View renders the TUI screen.
func (m Model) View() string {
	if !m.ready {
		return "Loading..."
	}

	if m.quitting {
		return ""
	}

	totalLines := m.viewport.TotalLineCount()
	currentLine := m.viewport.YOffset + 1
	if totalLines == 0 {
		currentLine = 0
	}

	// Combine viewport lines with vertical scrollbar on the right
	vpView := m.viewport.View()
	contentWithScrollbar := m.scrollbar.CombineWithViewport(
		vpView,
		m.width,
		m.viewport.Height,
		totalLines,
		m.viewport.YOffset,
	)

	// Bottom status bar
	barView := m.statusbar.View(m.width, m.viewport.ScrollPercent(), currentLine, totalLines)

	baseView := contentWithScrollbar + "\n" + barView

	// If exit dialog is active, overlay it centered on the base view
	if m.showDialog {
		return m.overlayDialog(baseView)
	}

	return baseView
}

// overlayDialog places the modal box centered over the base view.
func (m Model) overlayDialog(baseView string) string {
	dialogBox := m.dialog.View(0, 0)
	dialogLines := strings.Split(dialogBox, "\n")
	dialogH := len(dialogLines)
	dialogW := 0
	for _, l := range dialogLines {
		w := lipgloss.Width(l)
		if w > dialogW {
			dialogW = w
		}
	}

	baseLines := strings.Split(baseView, "\n")
	top := (m.height - dialogH) / 2
	left := (m.width - dialogW) / 2
	if top < 0 {
		top = 0
	}
	if left < 0 {
		left = 0
	}

	out := make([]string, len(baseLines))
	for i, line := range baseLines {
		if i >= top && i < top+dialogH {
			dIdx := i - top
			dLine := dialogLines[dIdx]

			leftPart := ""
			if left > 0 {
				leftPart = ansi.Truncate(line, left, "")
				leftWidth := lipgloss.Width(leftPart)
				if leftWidth < left {
					leftPart += strings.Repeat(" ", left-leftWidth)
				}
			}

			rightStart := left + dialogW
			rightPart := ""
			totalW := lipgloss.Width(line)
			if totalW > rightStart {
				rightPart = ansi.TruncateLeft(line, rightStart, "")
			}

			out[i] = leftPart + dLine + rightPart
		} else {
			out[i] = line
		}
	}

	return strings.Join(out, "\n")
}
