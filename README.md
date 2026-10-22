# mdpreview-cli

Terminal Markdown viewer and reader built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Lipgloss](https://github.com/charmbracelet/lipgloss), and [Glamour](https://github.com/charmbracelet/glamour).

Features visual styling inspired by `mdpreview`, including GitHub Alerts (`[!NOTE]`, `[!TIP]`, `[!IMPORTANT]`, `[!WARNING]`, `[!CAUTION]`), syntax highlighting, table borders, adaptive terminal theming, and an interactive exit dialog.

## Features

- Full-screen terminal user interface (TUI)
- Reads Markdown files directly or piped via standard input (stdin)
- Mouse wheel and keyboard scrolling support (arrows, Vim keys `j`/`k`, `PageUp`/`PageDown`, `Home`/`End`)
- Vertical progress scrollbar on the right edge
- Informative status bar displaying document title, scroll percentage, line count, and shortcuts
- Exit confirmation dialog triggered via `Ctrl+Q` with interactive `[ No ]` and `[ Yes ]` options (defaults to `[ Yes ]`)
- Automatic dark/light background detection based on the user's terminal profile
- Custom GitHub alert rendering matching `mdpreview` color specifications

## Installation & Build

```bash
go build -o bin/mdpreview-cli ./cmd/mdpreview-cli
```

Optionally install into your `$GOPATH/bin`:

```bash
go install ./cmd/mdpreview-cli
```

## Usage

### 1. Previewing a file

```bash
mdpreview-cli README.md
```

### 2. Piping from stdin

```bash
cat file.md | mdpreview-cli
```

Or piping output from any command:

```bash
curl -s https://raw.githubusercontent.com/charmbracelet/bubbletea/master/README.md | mdpreview-cli
```

### 3. Displaying help and version

```bash
mdpreview-cli --help
mdpreview-cli --version
```

## Keyboard and Mouse Controls

| Action | Shortcuts |
| --- | --- |
| Scroll Down | `↓`, `j`, Mouse Wheel Down |
| Scroll Up | `↑`, `k`, Mouse Wheel Up |
| Page Down | `PageDown`, `f`, `Ctrl+D` (half page) |
| Page Up | `PageUp`, `b`, `Ctrl+U` (half page) |
| Jump to Top | `Home`, `g` |
| Jump to Bottom | `End`, `G` |
| Exit Dialog | `Ctrl+Q`, `q` |
| Dialog Navigation | `←` / `→`, `Tab` / `Shift+Tab`, `h` / `l` |
| Confirm Dialog | `Enter` (executes selected button), `y` |
| Cancel Dialog | `Esc`, `n` |
| Force Quit | `Ctrl+C` |

## Running Tests

```bash
go test ./...
```
