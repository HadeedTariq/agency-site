package config

import (
	"agency-site/internal/types"
	_ "embed"
	"encoding/json"
)

//go:embed data/testimonial-data.json
var testimonialJSONBytes []byte

// LoadHeader now reads directly from the embedded memory slice
func LoadTestimonial() ([]types.TestimonialData, error) {
	var testimonialData []types.TestimonialData

	err := json.Unmarshal(testimonialJSONBytes, &testimonialData)
	if err != nil {
		return []types.TestimonialData{}, err
	}

	return testimonialData, nil
}
