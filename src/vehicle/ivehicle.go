package vehicle

import (
	"github.com/aryamansingh2008/ParkingLot/src/common/types"
)

type Vehicle interface {
	Type() VehicleType
	RegistrationNo() string
	Color() types.Color
}
