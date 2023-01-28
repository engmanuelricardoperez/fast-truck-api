package dependencies

import (
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/usecases/truck"
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/entrypoints"
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/entrypoints/handlers/api"
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/repositories/store/truck_repository"
)

type HandlerContainer struct {
	GetTruck entrypoints.Handler
}

type Dependencies struct {
}

func (d Dependencies) Start() *HandlerContainer {
	// Providers
	truckProvider := &truck_repository.Repository{}

	// Usecase
	getTruckUseCase := &truck.GetImplementation{
		TruckProvider: truckProvider,
	}

	// Handlers
	handlers := HandlerContainer{}
	handlers.GetTruck = &api.GetTruck{
		GetUseCase: getTruckUseCase,
	}

	return &handlers

}
