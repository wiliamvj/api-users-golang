package viacep

import (
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/wiliamvj/api-users-golang/config/env"
)

type ViaCepResponse struct {
  CEP         string `json:"cep"`
  Logradouro  string `json:"logradouro"`
  Complemento string `json:"complemento"`
  Bairro      string `json:"bairro"`
  Localidade  string `json:"localidade"`
  UF          string `json:"uf"`
  IBGE        string `json:"ibge"`
  GIA         string `json:"gia"`
  DDD         string `json:"ddd"`
  SIAFI       string `json:"siafi"`
}

func GetCep(cep string) (*ViaCepResponse, error) {
  url := fmt.Sprintf("%s/%s/json", env.Env.ViaCepURL, cep)
  var viaCepResponse ViaCepResponse

  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  err = json.NewDecoder(resp.Body).Decode(&viaCepResponse)
  if err != nil {
    return nil, err
  }
  if viaCepResponse.CEP == "" {
    return nil, fmt.Errorf("cep not found")
  }
  return &viaCepResponse, nil
}
