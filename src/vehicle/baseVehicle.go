package vehicle

import (
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/errors"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/types"
)

type BaseVehicle struct {
	registrationNo string
	color          types.Color
}

func NewBaseVehicle(registrationNo string, color types.Color) (*BaseVehicle, error) {
	if registrationNo == "" {
		return nil, errors.InvalidRegistration
	}

	return &BaseVehicle{
		registrationNo: registrationNo,
		color:          color,
	}, nil
}
