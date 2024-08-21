package vehicle

import (
	commonTypes "github.com/aryamansingh2008/ParkingLot/src/common/types"
	types "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type IVehicle interface {
	Type() types.VehicleType
	RegistrationNo() string
	Color() commonTypes.Color
}
