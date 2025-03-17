package commands

import "github.com/fatih/color"

func Version() {

	var version string = "v1.0.0"
	color.Blue("Version: " + version)
}
