package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type InfoLetra struct {
	title  string
	artist string
	lyrics string
}

func LetraDeCancion() {
	var cancion string

	fmt.Print("Escribe la canción: ")
	fmt.Scan(&cancion)

	url := "https://api.popcat.xyz/lyrics?song=" + cancion

	resp, err := http.Get(url)
	color.Blue("Obteniendo la letra...")
	if err != nil {
		color.Red("Hubo un error al intentar obtener información, puede ser que la API tenga problemas.")
		return
	}

	defer resp.Body.Close()

	var results []InfoLetra

	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		color.Red("Hubo un error al decodificar la respuesta, verifica que hayas escrito una cancion que exista.")
		return
	}

	if len(results) > 0 {

		color.Cyan("Al rato sigo chambieando")
	} else {
		color.Red("No se encontraron resultados para la cancion especificada.")
	}

}
