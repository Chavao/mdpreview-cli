package render

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// AlertType represents supported GitHub alert types.
type AlertType string

const (
	AlertNote      AlertType = "NOTE"
	AlertTip       AlertType = "TIP"
	AlertImportant AlertType = "IMPORTANT"
	AlertWarning   AlertType = "WARNING"
	AlertCaution   AlertType = "CAUTION"
)

// AlertBlock contains the extracted alert information.
type AlertBlock struct {
	Type AlertType
	Body string
}

var (
	quoteLinePattern     = regexp.MustCompile(`^(>+)\s?(.*)$`)
	calloutHeaderPattern = regexp.MustCompile(`^\[!(NOTE|TIP|IMPORTANT|WARNING|CAUTION)\]\s*$`)
)

// AlertThemeColors holds the visual colors for an alert type.
type AlertColor struct {
	Border lipgloss.TerminalColor
	Title  lipgloss.TerminalColor
}

// ExtractedAlert contains the placeholder and alert data.
type ExtractedAlert struct {
	Placeholder string
	Alert       AlertBlock
}

// ExtractAlerts parses markdown, replaces alert blocks with unique placeholders,
// and returns the modified markdown along with the extracted alert blocks.
// It ignores blockquote markers inside fenced code blocks.
func ExtractAlerts(input string) (string, []ExtractedAlert) {
	lines := strings.Split(input, "\n")
	var outputLines []string
	var alerts []ExtractedAlert

	inCodeBlock := false
	var codeFence string

	inAlert := false
	var currentAlert AlertBlock
	var currentBodyLines []string
	alertIndex := 0

	flushAlert := func() {
		if !inAlert {
			return
		}
		currentAlert.Body = strings.TrimSpace(strings.Join(currentBodyLines, "\n"))
		placeholder := fmt.Sprintf("MDPREVIEWALERT%s%dX", currentAlert.Type, alertIndex)
		alerts = append(alerts, ExtractedAlert{
			Placeholder: placeholder,
			Alert:       currentAlert,
		})
		outputLines = append(outputLines, placeholder)
		inAlert = false
		currentAlert = AlertBlock{}
		currentBodyLines = nil
		alertIndex++
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for code block fence
		if strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~") {
			fence := trimmed[:3]
			if !inCodeBlock {
				flushAlert()
				inCodeBlock = true
				codeFence = fence
			} else if strings.HasPrefix(trimmed, codeFence) {
				inCodeBlock = false
				codeFence = ""
			}
			outputLines = append(outputLines, line)
			continue
		}

		if inCodeBlock {
			outputLines = append(outputLines, line)
			continue
		}

		if inAlert {
			if isQuoteLine(line) {
				clean := stripQuotePrefix(line)
				currentBodyLines = append(currentBodyLines, clean)
				continue
			}
			// End of alert block
			flushAlert()
		}

		// Check if this line starts an alert block
		if alertType, ok := parseAlertHeader(line); ok {
			inAlert = true
			currentAlert = AlertBlock{Type: alertType}
			currentBodyLines = make([]string, 0, 4)
			continue
		}

		outputLines = append(outputLines, line)
	}

	flushAlert()

	return strings.Join(outputLines, "\n"), alerts
}

func parseAlertHeader(line string) (AlertType, bool) {
	trimmed := strings.TrimSpace(line)
	match := quoteLinePattern.FindStringSubmatch(trimmed)
	if match == nil {
		return "", false
	}
	content := strings.TrimSpace(match[2])
	submatch := calloutHeaderPattern.FindStringSubmatch(content)
	if submatch == nil {
		return "", false
	}
	return AlertType(submatch[1]), true
}

func isQuoteLine(line string) bool {
	return quoteLinePattern.MatchString(strings.TrimSpace(line))
}

func stripQuotePrefix(line string) string {
	trimmed := strings.TrimSpace(line)
	match := quoteLinePattern.FindStringSubmatch(trimmed)
	if match == nil {
		return line
	}
	return match[2]
}

// AlertColors returns theme colors for a given alert type and dark mode flag.
func GetAlertColors(alertType AlertType, isDark bool) AlertColor {
	switch alertType {
	case AlertNote:
		if isDark {
			return AlertColor{Border: lipgloss.Color("#8ab4ff"), Title: lipgloss.Color("#8ab4ff")}
		}
		return AlertColor{Border: lipgloss.Color("#2563eb"), Title: lipgloss.Color("#2563eb")}
	case AlertTip:
		return AlertColor{Border: lipgloss.Color("#16a34a"), Title: lipgloss.Color("#16a34a")}
	case AlertImportant:
		return AlertColor{Border: lipgloss.Color("#8b5cf6"), Title: lipgloss.Color("#8b5cf6")}
	case AlertWarning:
		return AlertColor{Border: lipgloss.Color("#ea580c"), Title: lipgloss.Color("#ea580c")}
	case AlertCaution:
		return AlertColor{Border: lipgloss.Color("#dc2626"), Title: lipgloss.Color("#dc2626")}
	default:
		if isDark {
			return AlertColor{Border: lipgloss.Color("#8ab4ff"), Title: lipgloss.Color("#8ab4ff")}
		}
		return AlertColor{Border: lipgloss.Color("#2563eb"), Title: lipgloss.Color("#2563eb")}
	}
}

// RenderAlert renders an alert box with lipgloss.
func RenderAlert(alert AlertBlock, renderedBody string, isDark bool, width int) string {
	colors := GetAlertColors(alert.Type, isDark)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colors.Title)

	borderChar := "▌"
	borderStyle := lipgloss.NewStyle().Foreground(colors.Border)

	var b strings.Builder

	// Header line
	b.WriteString(borderStyle.Render(borderChar) + " " + titleStyle.Render(string(alert.Type)) + "\n")

	// Body lines
	bodyTrimmed := strings.TrimRight(renderedBody, "\n")
	bodyLines := strings.Split(bodyTrimmed, "\n")
	for _, line := range bodyLines {
		if strings.TrimSpace(line) == "" {
			b.WriteString(borderStyle.Render(borderChar) + "\n")
		} else {
			b.WriteString(borderStyle.Render(borderChar) + " " + line + "\n")
		}
	}

	return b.String()
}
