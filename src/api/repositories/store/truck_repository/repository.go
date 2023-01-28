package truck_repository

import (
	"context"

	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/entities"
)

type Repository struct {
}

func (r Repository) GetTruck(ctx context.Context) (*entities.Truck, error) {
	var truck entities.Truck
	// var err error
	// var msg message

	truck = entities.NewTruck(12345, 12, "Repository Trial")
	return &truck, nil
}
