package commands

import (
	"github.com/fatih/color"
)

func Ayuda() {
	//hay que ameritar que debo crear y nombrar esta variable como: textoGigante
	//XD
	textoGigante := `
▒█▀▀█ ▒█░░░ ▀█▀ ░░ ▒█▀▀█ ▒█▀▀▀█
▒█░░░ ▒█░░░ ▒█░ ▀▀ ▒█░▄▄ ▒█░░▒█
▒█▄▄█ ▒█▄▄█ ▄█▄ ░░ ▒█▄▄█ ▒█▄▄▄█`

	color.Blue(textoGigante)

	color.White("Descubre que comandos te ofrece este CLI\n ")

	color.Yellow("Categoria de Ayuda")
	color.Cyan("-> cli-go clima <ciudad>\n-> cli-go gemini")

	color.Yellow("Extras")
	color.Cyan("-> cli-go creditos")
}
