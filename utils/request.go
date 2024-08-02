package ZeewUtils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Request[T any](url string, token string) (T, error) {
	var result T

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	// Establecer el encabezado de autorización
	req.Header.Set("Authorization", "token "+token)

	// Realizar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// Deserializar la respuesta en el tipo T
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
