package truck

import (
	"context"

	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/entities"
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/providers"
)

type Get interface {
	Execute(ctx context.Context) (*entities.Truck, error)
}

type GetImplementation struct {
	TruckProvider providers.TruckProvider
}

func (uc *GetImplementation) Execute(ctx context.Context) (*entities.Truck, error) {
	ctx = context.WithValue(ctx, "GET", "get_truck")
	truck, err := uc.TruckProvider.GetTruck(ctx)
	if err != nil {
		return nil, err
	}
	return truck, nil
}
