package valueobjects

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	CEP        string             `bson:"cep" json:"cep"`
	Logradouro string             `bson:"logradouro" json:"logradouro"`
	Bairro     string             `bson:"bairro" json:"bairro"`
	Localidade string             `bson:"localidade" json:"localidade"`
	UF         string             `bson:"uf" json:"uf"`
	IBGE       string             `bson:"ibge" json:"ibge"`
	DDD        string             `bson:"ddd" json:"ddd"`
	Siafi      string             `bson:"siafi" json:"siafi"`
}

func NewAddress(CEP, Logradouro, Bairro, Localidade, UF, IBGE, DDD, Siafi string) (*Address, error) {
	address := &Address{
		ID:         primitive.NewObjectID(),
		CEP:        CEP,
		Logradouro: Logradouro,
		Bairro:     Bairro,
		Localidade: Localidade,
		UF:         UF,
		IBGE:       IBGE,
		DDD:        DDD,
		Siafi:      Siafi,
	}

	if err := address.Validate(); err != nil {
		return nil, err
	}

	return address, nil
}

func (a *Address) Validate() error {
	if !isValidCEP(a.CEP) {
		return errors.New("CEP inválido")
	}
	if a.Logradouro == "" {
		return errors.New("logradouro não pode ser vazio")
	}
	if a.Bairro == "" {
		return errors.New("bairro não pode ser vazio")
	}
	if a.Localidade == "" {
		return errors.New("localidade não pode ser vazia")
	}
	if !isValidUF(a.UF) {
		return errors.New("uf inválido")
	}
	if len(a.DDD) != 2 {
		return errors.New("ddd deve ter 2 dígitos")
	}

	return nil
}

func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return re.MatchString(cep)
}

func isValidUF(uf string) bool {
	validUFs := map[string]bool{
		"AC": true, "AL": true, "AP": true, "AM": true,
		"BA": true, "CE": true, "DF": true, "ES": true,
		"GO": true, "MA": true, "MT": true, "MS": true,
		"MG": true, "PA": true, "PB": true, "PR": true,
		"PE": true, "PI": true, "RJ": true, "RN": true,
		"RS": true, "RO": true, "RR": true, "SC": true,
		"SP": true, "SE": true, "TO": true,
	}

	_, exists := validUFs[uf]
	return exists
}
