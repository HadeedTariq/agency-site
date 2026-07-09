package config

import (
	"agency-site/internal/types"
	_ "embed"
	"encoding/json"
)

//go:embed data/carousel-data.json
var carouselJSONBytes []byte

// LoadHeader now reads directly from the embedded memory slice
func LoadCarousel() ([]types.HeroData, error) {
	var carouselData []types.HeroData

	err := json.Unmarshal(carouselJSONBytes, &carouselData)
	if err != nil {
		return []types.HeroData{}, err
	}

	return carouselData, nil
}
