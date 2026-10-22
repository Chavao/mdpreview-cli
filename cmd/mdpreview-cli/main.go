package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Chavao/mdpreview-cli/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const version = "1.0.0"

func main() {
	isStdinTerm := term.IsTerminal(int(os.Stdin.Fd()))
	code := run(os.Args[1:], os.Stdin, isStdinTerm, os.Stdout, os.Stderr)
	os.Exit(code)
}

func printHelp(w io.Writer) {
	helpText := fmt.Sprintf(`mdpreview-cli v%s - Terminal Markdown Previewer

Usage:
  mdpreview-cli <file.md>         Preview markdown file
  cat <file.md> | mdpreview-cli   Preview markdown piped from stdin

Options:
  -h, --help      Show this help message
  -v, --version   Show application version

Controls:
  ↑/↓, j/k        Scroll up and down
  Mouse Wheel     Scroll with mouse
  PageUp/PageDown Scroll by full page
  Home/End, g/G   Jump to top / bottom
  Ctrl+U/Ctrl+D   Scroll by half page
  Ctrl+Q, q       Open exit confirmation dialog
  Enter           Confirm button in dialog
  Esc             Cancel dialog and return to preview
`, version)
	fmt.Fprint(w, helpText)
}

func run(args []string, stdin io.Reader, isStdinTerm bool, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("mdpreview-cli", flag.ContinueOnError)
	flags.SetOutput(stderr)

	showHelp := flags.Bool("help", false, "Show help")
	flags.BoolVar(showHelp, "h", false, "Show help (shorthand)")

	showVersion := flags.Bool("version", false, "Show version")
	flags.BoolVar(showVersion, "v", false, "Show version (shorthand)")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	if *showHelp {
		printHelp(stdout)
		return 0
	}

	if *showVersion {
		fmt.Fprintf(stdout, "mdpreview-cli version %s\n", version)
		return 0
	}

	var content string
	var title string
	var customTty *os.File

	if !isStdinTerm {
		// Read from piped stdin
		bytes, err := io.ReadAll(stdin)
		if err != nil {
			fmt.Fprintf(stderr, "Error reading from stdin: %v\n", err)
			return 1
		}
		content = string(bytes)
		title = "[stdin]"

		// Open /dev/tty so Bubbletea can receive user input from the terminal
		tty, err := os.Open("/dev/tty")
		if err != nil {
			fmt.Fprintf(stderr, "Error opening /dev/tty: %v\n", err)
			return 1
		}
		customTty = tty
		defer customTty.Close()
	} else {
		// No stdin pipe; require file argument
		restArgs := flags.Args()
		if len(restArgs) == 0 {
			printHelp(stdout)
			return 0
		}

		filePath := restArgs[0]
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Fprintf(stderr, "Error: file not found: %s\n", filePath)
			} else {
				fmt.Fprintf(stderr, "Error reading file %s: %v\n", filePath, err)
			}
			return 1
		}
		content = string(bytes)
		title = filepath.Base(filePath)
	}

	isDark := lipgloss.HasDarkBackground()
	model := ui.NewModel(content, title, isDark)

	teaOptions := []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}

	if customTty != nil {
		teaOptions = append(teaOptions, tea.WithInput(customTty))
	}

	p := tea.NewProgram(model, teaOptions...)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(stderr, "Error running TUI: %v\n", err)
		return 1
	}

	return 0
}
