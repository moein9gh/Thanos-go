package validation

import "github.com/thanos-go/config"

type Validate struct {
	config config.Config
}

// creates a new validation service
func NewValidation(config config.Config) *Validate {
	return &Validate{config: config}
}
