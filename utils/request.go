package ZeewUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Status   string `json:"status"`
	Message  string `json:"mensaje"`
	Endpoint string `json:"endpoint"`
}

func Request(url string, token string) (*bytes.Buffer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", INT, url), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err == nil && res.Message != "" {
		return nil, fmt.Errorf("error en la solicitud: %s", res.Message)
	}

	return bytes.NewBuffer(body), nil;
}
