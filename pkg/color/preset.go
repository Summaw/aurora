package color

var (
	GradientSunset = Gradient{
		Colors: []RGB{{255, 107, 107}, {254, 202, 87}, {255, 159, 243}},
	}

	GradientOcean = Gradient{
		Colors: []RGB{{0, 82, 212}, {67, 100, 247}, {111, 177, 252}},
	}

	GradientNeon = Gradient{
		Colors: []RGB{{255, 0, 255}, {0, 255, 255}},
	}

	GradientCyberpunk = Gradient{
		Colors: []RGB{{247, 37, 133}, {114, 9, 183}, {58, 12, 163}},
	}

	GradientMiami = Gradient{
		Colors: []RGB{{247, 37, 133}, {76, 201, 240}},
	}

	GradientFire = Gradient{
		Colors: []RGB{{241, 39, 17}, {245, 175, 25}},
	}

	GradientForest = Gradient{
		Colors: []RGB{{19, 78, 94}, {113, 178, 128}},
	}

	GradientGalaxy = Gradient{
		Colors: []RGB{{127, 0, 255}, {225, 0, 255}},
	}

	GradientRetro = Gradient{
		Colors: []RGB{{252, 70, 107}, {63, 94, 251}},
	}

	GradientAurora = Gradient{
		Colors: []RGB{{0, 198, 255}, {0, 114, 255}, {114, 9, 183}, {247, 37, 133}},
	}

	GradientMint = Gradient{
		Colors: []RGB{{0, 176, 155}, {150, 201, 61}},
	}

	GradientPeach = Gradient{
		Colors: []RGB{{255, 154, 158}, {250, 208, 196}},
	}

	GradientLavender = Gradient{
		Colors: []RGB{{150, 131, 236}, {246, 191, 255}},
	}

	GradientGold = Gradient{
		Colors: []RGB{{255, 215, 0}, {255, 165, 0}, {184, 134, 11}},
	}

	GradientIce = Gradient{
		Colors: []RGB{{230, 240, 255}, {135, 206, 250}, {70, 130, 180}},
	}

	GradientBlood = Gradient{
		Colors: []RGB{{139, 0, 0}, {220, 20, 60}, {178, 34, 34}},
	}

	GradientMatrix = Gradient{
		Colors: []RGB{{0, 50, 0}, {0, 255, 65}, {0, 100, 0}},
	}

	GradientVaporwave = Gradient{
		Colors: []RGB{{255, 113, 206}, {1, 205, 254}, {185, 103, 255}},
	}

	GradientRainbow = Gradient{
		Colors: []RGB{
			{255, 0, 0}, {255, 127, 0}, {255, 255, 0},
			{0, 255, 0}, {0, 0, 255}, {75, 0, 130}, {148, 0, 211},
		},
	}

	GradientTerminal = Gradient{
		Colors: []RGB{{0, 255, 0}, {0, 200, 0}},
	}

	GradientRose = Gradient{
		Colors: []RGB{{255, 0, 128}, {255, 102, 178}, {255, 179, 217}},
	}

	GradientSky = Gradient{
		Colors: []RGB{{135, 206, 235}, {70, 130, 180}, {25, 25, 112}},
	}
)

var gradientMap = map[string]Gradient{
	"sunset":    GradientSunset,
	"ocean":     GradientOcean,
	"neon":      GradientNeon,
	"cyberpunk": GradientCyberpunk,
	"miami":     GradientMiami,
	"fire":      GradientFire,
	"forest":    GradientForest,
	"galaxy":    GradientGalaxy,
	"retro":     GradientRetro,
	"aurora":    GradientAurora,
	"mint":      GradientMint,
	"peach":     GradientPeach,
	"lavender":  GradientLavender,
	"gold":      GradientGold,
	"ice":       GradientIce,
	"blood":     GradientBlood,
	"matrix":    GradientMatrix,
	"vaporwave": GradientVaporwave,
	"rainbow":   GradientRainbow,
	"terminal":  GradientTerminal,
	"rose":      GradientRose,
	"sky":       GradientSky,
}

func GetGradient(name string) Gradient {
	if g, ok := gradientMap[name]; ok {
		return g
	}
	return GradientAurora
}

func ListGradients() []string {
	names := make([]string, 0, len(gradientMap))
	for name := range gradientMap {
		names = append(names, name)
	}
	return names
}
