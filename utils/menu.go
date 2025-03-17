package utils

import (
	"github.com/fatih/color"
)

func Menu() {
	//hay que ameritar que debo crear y nombrar esta variable como: textoGigante
	//XD

	textoGigante := `
▒█▀▀█ ▒█░░░ ▀█▀ ░░ ▒█▀▀█ ▒█▀▀▀█
▒█░░░ ▒█░░░ ▒█░ ▀▀ ▒█░▄▄ ▒█░░▒█
▒█▄▄█ ▒█▄▄█ ▄█▄ ░░ ▒█▄▄█ ▒█▄▄▄█`

	color.Blue(textoGigante)

	color.Yellow("Bienvenido de nuevo, ¿que deseas hacer hoy?\n ")

	color.White("1. Gemini")
	color.White("2. Clima")
	color.White("3. Creditos")
	color.White("4. Version")
	color.White("5. Limpiar terminal")
	color.White("6. Letra de Cancion\n ")
}
