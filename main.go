package main

import (
	"cli-go/commands"
	"cli-go/utils"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {

	for {
		utils.Menu()

		var opcion string

		fmt.Print("Elige una opción (escribe salir para terminar la ejecucion): ")
		fmt.Scan(&opcion)

		switch opcion {
		case "salir":
			color.Yellow("Has elegido salir")
			return
		case "gemini":
			color.Yellow("Has elegido el comando de Gemini\n ")

			commands.Gemini()
			commands.ClearTerminal()
		case "clima":
			color.Yellow("Has elegido el comando de Clima\n ")

			commands.Clima()
			commands.ClearTerminal()

		case "spoo.me":
			color.Yellow("Has elegido el comando de Spoo.me\n ")

			commands.SpooMe()
			commands.ClearTerminal()
		default:
			color.Red("No se pudo encontrar ese comando, intentalo de nuevo.")

			time.Sleep(2 * time.Second) // Pausa de 2 segundos
			commands.ClearTerminal()
		}

	}
}
