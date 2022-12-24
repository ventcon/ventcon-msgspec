package msgspec

const CommandTopicPrefix string = "ventcon.command"

type Command string

const (
	SetIntakeAirLevel  Command = "setIntakeAirLevel"
	SetExhaustAirLevel Command = "setExhaustAirLevel"
)

type CommandMessage struct {
	// Address is the numeric address of the ventilator this command is for
	Address int     `json:"address"`
	Command Command `json:"command"`
}

type SetIntakeAirLevelCommandMessage struct {
	CommandMessage
	// IntakeAirLevel is a number from 0 to 10 indicating on which level the air intake fan should be running
	IntakeAirLevel int `json:"intakeAirLevel"`
}

type SetExhaustAirLevelCommandMessage struct {
	CommandMessage
	// ExhaustAirLevel is a number from 0 to 10 indicating on which level the air exhaust fan should be running
	ExhaustAirLevel int `json:"exhaustAirLevel"`
}
