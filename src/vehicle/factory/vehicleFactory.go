package factory

import (
	commonTypes "github.com/aryamansingh2008/ParkingLot/src/common/types"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle"
	vehicleTypes "github.com/aryamansingh2008/ParkingLot/src/vehicle/constants/vehicleType"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/errors"
	types "github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type vehicleFactory struct{}

func NewVehicleFactory() *vehicleFactory {
	return &vehicleFactory{}
}

func (f *vehicleFactory) CreateVehicle(vehicleType types.VehicleType, registrationNo string,
	color commonTypes.Color) (vehicle.IVehicle, error) {
	switch vehicleType {
	case vehicleTypes.Car:
		return vehicle.NewCar(registrationNo, color)
	default:
		return nil, errors.InvalidType
	}
}
