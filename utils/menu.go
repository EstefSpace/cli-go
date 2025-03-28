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

	color.White("gemini - Habla con la IA de Google")
	color.White("clima - Obten informacion del clima de una ciudad")
	color.White("spoo.me - Acorta enlaces y revisa estadisticas de ellos\n ")
}
