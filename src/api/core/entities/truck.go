package entities

type Truck struct {
	TruckID        int64
	TotalInventory int64
	Description    string
}

func NewTruck(truckID int64, totalInventory int64, description string) Truck {
	return Truck{
		TruckID:        truckID,
		TotalInventory: totalInventory,
		Description:    description,
	}
}
