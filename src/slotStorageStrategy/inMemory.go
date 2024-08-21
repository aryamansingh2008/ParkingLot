package slotStorageStrategy

import (
	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot"
	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot/parkingSlotType"
	"github.com/aryamansingh2008/ParkingLot/src/slotStorageStrategy/errors"
)

type inMemoryStorage struct {
	slots []*parkingSlot.ParkingSlot
}

func NewInMemoryStorage(size int) (*inMemoryStorage, error) {
	slots := make([]*parkingSlot.ParkingSlot, 0, size)
	for slotID := 1; slotID <= size; slotID++ {
		slot, err := parkingSlot.NewParkingSlot(slotID, parkingSlotType.LargeSlot{})
		if err != nil {
			return nil, err
		}

		slots = append(slots, slot)
	}

	return &inMemoryStorage{
		slots: slots,
	}, nil
}

func (s *inMemoryStorage) GetSlot(slotID int) (*parkingSlot.ParkingSlot, error) {
	for _, slot := range s.slots {
		if slot.ID() == slotID {
			return slot, nil
		}
	}

	return nil, errors.ErrSlotIdNotFound
}

func (s *inMemoryStorage) UpdateSlot(updatedSlot *parkingSlot.ParkingSlot) error {
	for _, slot := range s.slots {
		if slot.ID() == updatedSlot.ID() {
			slot = updatedSlot
			return nil
		}
	}

	return errors.ErrSlotIdNotFound
}

func (s *inMemoryStorage) GetAllSlots() ([]*parkingSlot.ParkingSlot, error) {
	return s.slots, nil
}
