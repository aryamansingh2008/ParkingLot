package main

import (
	"github.com/aryamansingh2008/ParkingLot/src/driver"
	"github.com/aryamansingh2008/ParkingLot/src/driver/commandLineDriver"
)

func main() {
	var driver driver.IDriver
	driver = commandLineDriver.NewCommandLineInput()

	driver.Start()
}
