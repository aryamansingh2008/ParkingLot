package factory

import (
	commonTypes "github.com/aryamansingh2008/ParkingLot/src/common/types"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
	types "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type IVehicleFactory interface {
	CreateVehicle(vehicleType types.VehicleType, registrationNo string,
		color commonTypes.Color) (vehicle.IVehicle, error)
}
