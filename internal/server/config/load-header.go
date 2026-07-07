package config

import (
	"agency-site/internal/types"
	_ "embed"
	"encoding/json"
)

//go:embed data/header-data.json
var headerJSONBytes []byte

// LoadHeader now reads directly from the embedded memory slice
func LoadHeader() (types.HeaderConfig, error) {
	var headerData types.HeaderConfig

	err := json.Unmarshal(headerJSONBytes, &headerData)
	if err != nil {
		return types.HeaderConfig{}, err
	}

	return headerData, nil
}
