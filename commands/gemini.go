package commands

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type InfoGemini struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func Gemini() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("¡Bienvenido de nuevo! ¿Qué deseas preguntar hoy?: ")
	scanner.Scan()
	prompt := scanner.Text()

	if err := scanner.Err(); err != nil {
		color.Red("Hubo un error al tratar de leer tu prompt")
		return
	}

	err := godotenv.Load()
	if err != nil {
		color.Red("Error al tratar de cargar las variables de entorno (.env)")
	}
	API_KEY := os.Getenv("API_KEY_GEMINI")
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=" + API_KEY

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": prompt,
					},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		color.Red("Error al intentar ejecutar el comando.") // Usamos color.Red()
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		color.Red("Error al crear la petición a la API de Gemini") // Usamos color.Red()
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		color.Red("Error al realizar la petición a la API de Gemini") // Usamos color.Red()
		os.Exit(1)
	}
	defer resp.Body.Close()

	var info InfoGemini

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Fatalf("Error al decodificar JSON: %v", err)
	}

	// Extraer y mostrar el texto
	if len(info.Candidates) > 0 && len(info.Candidates[0].Content.Parts) > 0 {
		texto := info.Candidates[0].Content.Parts[0].Text
		color.Blue("Esperando respuesta de Gemini...")
		fmt.Println(texto)
	} else {
		color.Red("No se pudo mostrar el texto generado por Gemini.")
	}
}
