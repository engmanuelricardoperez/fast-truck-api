package providers

import (
	"context"

	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/entities"
)

type TruckProvider interface {
	GetTruck(ctx context.Context) (*entities.Truck, error)
}
