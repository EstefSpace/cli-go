package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type InfoClima struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Temperature string `json:"temperature"`
		Skytext     string `json:"skytext"`
	} `json:"current"`
}

func Clima() {

	var ciudad string
	fmt.Print("Introduce la ciudad a revisar: ")

	fmt.Scan(&ciudad)
	url := "https://api.popcat.xyz/weather?q=" + ciudad

	resp, err := http.Get(url)
	color.Blue("Obteniendo el clima...")
	if err != nil {
		color.Red("Hubo un error al intentar obtener información, puede ser que la API tenga problemas.")
		time.Sleep(2 * time.Second) // Pausa de 2 segundos
		return
	}

	defer resp.Body.Close()

	var results []InfoClima

	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		color.Red("Hubo un error al decodificar la respuesta, verifica que hayas escrito una ciudad que exista.")
		time.Sleep(2 * time.Second) // Pausa de 2 segundos
		return
	}

	if len(results) > 0 {
		info := results[0]
		color.Cyan("Clima en %s: %s°C, %s\n", info.Location.Name, info.Current.Temperature, info.Current.Skytext)
		time.Sleep(2 * time.Second) // Pausa de 2 segundos
	} else {
		color.Red("No se encontraron resultados para la ciudad especificada.")
		time.Sleep(2 * time.Second) // Pausa de 2 segundos
	}

}
