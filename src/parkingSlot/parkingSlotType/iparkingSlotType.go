package parkingSloTtype

import "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"

type ParkingSlotType interface {
	CanAccommodate(types.VehicleType) bool
}
