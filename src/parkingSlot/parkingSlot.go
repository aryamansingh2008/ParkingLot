package parkingSlot

import (
	"fmt"

	"github.com/aryamansingh2008/ParkingLot/src/parkingSlot/errors"
	parkingSlotType "github.com/aryamansingh2008/ParkingLot/src/parkingSlot/parkingSlotType"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
)

type ParkingSlot struct {
	id       int
	slotType parkingSlotType.ParkingSlotType
	vehicle  vehicle.IVehicle
}

func NewParkingSlot(id int, slotType parkingSlotType.ParkingSlotType) (*ParkingSlot, error) {
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

func (ps *ParkingSlot) Vehicle() vehicle.IVehicle {
	return ps.vehicle
}

func (ps *ParkingSlot) IsOccupied() bool {
	return ps.vehicle != nil
}

func (ps *ParkingSlot) CanPark(vehicle vehicle.IVehicle) bool {
	if ps.IsOccupied() {
		return false
	}

	if !ps.slotType.CanAccommodate(vehicle.Type()) {
		return false
	}

	return true
}

func (ps *ParkingSlot) Park(vehicle vehicle.IVehicle) error {
	if !ps.CanPark(vehicle) {
		return errors.InvalidPark
	}

	ps.vehicle = vehicle
	return nil
}

func (ps *ParkingSlot) Vacate() {
	ps.vehicle = nil
}

func (ps *ParkingSlot) Status() string {
	if ps.IsOccupied() {
		return fmt.Sprintf("Slot %d: Vehicle %s is parked\n", ps.id, ps.vehicle.RegistrationNo())
	}

	return fmt.Sprintf("Slot %d: Empty\n", ps.id)
}
