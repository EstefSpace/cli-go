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

	color.White("1. Gemini - Habla con la IA de Google")
	color.White("2. Clima - Obten informacion del clima de una ciudad")
	color.White("3. Creditos - Obten información de quien creo este CLI")
	color.White("4. Version - Revisa en que version tiene tu CLI")
	color.White("5. Limpiar terminal - Limpia tu terminal\n ")
}
