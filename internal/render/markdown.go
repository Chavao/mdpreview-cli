package render

import (
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// Renderer renders Markdown content into styled terminal ANSI output.
type Renderer struct {
	isDark bool
}

// NewRenderer creates a new Markdown renderer with dark/light mode configured.
func NewRenderer(isDark bool) *Renderer {
	return &Renderer{
		isDark: isDark,
	}
}

// Render converts Markdown text into terminal formatted ANSI strings.
func (r *Renderer) Render(input string, width int) (string, error) {
	if strings.TrimSpace(input) == "" {
		emptyStyle := lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("#6b7280"))
		return emptyStyle.Render("No markdown content to preview."), nil
	}

	// 1. Extract GitHub alert blocks and replace with placeholders
	cleanMarkdown, extractedAlerts := ExtractAlerts(input)

	// Determine wrap width for terminal
	wrapWidth := width
	if wrapWidth <= 0 {
		wrapWidth = 80
	}

	// 2. Configure Glamour renderer
	styleCfg := NewLightStyleConfig()
	if r.isDark {
		styleCfg = NewDarkStyleConfig()
	}

	glamourRenderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(styleCfg),
		glamour.WithWordWrap(wrapWidth),
	)
	if err != nil {
		return "", err
	}

	mainOutput, err := glamourRenderer.Render(cleanMarkdown)
	if err != nil {
		return "", err
	}

	// 3. Render and replace extracted alerts
	if len(extractedAlerts) > 0 {
		// Renderer for inside alert boxes with narrower width
		alertWrapWidth := wrapWidth - 8
		if alertWrapWidth < 20 {
			alertWrapWidth = 20
		}

		alertBodyRenderer, err := glamour.NewTermRenderer(
			glamour.WithStyles(styleCfg),
			glamour.WithWordWrap(alertWrapWidth),
		)

		for _, item := range extractedAlerts {
			var renderedAlertBody string
			if err == nil && strings.TrimSpace(item.Alert.Body) != "" {
				bodyOut, bodyErr := alertBodyRenderer.Render(item.Alert.Body)
				if bodyErr == nil {
					renderedAlertBody = strings.TrimSpace(bodyOut)
				} else {
					renderedAlertBody = item.Alert.Body
				}
			} else {
				renderedAlertBody = item.Alert.Body
			}

			renderedAlert := RenderAlert(item.Alert, renderedAlertBody, r.isDark, wrapWidth)

			// Replace placeholder token in main output
			mainOutput = strings.Replace(mainOutput, item.Placeholder, "\n"+renderedAlert+"\n", 1)
		}
	}

	return mainOutput, nil
}
