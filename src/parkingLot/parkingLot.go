package parkingLot

import (
	"strconv"

	"github.com/aryamansingh2008/ParkingLot/src/parkingLot/errors"
	"github.com/aryamansingh2008/ParkingLot/src/slotAllocationStrategy"
	"github.com/aryamansingh2008/ParkingLot/src/slotStorageStrategy"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/factory"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type ParkingLot struct {
	storageStrategy    slotStorageStrategy.ISlotStorageStrategy
	allocationStrategy slotAllocationStrategy.ISlotAllocationStrategy
	vehicleFactory     factory.IVehicleFactory
}

func NewParkingLot(storageStrategy slotStorageStrategy.ISlotStorageStrategy,
	allocationStrategy slotAllocationStrategy.ISlotAllocationStrategy,
	vehicleFactory factory.IVehicleFactory) (*ParkingLot, error) {

	return &ParkingLot{
		storageStrategy:    storageStrategy,
		allocationStrategy: allocationStrategy,
		vehicleFactory:     vehicleFactory,
	}, nil
}

func (pl *ParkingLot) ParkVehicle(vehicleType types.VehicleType, registrationNo string,
	color types.Color) (string, error) {
	vehicle, err := pl.vehicleFactory.CreateVehicle(vehicleType, registrationNo, color)
	if err != nil {
		return "", err
	}

	slot, err := pl.allocationStrategy.AllocateSlot(pl.storageStrategy, vehicle)
	if err != nil {
		return "", err
	}

	if err := slot.Park(vehicle); err != nil {
		return "", err
	}

	if err := pl.storageStrategy.UpdateSlot(slot); err != nil {
		return "", err
	}

	return strconv.Itoa(slot.ID()), nil
}

func (pl *ParkingLot) Vacate(slotID int) error {
	slot, err := pl.storageStrategy.GetSlot(slotID)
	if err != nil {
		return err
	}

	if !slot.IsOccupied() {
		return errors.VehicleNotFound
	}

	slot.Vacate()
	return nil
}

func (pl *ParkingLot) Status() (string, error) {
	slots, err := pl.storageStrategy.GetAllSlots()
	if err != nil {
		return "", err
	}

	status := ""
	for _, slot := range slots {
		status += slot.Status()
	}
	return status, nil
}

func (pl *ParkingLot) FindVehiclesByColor(color types.Color) (string, error) {
	slots, err := pl.storageStrategy.GetAllSlots()
	if err != nil {
		return "", err
	}

	result := ""
	for _, slot := range slots {
		if slot.IsOccupied() && slot.Vehicle().Color() == color {
			result += slot.Status()
		}
	}
	return result, nil
}

func (pl *ParkingLot) FindVehicleByRegistrationNo(registrationNo string) (string, error) {
	slots, err := pl.storageStrategy.GetAllSlots()
	if err != nil {
		return "", err
	}

	for _, slot := range slots {
		if slot.IsOccupied() && slot.Vehicle().RegistrationNo() == registrationNo {
			return slot.Status(), nil
		}
	}

	return "", errors.VehicleNotFound
}

func (pl *ParkingLot) FindVehicleBySlotID(slotID int) (string, error) {
	slot, err := pl.storageStrategy.GetSlot(slotID)
	if err != nil {
		return "", err
	}

	return slot.Status(), nil
}
