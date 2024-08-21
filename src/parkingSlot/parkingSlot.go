package parkingslot

import (
	"fmt"

	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot/errors"
	parkingslottype "github.com/aryamansingh2008/ParkingLot/src/parkingSlot/parkingSlotType"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
)

type ParkingSlot struct {
	id       int
	slotType parkingslottype.ParkingSlotType
	vehicle  vehicle.IVehicle
}

func NewParkingSlot(id int, slotType parkingslottype.ParkingSlotType) (*ParkingSlot, error) {
	if id < 1 {
		return nil, errors.InvalidID
	}

	return &ParkingSlot{
		id:       id,
		slotType: slotType,
	}, nil
}

func (ps *ParkingSlot) ID() int {
	return ps.id
}

func (ps *ParkingSlot) IsOccupied() bool {
	return ps.vehicle != nil
}

func (ps *ParkingSlot) Park(vehicle vehicle.IVehicle) error {
	if ps.IsOccupied() {
		return errors.InvalidParkOccupied
	}

	if !ps.slotType.CanAccommodate(vehicle.Type()) {
		return errors.InvalidParkCantAccomodate
	}

	ps.vehicle = vehicle
	return nil
}

func (ps *ParkingSlot) Vacate() {
	ps.vehicle = nil
}

func (ps *ParkingSlot) Status() string {
	if ps.IsOccupied() {
		return fmt.Sprintf("Slot %d: Vehicle %s is parked", ps.id, ps.vehicle.RegistrationNo())
	}

	return fmt.Sprintf("Slot %d: Empty", ps.id)
}
