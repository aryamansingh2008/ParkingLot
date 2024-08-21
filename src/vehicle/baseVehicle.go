package vehicle

import (
	colorConstants "github.com/aryamansingh2008/ParkingLot/src/common/constants/color"
	"github.com/aryamansingh2008/ParkingLot/src/common/types"
	"github.com/aryamansingh2008/ParkingLot/src/vehicle/errors"
)

type BaseVehicle struct {
	registrationNo string
	color          types.Color
}

func NewBaseVehicle(registrationNo string, color types.Color) (*BaseVehicle, error) {
	if registrationNo == "" {
		return nil, errors.InvalidRegistration
	}

	if color == colorConstants.None {
		return nil, errors.InvalidColor
	}

	return &BaseVehicle{
		registrationNo: registrationNo,
		color:          color,
	}, nil
}
