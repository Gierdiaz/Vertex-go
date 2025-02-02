package integrations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const viaCepURL = "https://viacep.com.br/ws/%s/json/"

type ViaCEP struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	IBGE       string `json:"ibge"`
	DDD        string `json:"ddd"`
	Siafi      string `json:"siafi"`
}

func GetAddressByCEP(cep string) (*ViaCEP, error) {
	url := fmt.Sprintf(viaCepURL, cep)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar endereço: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao buscar endereço: %s", resp.Status)
	}

	var address ViaCEP
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, fmt.Errorf("erro ao decodificar endereço: %v", err)
	}

	return &address, nil
}
