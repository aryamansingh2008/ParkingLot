package vehicle

import (
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/constants/vehicleTypes"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type Car struct {
	BaseVehicle
}

func (c *Car) Type() types.VehicleType {
	return vehicleTypes.Car
}

func (c *Car) RegistrationNo() string {
	return c.registrationNo
}

func (c *Car) Color() types.Color {
	return c.color
}

func NewCar(registrationNo string, color types.Color) (*Car, error) {
	baseVehicle, err := NewBaseVehicle(registrationNo, color)
	if err != nil {
		return nil, err
	}

	return &Car{
		BaseVehicle: *baseVehicle,
	}, nil
}
