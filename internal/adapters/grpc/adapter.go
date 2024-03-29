package adapters

import (
	"github.com/NBN23dev/gcr-go-template/internal/core/ports"
)

// GRPCAdapter
type GRPCAdapter struct {
	service ports.Service
}

// NewGRPCAdapter
func NewGRPCAdapter(service ports.Service) *GRPCAdapter {
	return &GRPCAdapter{service: service}
}
