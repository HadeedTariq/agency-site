package config

import (
	"agency-site/internal/types"
	_ "embed"
	"encoding/json"
)

//go:embed data/advantage-data.json
var advantagesJSONBytes []byte

// LoadHeader now reads directly from the embedded memory slice
func LoadAdvantages() ([]types.AdvantageData, error) {
	var advantageData []types.AdvantageData

	err := json.Unmarshal(advantagesJSONBytes, &advantageData)
	if err != nil {
		return []types.AdvantageData{}, err
	}

	return advantageData, nil
}
