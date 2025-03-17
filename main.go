package main

import (
	"cli-go/commands"
	"cli-go/utils"
	"fmt"

	"github.com/fatih/color"
)

func main() {

	for {
		utils.Menu()

		var opcion int

		fmt.Print("Elige una opción (escribe 10 para salir): ")
		fmt.Scan(&opcion)

		switch opcion {
		case 10:
			color.Yellow("Has elegido salir")
			return
		case 1:
			color.Yellow("Has elegido el comando de Gemini\n ")

			commands.Gemini()
		case 2:
			color.Yellow("Has elegido el comando de Clima\n ")

			commands.Clima()
		case 3:
			color.Yellow("Has elegido el comando de Creditos\n ")

			commands.Creditos()
		case 4:
			color.Yellow("Has elegido el comando de Version\n ")

			commands.Version()
		case 5:
			color.Yellow("Has elegido el comando de Limpiar terminal\n ")

			commands.ClearTerminal()
		default:
			color.Red("No se pudo encontrar ese comando")
		}

	}
}
