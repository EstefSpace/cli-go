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
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
		case 2:
			color.Yellow("Has elegido el comando de Clima\n ")

			commands.Clima()
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
		case 3:
			color.Yellow("Has elegido el comando de Creditos\n ")

			commands.Creditos()
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
		case 4:
			color.Yellow("Has elegido el comando de Version\n ")

			commands.Version()
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
		case 5:
			color.Yellow("Has elegido el comando de Limpiar terminal\n ")

			commands.ClearTerminal()
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
		default:
			color.Red("No se pudo encontrar ese comando")
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
		}

	}
}
