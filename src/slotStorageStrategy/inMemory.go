package slotStorageStrategy

import (
	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot"
	parkingSlotType "github.com/aryamansingh2008/ParkingLot/src/parkingSlot/parkingSlotType"
	"github.com/aryamansingh2008/ParkingLot/src/slotStorageStrategy/errors"
)

type InMemoryStorage struct {
	slots map[int]*parkingSlot.ParkingSlot
}

func NewInMemoryStorage(size int) (*InMemoryStorage, error) {
	slots := make(map[int]*parkingSlot.ParkingSlot, size)
	for slotID := 1; slotID <= size; slotID++ {
		slot, err := parkingSlot.NewParkingSlot(slotID, parkingSlotType.LargeSlot{})
		if err != nil {
			return nil, err
		}

		slots[slotID] = slot
	}

	return &InMemoryStorage{
		slots: slots,
	}, nil
}

func (s *InMemoryStorage) GetSlot(slotID int) (*parkingSlot.ParkingSlot, error) {
	slot, exists := s.slots[slotID]
	if !exists {
		return nil, errors.SlotIdNotFound
	}

	return slot, nil
}

func (s *InMemoryStorage) UpdateSlot(slot *parkingSlot.ParkingSlot) error {
	s.slots[slot.ID()] = slot
	return nil
}

func (s *InMemoryStorage) GetAllSlots() ([]*parkingSlot.ParkingSlot, error) {
	allSlots := make([]*parkingSlot.ParkingSlot, 0, len(s.slots))
	for _, slot := range s.slots {
		allSlots = append(allSlots, slot)
	}

	return allSlots, nil
}
