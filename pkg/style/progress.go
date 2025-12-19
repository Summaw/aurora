package style

import (
	"fmt"
	"os"
	"strings"

	"github.com/Summaw/aurora/pkg/color"
)

type ProgressBar struct {
	label    string
	total    int
	current  int
	width    int
	gradient color.Gradient
	output   *os.File
	complete string
	pending  string
}

func NewProgressBar(label string, total int) *ProgressBar {
	return &ProgressBar{
		label:    label,
		total:    total,
		current:  0,
		width:    40,
		gradient: color.GradientAurora,
		output:   os.Stdout,
		complete: "█",
		pending:  "░",
	}
}

func (p *ProgressBar) Width(w int) *ProgressBar {
	p.width = w
	return p
}

func (p *ProgressBar) Gradient(name string) *ProgressBar {
	p.gradient = color.GetGradient(name)
	return p
}

func (p *ProgressBar) Chars(complete, pending string) *ProgressBar {
	p.complete = complete
	p.pending = pending
	return p
}

func (p *ProgressBar) Set(value int) {
	p.current = value
	p.render()
}

func (p *ProgressBar) Increment() {
	p.current++
	p.render()
}

func (p *ProgressBar) render() {
	percent := float64(p.current) / float64(p.total)
	if percent > 1 {
		percent = 1
	}

	filled := int(percent * float64(p.width))
	empty := p.width - filled

	bar := strings.Repeat(p.complete, filled) + strings.Repeat(p.pending, empty)
	coloredBar := p.gradient.Apply(bar)

	percentStr := fmt.Sprintf("%3.0f%%", percent*100)

	fmt.Fprintf(p.output, "\r  %s %s %s", p.label, coloredBar, percentStr)
}

func (p *ProgressBar) Done() {
	p.current = p.total
	p.render()
	fmt.Fprintln(p.output)
}

func (p *ProgressBar) Clear() {
	fmt.Fprint(p.output, "\r\033[K")
}
