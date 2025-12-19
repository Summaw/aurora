package color

type Gradient struct {
	Colors []RGB
}

func NewGradient(start, end RGB) Gradient {
	return Gradient{Colors: []RGB{start, end}}
}

func NewMultiGradient(hexColors ...string) Gradient {
	colors := make([]RGB, len(hexColors))
	for i, hex := range hexColors {
		colors[i] = FromHex(hex)
	}
	return Gradient{Colors: colors}
}

func (g Gradient) At(t float64) RGB {
	if len(g.Colors) == 0 {
		return White
	}
	if len(g.Colors) == 1 {
		return g.Colors[0]
	}
	if t <= 0 {
		return g.Colors[0]
	}
	if t >= 1 {
		return g.Colors[len(g.Colors)-1]
	}

	segments := float64(len(g.Colors) - 1)
	segment := int(t * segments)
	if segment >= len(g.Colors)-1 {
		segment = len(g.Colors) - 2
	}

	localT := (t*segments - float64(segment))
	return g.Colors[segment].Blend(g.Colors[segment+1], localT)
}

func (g Gradient) Apply(text string) string {
	if len(text) == 0 {
		return ""
	}

	runes := []rune(text)
	result := ""

	for i, r := range runes {
		var t float64
		if len(runes) > 1 {
			t = float64(i) / float64(len(runes)-1)
		} else {
			t = 0.5
		}
		c := g.At(t)
		result += c.ANSI() + string(r)
	}

	return result + Reset
}

func (g Gradient) ApplyLines(lines []string) []string {
	result := make([]string, len(lines))

	totalChars := 0
	for _, line := range lines {
		totalChars += len([]rune(line))
	}

	charIndex := 0
	for i, line := range lines {
		runes := []rune(line)
		coloredLine := ""

		for _, r := range runes {
			var t float64
			if totalChars > 1 {
				t = float64(charIndex) / float64(totalChars-1)
			} else {
				t = 0.5
			}
			c := g.At(t)
			coloredLine += c.ANSI() + string(r)
			charIndex++
		}

		result[i] = coloredLine + Reset
	}

	return result
}

func (g Gradient) ApplyVertical(lines []string) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		var t float64
		if len(lines) > 1 {
			t = float64(i) / float64(len(lines)-1)
		} else {
			t = 0.5
		}
		c := g.At(t)
		result[i] = c.ANSI() + line + Reset
	}

	return result
}

func (g Gradient) ApplyDiagonal(lines []string) []string {
	result := make([]string, len(lines))

	maxLen := 0
	for _, line := range lines {
		if len([]rune(line)) > maxLen {
			maxLen = len([]rune(line))
		}
	}

	for i, line := range lines {
		runes := []rune(line)
		coloredLine := ""

		for j, r := range runes {
			var diagonal float64
			if len(lines)+maxLen > 2 {
				diagonal = float64(i+j) / float64(len(lines)+maxLen-2)
			} else {
				diagonal = 0.5
			}
			c := g.At(diagonal)
			coloredLine += c.ANSI() + string(r)
		}

		result[i] = coloredLine + Reset
	}

	return result
}
