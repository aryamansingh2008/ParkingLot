package vehicle

import (
	commonTypes "github.com/aryamansingh2008/ParkingLot/src/common/types"
	vehicleType "github.com/aryamansingh2008/ParkingLot/src/vehicle/constants/vehicleType"
	types "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type Car struct {
	BaseVehicle
}

func (c *Car) Type() types.VehicleType {
	return vehicleType.Car
}

func (c *Car) RegistrationNo() string {
	return c.registrationNo
}

func (c *Car) Color() commonTypes.Color {
	return c.color
}

func NewCar(registrationNo string, color commonTypes.Color) (*Car, error) {
	baseVehicle, err := NewBaseVehicle(registrationNo, color)
	if err != nil {
		return nil, err
	}

	return &Car{
		BaseVehicle: *baseVehicle,
	}, nil
}
