package style

import (
	"fmt"
	"os"
	"strings"

	"github.com/Summaw/aurora/pkg/color"
)

type DividerBuilder struct {
	text     string
	width    int
	char     string
	gradient color.Gradient
	output   *os.File
}

func NewDivider(text string) *DividerBuilder {
	return &DividerBuilder{
		text:     text,
		width:    60,
		char:     "─",
		gradient: color.GradientAurora,
		output:   os.Stdout,
	}
}

func (d *DividerBuilder) Width(w int) *DividerBuilder {
	d.width = w
	return d
}

func (d *DividerBuilder) Char(c string) *DividerBuilder {
	d.char = c
	return d
}

func (d *DividerBuilder) Gradient(name string) *DividerBuilder {
	d.gradient = color.GetGradient(name)
	return d
}

func (d *DividerBuilder) Build() string {
	if d.text == "" {
		line := strings.Repeat(d.char, d.width)
		return d.gradient.Apply(line) + "\n"
	}

	textLen := len([]rune(d.text)) + 2
	sideLen := (d.width - textLen) / 2
	if sideLen < 0 {
		sideLen = 0
	}

	left := strings.Repeat(d.char, sideLen)
	right := strings.Repeat(d.char, d.width-sideLen-textLen)

	fullLine := left + " " + d.text + " " + right
	return d.gradient.Apply(fullLine) + "\n"
}

func (d *DividerBuilder) Render() {
	fmt.Fprint(d.output, d.Build())
}

func (d *DividerBuilder) String() string {
	return d.Build()
}
