package style

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Summaw/aurora/pkg/color"
)

type KVBuilder struct {
	pairs    map[string]any
	gradient color.Gradient
	output   *os.File
}

func NewKV(pairs map[string]any) *KVBuilder {
	return &KVBuilder{
		pairs:    pairs,
		gradient: color.GradientAurora,
		output:   os.Stdout,
	}
}

func (k *KVBuilder) Gradient(name string) *KVBuilder {
	k.gradient = color.GetGradient(name)
	return k
}

func (k *KVBuilder) Build() string {
	maxKeyLen := 0
	keys := make([]string, 0, len(k.pairs))
	for key := range k.pairs {
		keys = append(keys, key)
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
	}
	sort.Strings(keys)

	var result strings.Builder
	result.WriteString("\n")

	for _, key := range keys {
		value := k.pairs[key]
		coloredKey := k.gradient.Apply(key)
		padding := strings.Repeat(" ", maxKeyLen-len(key))
		result.WriteString(fmt.Sprintf("  %s%s   %v\n", coloredKey, padding, value))
	}

	result.WriteString("\n")
	return result.String()
}

func (k *KVBuilder) Render() {
	fmt.Fprint(k.output, k.Build())
}

func (k *KVBuilder) String() string {
	return k.Build()
}
