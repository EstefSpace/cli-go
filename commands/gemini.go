package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/estefspace/gemini-go"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func Gemini() {
	scanner := bufio.NewScanner(os.Stdin)

	err := godotenv.Load()
	if err != nil {
		color.Red("Error al tratar de cargar las variables de entorno (.env)")
		return
	}
	API_KEY := os.Getenv("API_KEY_GEMINI")

	client := gemini.NewClient(API_KEY)
	color.Yellow("Escribe salir para poder terminar la conversación")
	for {
		fmt.Print("Tu: ")
		scanner.Scan()
		prompt := scanner.Text()

		if err := scanner.Err(); err != nil {
			color.Red("Hubo un error al tratar de leer tu prompt")
			return
		}

		if strings.ToLower(prompt) == "salir" {
			color.Green("Conversación terminada.")
			break
		}

		resp, err := client.Ask("Necesito que respondas a lo siguiente con estas reglas: 1. NO USES ** ni # en tu texto 2. Responde de manera eficiente y rapida sin tanto rollo. Prompt:" + prompt)
		if err != nil {
			color.Red("Ocurrio un error al intentar mostrar la respuesta de Gemini: %v", err)
			continue // Continúa con el siguiente prompt
		}

		fmt.Println("Gemini: " + resp)
	}
}
