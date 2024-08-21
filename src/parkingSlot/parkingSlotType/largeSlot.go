package parkingSloTtype

import "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"

type LargeSlot struct{}

func (ls LargeSlot) CanAccommodate(vehicle types.VehicleType) bool {
	return true
}
