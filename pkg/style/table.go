package style

import (
	"fmt"
	"os"
	"strings"

	"github.com/Summaw/aurora/pkg/color"
)

func displayWidth(s string) int {
	width := 0
	for _, r := range s {
		if r >= 0x1100 &&
			(r <= 0x115F ||
				r == 0x2329 || r == 0x232A ||
				(r >= 0x2E80 && r <= 0xA4CF && r != 0x303F) ||
				(r >= 0xAC00 && r <= 0xD7A3) ||
				(r >= 0xF900 && r <= 0xFAFF) ||
				(r >= 0xFE10 && r <= 0xFE19) ||
				(r >= 0xFE30 && r <= 0xFE6F) ||
				(r >= 0xFF00 && r <= 0xFF60) ||
				(r >= 0xFFE0 && r <= 0xFFE6) ||
				(r >= 0x20000 && r <= 0x2FFFD) ||
				(r >= 0x30000 && r <= 0x3FFFD)) {
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

type TableBuilder struct {
	headers  []string
	rows     [][]string
	border   string
	gradient color.Gradient
	output   *os.File
}

func NewTable(headers []string, rows [][]string) *TableBuilder {
	return &TableBuilder{
		headers:  headers,
		rows:     rows,
		border:   "rounded",
		gradient: color.GradientAurora,
		output:   os.Stdout,
	}
}

func (t *TableBuilder) Border(style string) *TableBuilder {
	t.border = style
	return t
}

func (t *TableBuilder) Gradient(name string) *TableBuilder {
	t.gradient = color.GetGradient(name)
	return t
}

func (t *TableBuilder) Build() string {
	colWidths := make([]int, len(t.headers))
	for i, h := range t.headers {
		colWidths[i] = displayWidth(h)
	}
	for _, row := range t.rows {
		for i, cell := range row {
			w := displayWidth(cell)
			if i < len(colWidths) && w > colWidths[i] {
				colWidths[i] = w
			}
		}
	}

	chars := getTableChars(t.border)
	var result strings.Builder

	result.WriteString(chars.topLeft)
	for i, w := range colWidths {
		result.WriteString(strings.Repeat(chars.horizontal, w+2))
		if i < len(colWidths)-1 {
			result.WriteString(chars.topMid)
		}
	}
	result.WriteString(chars.topRight)
	result.WriteString("\n")

	result.WriteString(chars.vertical)
	for i, h := range t.headers {
		result.WriteString(" ")
		colored := t.gradient.Apply(h)
		result.WriteString(colored)
		padding := colWidths[i] - displayWidth(h) + 1
		result.WriteString(strings.Repeat(" ", padding))
		result.WriteString(chars.vertical)
	}
	result.WriteString("\n")

	result.WriteString(chars.midLeft)
	for i, w := range colWidths {
		result.WriteString(strings.Repeat(chars.horizontal, w+2))
		if i < len(colWidths)-1 {
			result.WriteString(chars.midMid)
		}
	}
	result.WriteString(chars.midRight)
	result.WriteString("\n")

	for _, row := range t.rows {
		result.WriteString(chars.vertical)
		for i := range colWidths {
			cell := ""
			if i < len(row) {
				cell = row[i]
			}
			result.WriteString(" ")
			result.WriteString(cell)
			padding := colWidths[i] - displayWidth(cell) + 1
			result.WriteString(strings.Repeat(" ", padding))
			result.WriteString(chars.vertical)
		}
		result.WriteString("\n")
	}

	result.WriteString(chars.bottomLeft)
	for i, w := range colWidths {
		result.WriteString(strings.Repeat(chars.horizontal, w+2))
		if i < len(colWidths)-1 {
			result.WriteString(chars.bottomMid)
		}
	}
	result.WriteString(chars.bottomRight)
	result.WriteString("\n")

	return result.String()
}

func (t *TableBuilder) Render() {
	fmt.Fprint(t.output, t.Build())
}

func (t *TableBuilder) String() string {
	return t.Build()
}

type tableChars struct {
	topLeft     string
	topRight    string
	topMid      string
	bottomLeft  string
	bottomRight string
	bottomMid   string
	midLeft     string
	midRight    string
	midMid      string
	horizontal  string
	vertical    string
}

func getTableChars(style string) tableChars {
	switch style {
	case "rounded":
		return tableChars{"╭", "╮", "┬", "╰", "╯", "┴", "├", "┤", "┼", "─", "│"}
	case "sharp":
		return tableChars{"┌", "┐", "┬", "└", "┘", "┴", "├", "┤", "┼", "─", "│"}
	case "double":
		return tableChars{"╔", "╗", "╦", "╚", "╝", "╩", "╠", "╣", "╬", "═", "║"}
	case "heavy":
		return tableChars{"┏", "┓", "┳", "┗", "┛", "┻", "┣", "┫", "╋", "━", "┃"}
	default:
		return tableChars{"╭", "╮", "┬", "╰", "╯", "┴", "├", "┤", "┼", "─", "│"}
	}
}
