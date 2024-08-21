package slotAllocationStrategy

import (
	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot"
	"github.com/aryamansingh2008/ParkingLot/src/slotStorageStrategy"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
)

type ISlotAllocationStrategy interface {
	AllocateSlot(storageStrategy slotStorageStrategy.ISlotStorageStrategy,
		vehicle vehicle.IVehicle) (*parkingSlot.ParkingSlot, error)
}
