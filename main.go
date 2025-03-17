package main

import (
	"cli-go/commands"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Red("Mal uso del CLI")
		color.Blue("Uso: cli-go <comando> [argumentos]")
		os.Exit(1)
	}

	comando := os.Args[1]

	switch comando {
	case "ayuda":
		commands.Ayuda()
	case "creditos":
		commands.Creditos()
	case "clima":
		if len(os.Args) < 3 {
			color.Red("Mal uso del comando")
			color.Blue("Uso: cli-go clima <ciudad>")
			os.Exit(1)
		}
		ciudad := strings.Join(os.Args[2:], " ")
		commands.Clima(ciudad)
	case "gemini":

		commands.Gemini()
	default:
		color.Red("Ese comando no existe")
		os.Exit(1)
	}
}
