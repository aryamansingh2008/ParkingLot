package commandLineDriver

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aryamansingh2008/ParkingLot/src/driver/errors"
	"github.com/aryamansingh2008/ParkingLot/src/parkingLot"
)

type CommandCallback func(args []string) (string, error)

type CommandLineDriver struct {
	parkingLot       *parkingLot.ParkingLot
	commandCallbacks map[string]CommandCallback
}

func NewCommandLineInput() *CommandLineDriver {
	return &CommandLineDriver{
		commandCallbacks: make(map[string]CommandCallback),
	}
}

func (d *CommandLineDriver) Start() {
	d.registerCallbacks()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		output, err := d.HandleCommand(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("ProcessInput: error: ", err)
		} else {
			fmt.Println(output)
		}
	}
}

func (d *CommandLineDriver) HandleCommand(input string) (string, error) {
	args := strings.Split(input, " ")

	command, err := d.getCommand(args)
	if err != nil {
		return "", err
	}

	if cb, exists := d.commandCallbacks[command]; exists {
		return cb(args)
	}
	return "", errors.ErrInvalidRequest
}

func (d *CommandLineDriver) isValidCommand(args []string, minLength int) bool {
	if len(args) < minLength {
		return false
	}

	return true
}

func (d *CommandLineDriver) getCommand(args []string) (string, error) {
	if !d.isValidCommand(args, 1) {
		return "", errors.ErrInvalidRequest
	}

	return args[0], nil
}
