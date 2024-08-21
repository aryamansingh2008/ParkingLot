package vehicle

import (
	types "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type IVehicle interface {
	Type() types.VehicleType
	RegistrationNo() string
	Color() types.Color
}
