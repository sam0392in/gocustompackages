package headingstyler

import "github.com/mbndr/figlet4go"

func Styler(data, color, font *string) string {
	text := *data
	textFont := *font

	ascii := figlet4go.NewAsciiRender()
	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	if textFont == "3d" {
		options.FontName = "larry3d"
	}
	switch *color {
	case "Blue":
		options.FontColor = []figlet4go.Color{figlet4go.ColorBlue}
	case "Yellow":
		options.FontColor = []figlet4go.Color{figlet4go.ColorYellow}
	case "Green":
		options.FontColor = []figlet4go.Color{figlet4go.ColorGreen}
	case "Red":
		options.FontColor = []figlet4go.Color{figlet4go.ColorRed}
	case "Cyan":
		options.FontColor = []figlet4go.Color{figlet4go.ColorCyan}
	default:
		options.FontColor = []figlet4go.Color{figlet4go.ColorWhite}
	}

	renderStr, _ := ascii.RenderOpts(text, options)
	return renderStr
}
