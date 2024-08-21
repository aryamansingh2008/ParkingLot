package factory

import (
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
	types "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type IVehicleFactory interface {
	CreateVehicle(vehicleType types.VehicleType, registrationNo string,
		color types.Color) (vehicle.IVehicle, error)
}
