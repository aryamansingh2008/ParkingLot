package parkingslotstoragestrategy

import parkingslot "github.com/aryamansingh2008/ParkingLot/src/parkingSlot"

type IParkingSlotStorageStragefy interface {
	GetSlot(slotID int) (*parkingslot.ParkingSlot, error)
	UpdateSlot(slot *parkingslot.ParkingSlot) error
	GetAllSlots() ([]*parkingslot.ParkingSlot, error)
}
