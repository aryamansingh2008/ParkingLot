package factory

import (
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/constants/vehicleTypes"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/errors"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type vehicleFactory struct{}

func NewVehicleFactory() *vehicleFactory {
	return &vehicleFactory{}
}

func (f *vehicleFactory) CreateVehicle(vehicleType types.VehicleType, registrationNo string,
	color types.Color) (vehicle.IVehicle, error) {
	switch vehicleType {
	case vehicleTypes.Car:
		return vehicle.NewCar(registrationNo, color)
	default:
		return nil, errors.InvalidType
	}
}
