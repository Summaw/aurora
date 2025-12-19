package style

import (
	"fmt"
	"os"
	"strings"

	"github.com/Summaw/aurora/pkg/color"
)

type BoxBuilder struct {
	content  string
	border   string
	padding  int
	width    int
	title    string
	gradient color.Gradient
	output   *os.File
}

func NewBox(content string) *BoxBuilder {
	return &BoxBuilder{
		content:  content,
		border:   "rounded",
		padding:  1,
		gradient: color.GradientAurora,
		output:   os.Stdout,
	}
}

func (b *BoxBuilder) Border(style string) *BoxBuilder {
	b.border = style
	return b
}

func (b *BoxBuilder) Padding(p int) *BoxBuilder {
	b.padding = p
	return b
}

func (b *BoxBuilder) Width(w int) *BoxBuilder {
	b.width = w
	return b
}

func (b *BoxBuilder) Title(t string) *BoxBuilder {
	b.title = t
	return b
}

func (b *BoxBuilder) Gradient(name string) *BoxBuilder {
	b.gradient = color.GetGradient(name)
	return b
}

func (b *BoxBuilder) Build() string {
	lines := strings.Split(b.content, "\n")

	maxLen := 0
	for _, line := range lines {
		lineLen := len([]rune(line))
		if lineLen > maxLen {
			maxLen = lineLen
		}
	}
	titleLen := len([]rune(b.title))
	if b.title != "" && titleLen+4 > maxLen {
		maxLen = titleLen + 4
	}
	if b.width > maxLen {
		maxLen = b.width
	}

	maxLen += b.padding * 2

	chars := getBoxChars(b.border)
	var result strings.Builder

	if b.title != "" {
		titlePad := (maxLen - titleLen - 2) / 2
		result.WriteString(chars.topLeft)
		result.WriteString(strings.Repeat(chars.horizontal, titlePad))
		result.WriteString(" ")
		result.WriteString(b.gradient.Apply(b.title))
		result.WriteString(" ")
		result.WriteString(strings.Repeat(chars.horizontal, maxLen-titlePad-titleLen-2))
		result.WriteString(chars.topRight)
	} else {
		result.WriteString(chars.topLeft)
		result.WriteString(strings.Repeat(chars.horizontal, maxLen+2))
		result.WriteString(chars.topRight)
	}
	result.WriteString("\n")

	for i := 0; i < b.padding; i++ {
		result.WriteString(chars.vertical)
		result.WriteString(strings.Repeat(" ", maxLen+2))
		result.WriteString(chars.vertical)
		result.WriteString("\n")
	}

	for _, line := range lines {
		result.WriteString(chars.vertical)
		result.WriteString(" ")
		lineLen := len([]rune(line))
		pad := maxLen - lineLen
		leftPad := b.padding
		rightPad := pad - leftPad + b.padding
		result.WriteString(strings.Repeat(" ", leftPad))
		result.WriteString(line)
		result.WriteString(strings.Repeat(" ", rightPad))
		result.WriteString(" ")
		result.WriteString(chars.vertical)
		result.WriteString("\n")
	}

	for i := 0; i < b.padding; i++ {
		result.WriteString(chars.vertical)
		result.WriteString(strings.Repeat(" ", maxLen+2))
		result.WriteString(chars.vertical)
		result.WriteString("\n")
	}

	result.WriteString(chars.bottomLeft)
	result.WriteString(strings.Repeat(chars.horizontal, maxLen+2))
	result.WriteString(chars.bottomRight)
	result.WriteString("\n")

	return result.String()
}

func (b *BoxBuilder) Render() {
	fmt.Fprint(b.output, b.Build())
}

func (b *BoxBuilder) String() string {
	return b.Build()
}

type boxChars struct {
	topLeft     string
	topRight    string
	bottomLeft  string
	bottomRight string
	horizontal  string
	vertical    string
}

func getBoxChars(style string) boxChars {
	switch style {
	case "rounded":
		return boxChars{"╭", "╮", "╰", "╯", "─", "│"}
	case "sharp":
		return boxChars{"┌", "┐", "└", "┘", "─", "│"}
	case "double":
		return boxChars{"╔", "╗", "╚", "╝", "═", "║"}
	case "heavy":
		return boxChars{"┏", "┓", "┗", "┛", "━", "┃"}
	case "ascii":
		return boxChars{"+", "+", "+", "+", "-", "|"}
	default:
		return boxChars{"╭", "╮", "╰", "╯", "─", "│"}
	}
}
