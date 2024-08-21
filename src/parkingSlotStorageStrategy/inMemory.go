package parkingSlotStorageStrategy

import (
	parkingslot "github.com/aryamansingh2008/ParkingLot/src/parkingSlot"
	parkingslottype "github.com/aryamansingh2008/ParkingLot/src/parkingSlot/parkingSlotType"
	"github.com/aryamansingh2008/ParkingLot/src/parkingSlotStorageStrategy/errors"
)

type InMemoryStorage struct {
	slots map[int]*parkingslot.ParkingSlot
}

func NewInMemoryStorage(size int) (*InMemoryStorage, error) {
	slots := make(map[int]*parkingslot.ParkingSlot, size)
	for slotID := 1; slotID <= size; slotID++ {
		slot, err := parkingslot.NewParkingSlot(slotID, parkingslottype.LargeSlot{})
		if err != nil {
			return nil, err
		}

		slots[slotID] = slot
	}

	return &InMemoryStorage{
		slots: slots,
	}, nil
}

func (s *InMemoryStorage) GetSlot(slotID int) (*parkingslot.ParkingSlot, error) {
	slot, exists := s.slots[slotID]
	if !exists {
		return nil, errors.SlotIdNotFound
	}

	return slot, nil
}

func (s *InMemoryStorage) UpdateSlot(slot *parkingslot.ParkingSlot) error {
	s.slots[slot.ID()] = slot
	return nil
}

func (s *InMemoryStorage) GetAllSlots() ([]*parkingslot.ParkingSlot, error) {
	allSlots := make([]*parkingslot.ParkingSlot, 0, len(s.slots))
	for _, slot := range s.slots {
		allSlots = append(allSlots, slot)
	}

	return allSlots, nil
}
