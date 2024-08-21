package driver

type IDriver interface {
	Start()

	HandleCommand(command string) (string, error)
}
