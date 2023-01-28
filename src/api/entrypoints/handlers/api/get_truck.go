package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/contracts/gettruck"
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/entities"
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/errors"
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/usecases/truck"
)

type Get interface {
	Execute(ctx context.Context) (*entities.Truck, error)
}

type GetImplementation struct {
}

type GetTruck struct {
	GetUseCase truck.Get
}

func (handler *GetTruck) Handler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	truck, err := handler.GetUseCase.Execute(ctx)
	if err != nil {
		return errors.GetCmmonAPIErrors(err)
	}

	response := gettruck.NewResponse(*truck)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}
