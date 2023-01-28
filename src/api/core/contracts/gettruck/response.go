package gettruck

import "github.com/engmanuelricardoperez/fast-truck-api/src/api/core/entities"

type Response struct {
	TruckID        int64  `json:"truck_id"`
	TotalInventory int64  `json:"total_inventory"`
	Description    string `json:"description"`
}

func NewResponse(truck entities.Truck) Response {
	return Response{
		TruckID:        truck.TruckID,
		TotalInventory: truck.TotalInventory,
		Description:    truck.Description,
	}
}
