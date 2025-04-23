package msgspec

import (
	"encoding/json"
	"errors"
)

type commandMessage struct {
	// Address is the numeric address of the ventilator this command is for
	Address_ int `json:"address"`
	// CommandType is the type of the command
	CommandType CommandType `json:"commandType"`
	// Command is the actual command to be executed
	Command_   Command          `json:"-"`
	RawCommand *json.RawMessage `json:"command"`
}

func (c *commandMessage) Address() int {
	return c.Address_
}
func (c *commandMessage) Command() Command {
	return c.Command_
}

func NewCommandMessage(address int, command Command) CommandMessage {
	return &commandMessage{
		Address_:    address,
		CommandType: command.getCommandType(),
		Command_:    command,
		RawCommand:  nil,
	}
}

func EmptyCommandMessage() CommandMessage {
	return &commandMessage{
		Address_:    0,
		CommandType: "",
		Command_:    nil,
		RawCommand:  nil,
	}
}

func (c *commandMessage) UnmarshalJSON(b []byte) error {
	type commandMessage_ commandMessage
	err := json.Unmarshal(b, (*commandMessage_)(c))
	if err != nil {
		return err
	}
	switch c.CommandType {
	case PollVentilatorNow:
		c.Command_ = &PollVentilatorNowCommand{}
	case SetRemoteCommander:
		c.Command_ = &SetRemoteCommanderCommand{}
	case SetVentilationMode:
		c.Command_ = &SetVentilationModeCommand{}
	case SetIntakeAirLevel:
		c.Command_ = &SetIntakeAirLevelCommand{}
	case SetExhaustAirLevel:
		c.Command_ = &SetExhaustAirLevelCommand{}
	case SetBothAirLevel:
		c.Command_ = &SetBothAirLevelCommand{}
	default:
		return errors.New("unknown command type")
	}
	err = json.Unmarshal(*c.RawCommand, c.Command_)
	if err != nil {
		return err
	}
	c.RawCommand = nil
	return nil
}

func (c *commandMessage) MarshalJSON() ([]byte, error) {
	type commandMessage_ commandMessage

	b, err := json.Marshal(c.Command_)
	if err != nil {
		return nil, err
	}
	var rawCommand json.RawMessage = b
	c.RawCommand = &rawCommand

	return json.Marshal((*commandMessage_)(c))
}
