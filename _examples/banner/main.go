package main

import "github.com/Summaw/aurora"

func main() {
	aurora.Banner("AURORA").
		Gradient("aurora").
		Tagline("Beautiful Console Logging").
		Version("v1.0.0").
		Render()

	aurora.Banner("CYBERPUNK").
		Gradient("cyberpunk").
		Border("double").
		Render()

	aurora.Banner("SUNSET").
		Gradient("sunset").
		Font("slant").
		Border("heavy").
		Render()

	aurora.Banner("MINIMAL").
		Gradient("neon").
		Font("minimal").
		Border("sharp").
		Render()

	aurora.Banner("CUSTOM").
		GradientMulti("#ff0000", "#ff7f00", "#ffff00", "#00ff00", "#0000ff").
		Tagline("Rainbow Gradient").
		Render()
}
