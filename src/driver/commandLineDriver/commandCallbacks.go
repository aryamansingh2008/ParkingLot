package commandLineDriver

import (
	"strconv"

	"github.com/aryamansingh2008/ParkingLot/src/driver/constants/commands"
	"github.com/aryamansingh2008/ParkingLot/src/driver/errors"
	"github.com/aryamansingh2008/ParkingLot/src/parkingLot"
	"github.com/aryamansingh2008/ParkingLot/src/slotAllocationStrategy"
	"github.com/aryamansingh2008/ParkingLot/src/slotStorageStrategy"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/factory"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

func (d *CommandLineDriver) registerCallbacks() {
	d.commandCallbacks[commands.Create] = d.CreateCallback
	d.commandCallbacks[commands.Status] = d.StatusCallback
	d.commandCallbacks[commands.Park] = d.ParkCallback
	d.commandCallbacks[commands.Vacate] = d.VacateBySlotIDCallback
	d.commandCallbacks[commands.FindVehiclesByColor] = d.FindVehiclesByColorCallback
	d.commandCallbacks[commands.FindVehicleBySlotID] = d.FindVehicleBySlotIDCallback
	d.commandCallbacks[commands.FindVehicleByRegistrationNo] = d.FindVehicleByRegistrationNoCallback
}

func (d *CommandLineDriver) CreateCallback(args []string) (string, error) {
	if !d.isValidCommand(args, 2) {
		return "", errors.ErrInvalidRequest
	}

	size, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	storageStrategy, err := slotStorageStrategy.NewInMemoryStorage(size)
	if err != nil {
		return "", err
	}
	allocationStrategy := slotAllocationStrategy.NewNearestFirstStrategy()
	vehicleFactory := factory.NewVehicleFactory()

	parkingLot, err := parkingLot.NewParkingLot(storageStrategy, allocationStrategy, vehicleFactory)
	if err != nil {
		return "", err
	}

	d.parkingLot = parkingLot
	return "Parking lot created.\n", nil
}

func (d *CommandLineDriver) StatusCallback(args []string) (string, error) {
	if d.parkingLot == nil {
		return "", errors.ErrInvalidRequest
	}

	return d.parkingLot.Status()
}

func (d *CommandLineDriver) ParkCallback(args []string) (string, error) {
	if d.parkingLot == nil {
		return "", errors.ErrInvalidRequest
	}

	if !d.isValidCommand(args, 4) {
		return "", errors.ErrInvalidRequest
	}

	vehicleTypeStr := args[1]
	registrationNo := args[2]
	color := args[3]

	output, err := d.parkingLot.ParkVehicle(types.VehicleType(vehicleTypeStr), registrationNo, types.Color(color))
	return output + "\n", err
}

func (d *CommandLineDriver) VacateBySlotIDCallback(args []string) (string, error) {
	if d.parkingLot == nil {
		return "", errors.ErrInvalidRequest
	}

	if !d.isValidCommand(args, 2) {
		return "", errors.ErrInvalidRequest
	}

	slotID, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	err = d.parkingLot.Vacate(slotID)
	if err != nil {
		return "", err
	}

	return "Parking Slot Vacated\n", nil
}

func (d *CommandLineDriver) FindVehiclesByColorCallback(args []string) (string, error) {
	if d.parkingLot == nil {
		return "", errors.ErrInvalidRequest
	}

	if !d.isValidCommand(args, 2) {
		return "", errors.ErrInvalidRequest
	}

	color := args[1]
	return d.parkingLot.FindVehiclesByColor(types.Color(color))
}

func (d *CommandLineDriver) FindVehicleBySlotIDCallback(args []string) (string, error) {
	if d.parkingLot == nil {
		return "", errors.ErrInvalidRequest
	}

	if !d.isValidCommand(args, 2) {
		return "", errors.ErrInvalidRequest
	}

	slotID, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	return d.parkingLot.FindVehicleBySlotID(slotID)
}

func (d *CommandLineDriver) FindVehicleByRegistrationNoCallback(args []string) (string, error) {
	if d.parkingLot == nil {
		return "", errors.ErrInvalidRequest
	}

	if !d.isValidCommand(args, 2) {
		return "", errors.ErrInvalidRequest
	}

	registrationNo := args[1]
	return d.parkingLot.FindVehicleByRegistrationNo(registrationNo)
}
