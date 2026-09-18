package msgspec

const CommandTopicPrefix string = "ventcon.command"

type CommandType string

const (
	PollVentilatorNow  CommandType = "pollVentilatorNow"
	SetRemoteCommander CommandType = "setRemoteCommander"
	SetVentilationMode CommandType = "setVentilationMode"
	SetIntakeAirLevel  CommandType = "setIntakeAirLevel"
	SetExhaustAirLevel CommandType = "setExhaustAirLevel"
	SetBothAirLevel    CommandType = "setBothAirLevel"
)

type Command interface {
	getCommandType() CommandType
}

type CommandMessage interface {
	// Address is the numeric address of the ventilator this command is for
	Address() int
	// Command is the actual command to be executed
	Command() Command
}

type PollVentilatorNowCommand struct{}

func (PollVentilatorNowCommand) getCommandType() CommandType {
	return PollVentilatorNow
}

type SetRemoteCommanderCommand struct {
	// RemoteCommander is a flag indicating whether the ventilator should be controlled by a remote commander
	RemoteCommander bool `json:"remoteCommander"`
}

func (SetRemoteCommanderCommand) getCommandType() CommandType {
	return SetRemoteCommander
}

type SetVentilationModeCommand struct {
	// VentilationMode is a flag indicating whether the ventilator should be in ventilation mode
	VentilationMode bool `json:"ventilationMode"`
}

func (SetVentilationModeCommand) getCommandType() CommandType {
	return SetVentilationMode
}

type SetIntakeAirLevelCommand struct {
	// IntakeAirLevel is a number from 0 to 10 indicating on which level the air intake fan should be running
	IntakeAirLevel int `json:"intakeAirLevel"`
}

func (SetIntakeAirLevelCommand) getCommandType() CommandType {
	return SetIntakeAirLevel
}

type SetExhaustAirLevelCommand struct {
	// ExhaustAirLevel is a number from 0 to 10 indicating on which level the air exhaust fan should be running
	ExhaustAirLevel int `json:"exhaustAirLevel"`
}

func (SetExhaustAirLevelCommand) getCommandType() CommandType {
	return SetExhaustAirLevel
}

type SetBothAirLevelCommand struct {
	// IntakeAirLevel is a number from 0 to 10 indicating on which level the air intake fan should be running
	IntakeAirLevel int `json:"intakeAirLevel"`
	// ExhaustAirLevel is a number from 0 to 10 indicating on which level the air exhaust fan should be running
	ExhaustAirLevel int `json:"exhaustAirLevel"`
}

func (SetBothAirLevelCommand) getCommandType() CommandType {
	return SetBothAirLevel
}
