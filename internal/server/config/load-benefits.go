package config

import (
	"agency-site/internal/types"
	_ "embed"
	"encoding/json"
)

//go:embed data/benefits-data.json
var benefitsJSONBytes []byte

// LoadHeader now reads directly from the embedded memory slice
func LoadBenefits() ([]types.BenefitsData, error) {
	var benefitsData []types.BenefitsData

	err := json.Unmarshal(benefitsJSONBytes, &benefitsData)
	if err != nil {
		return []types.BenefitsData{}, err
	}

	return benefitsData, nil
}
