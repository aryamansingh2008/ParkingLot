package slotAllocationStrategy

import (
	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot"
	"github.com/aryamansingh2008/ParkingLot/src/slotAllocationStrategy/errors"
	"github.com/aryamansingh2008/ParkingLot/src/slotStorageStrategy"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
)

type NearestFirstStrategy struct{}

func NewNearestFirstStrategy() *NearestFirstStrategy {
	return &NearestFirstStrategy{}
}

func (s *NearestFirstStrategy) AllocateSlot(storageStrategy slotStorageStrategy.IParkingSlotStorageStrategy,
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
	return nil, errors.NoValidSlotFound
}
