# Markdown Showcase

This document demonstrates all supported standard Markdown features and extensions in `mdpreview-cli`.

---

## 1. Headings (H1 to H6)

# Heading 1
## Heading 2
### Heading 3
#### Heading 4
##### Heading 5
###### Heading 6

---

## 2. Text Formatting & Inline Elements

### Emphasis
- Italic with asterisks: *italic text*
- Italic with underscores: _italic text_
- Bold with asterisks: **bold text**
- Bold with underscores: __bold text__
- Bold and italic with asterisks: ***bold and italic text***
- Bold and italic with underscores: ___bold and italic text___
- Strikethrough: ~~strikethrough text~~

### Inline Code
Use the `fmt.Println()` function to print text to standard output in Go.

### Links & Autolinks
- Link with title: [Go Dev](https://go.dev "The Go Programming Language")
- Delimited autolink: <https://github.com>
- Plain URL autolink: https://charm.sh

### Images
![Sample Preview Image](https://images.unsplash.com/photo-1579783900882-c0d3dad7b119?w=400 "Sample Image")

---

## 3. Paragraphs & Line Breaks

This is a standard paragraph demonstrating text flow across lines with dynamic word wrapping based on your terminal width.

This line ends with two spaces to create a line break:  
Here is the broken line.  
Alternatively, you can use the break tag: Line one<br>Line two.

---

## 4. Blockquotes

> This is a single-level blockquote spanning multiple lines.
> Knowledge increases by sharing, not by saving.
>> This is a nested blockquote providing further detail.

---

## 5. Lists

### Unordered Lists
- Unordered item with hyphen
  - Sub-item level 2 (2 spaces indented)
    - Sub-item level 3 (4 spaces indented)
* Unordered item with asterisk
+ Unordered item with plus sign

### Ordered Lists
1. First step
2. Second step
   1. Sub-step 2.1
   2. Sub-step 2.2
3. Third step

### Task Lists
- [x] Design TUI architecture with Bubble Tea
- [x] Configure Lipgloss themes and palette
- [x] Implement GitHub Alerts renderer
- [ ] Add custom user stylesheets
- [ ] Support custom keybindings config

---

## 6. Fenced Code Blocks

### Go
```go
package main

import "fmt"

type Server struct {
	Host string
	Port int
}

func (s *Server) Start() error {
	fmt.Printf("Server listening on %s:%d\n", s.Host, s.Port)
	return nil
}
```

### TypeScript
```typescript
interface UserProfile {
  id: string;
  username: string;
  active: boolean;
}

export function greet(user: UserProfile): string {
  return `Hello, ${user.username}!`;
}
```

### Bash
```bash
# Preview via file argument
./bin/mdpreview-cli examples/showcase.md

# Or pipe via stdin
cat examples/showcase.md | ./bin/mdpreview-cli
```

---

## 7. Tables with Column Alignment

| Feature | Type | Status | Latency |
| :--- | :---: | :---: | ---: |
| GitHub Alerts | Custom Lipgloss Box | Supported | < 1ms |
| Syntax Highlighting | Chroma ANSI Engine | Supported | 2ms |
| Table Layout | Styled Grid | Supported | < 1ms |
| Vertical Scrollbar | Viewport Track | Supported | Realtime |

---

## 8. GitHub Alerts & Callouts

> [!NOTE]
> Highlights useful information that users should know, even when skimming.

> [!TIP]
> Helpful advice for doing things better or more easily.

> [!IMPORTANT]
> Key information users need to know to achieve their goal.

> [!WARNING]
> Urgent info that needs immediate user attention to avoid problems.

> [!CAUTION]
> Advises about risks or negative outcomes of certain actions.
