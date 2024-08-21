package slotAllocationStrategy

import (
	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot"
	"github.com/aryamansingh2008/ParkingLot/src/slotAllocationStrategy/errors"
	"github.com/aryamansingh2008/ParkingLot/src/slotStorageStrategy"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
)

type nearestFirstStrategy struct{}

func NewNearestFirstStrategy() *nearestFirstStrategy {
	return &nearestFirstStrategy{}
}

func (s *nearestFirstStrategy) AllocateSlot(storageStrategy slotStorageStrategy.ISlotStorageStrategy,
	vehicle vehicle.IVehicle) (*parkingSlot.ParkingSlot, error) {
	slots, err := storageStrategy.GetAllSlots()
	if err != nil {
		return nil, err
	}

	for _, slot := range slots {
		if slot.CanPark(vehicle) {
			return slot, nil
		}
	}
	return nil, errors.ErrNoValidSlotFound
}
