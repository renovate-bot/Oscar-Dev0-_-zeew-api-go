package zeewImg

import (
	"encoding/json"
	"fmt"
	"net/url"
	ZeewUtils "github.com/Oscar-Dev0/zeew-api-go/utils"
)

type Card struct {
	Token      string
	Bienvenida *Bienvenida
}

func NewCards(token string) Card {
	return Card{
		Token: token,
		Bienvenida: &Bienvenida{
			Token: token,
		},
	}
}

func (c Card) Render(bn Bienvenida) (*byte, *ZeewUtils.ZeewError) {
	if bn.Token == "" {
		return nil, ZeewUtils.NewZeewError(ZeewUtils.TokenEmpty)
	};

	if(bn.Estilo == "") {
		return nil, ZeewUtils.NewZeewError(ZeewUtils.EstiloEmpty)
	};

	jsonData, err := json.Marshal(bn)
    if err != nil {
		return nil, ZeewUtils.NewZeewError(ZeewUtils.ErrorMarshallingJson)
    }

    // Convertir JSON a un mapa para crear parámetros de la URL
    var jsonMap map[string]interface{}
    if err := json.Unmarshal(jsonData, &jsonMap); err != nil {
		return nil, ZeewUtils.NewZeewError(ZeewUtils.ErrorUnmarshallingJson)
    }

    // Crear parámetros de la URL
    query := url.Values{}
    for key, value := range jsonMap {
        query.Add(key, fmt.Sprint(value))
    }


    requestURL := fmt.Sprintf("%s/bw/%s?%s", ZeewUtils.INT, bn.Estilo, query.Encode())

	var r,e  = ZeewUtils.Request[byte](requestURL, bn.Token);
	if e != nil {
		return nil, ZeewUtils.NewZeewError(ZeewUtils.ErrorRequestingImage)
	}
	return &r, nil
}