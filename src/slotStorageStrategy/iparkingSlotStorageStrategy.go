package slotStorageStrategy

import "github.com/aryamansingh2008/ParkingLot/src/parkingSlot"

type IParkingSlotStorageStrategy interface {
	GetSlot(slotID int) (*parkingSlot.ParkingSlot, error)
	UpdateSlot(slot *parkingSlot.ParkingSlot) error
	GetAllSlots() ([]*parkingSlot.ParkingSlot, error)
}
