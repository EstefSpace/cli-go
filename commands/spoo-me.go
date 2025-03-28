package commands

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
)

func SpooMe() {
	var herramienta string

	color.Blue("acortar - Acorta un enlace para obtener por ejemplo: https://spoo.me/tuEnlaceAcortado")
	color.Blue("estadisticas - Revisa estadisticas del enlace nuevo (https://spoo.me/tuEnlaceAcortado)")

	fmt.Print("Elige una opcion (escribe cualquier otra cosa para regresar al menu): ")
	fmt.Scan(&herramienta)

	switch herramienta {
	case "acortar":
		var urlLargo string
		var alias string
		fmt.Print(" \nEscribe el enlace para acortar: ")
		fmt.Scan(&urlLargo)

		if !strings.HasPrefix(urlLargo, "https://") {
			color.Red("Ese enlace no es valido porque no contiene https://")
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
			return
		}

		fmt.Print("Escribe el alias de tu nuevo enlace (por ejemplo tuEnlaceAcortado): ")
		fmt.Scan(&alias)

		// Set DNS resolver to use Google's public DNS server
		net.DefaultResolver = &net.Resolver{
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{Timeout: 500 * time.Millisecond}
				return d.DialContext(ctx, "udp", "8.8.8.8:53")
			},
		}

		apiUrl := "https://spoo.me/" // muchas gracias por existir :heart:

		encodedLongUrl := url.QueryEscape(urlLargo)
		payload := strings.NewReader("url=" + encodedLongUrl + "&alias=" + alias)

		req, err := http.NewRequest("POST", apiUrl, payload)
		if err != nil {
			color.Red("Error al crear la peticion")
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
			return
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Accept", "application/json")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			color.Red("Hubo un error inesperado")
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
			return
		}

		defer res.Body.Close()

		var opcion string
		color.Green("¡Tu enlace acortado esta listo! Aparece como: https://spoo.me/" + alias)
		fmt.Print("Si deseas regresar al menu principal escribe salir: ")
		fmt.Scan(&opcion)

		if opcion == "salir" {
			color.Yellow("Regresando al menu...")
			time.Sleep(2 * time.Second) // Pausa de 2 segundos
		} else { // ya se que me puedo ahorrar esto de abajo pero pues para mi queda chistoso nose
			color.Yellow("¡Parece que usted no sabe seguir ordenes! pero aun asi regresara al menu...")
			time.Sleep(3 * time.Second) // Pausa de 3 segundos
		}
	case "estadisticas":
		fmt.Print("Proximamente...")
		time.Sleep(2 * time.Second) // Pausa de 2 segundos
	default:
		color.Red(" \nEsa opcion no existe, intentalo de nuevo.")
		time.Sleep(2 * time.Second) // Pausa de 2 segundos
	}
}
