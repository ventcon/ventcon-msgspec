package msgspec

import (
	"encoding/json"
	"testing"

	"github.com/shoenig/test"
	"github.com/shoenig/test/must"
)

func TestCommandMessageSerialization(t *testing.T) {
	message := NewCommandMessage(
		123,
		&SetRemoteCommanderCommand{
			RemoteCommander: true,
		},
	)
	bytes, err := json.Marshal(message)
	must.NoError(t, err)
	jsonString := string(bytes)
	test.Eq(t, "{\"address\":123,\"commandType\":\"setRemoteCommander\",\"command\":{\"remoteCommander\":true}}", jsonString)
}

func TestCommandMessageDeserialization(t *testing.T) {
	message := EmptyCommandMessage()
	err := json.Unmarshal([]byte("{\"address\":123,\"commandType\":\"setRemoteCommander\",\"command\":{\"remoteCommander\":true}}"), &message)
	must.NoError(t, err)
	test.Eq(t, 123, message.Address())

	val, ok := message.Command().(*SetRemoteCommanderCommand)
	test.True(t, ok)
	test.True(t, val.RemoteCommander)
}

func TestCommandMessageDeserializationWithMissingCommand(t *testing.T) {
	message := EmptyCommandMessage()
	err := json.Unmarshal([]byte("{\"address\":123,\"commandType\":\"pollVentilatorNow\"}"), &message)
	must.NoError(t, err)
	test.Eq(t, 123, message.Address())

	_, ok := message.Command().(*PollVentilatorNowCommand)
	test.True(t, ok)
}

func TestAllCommandMessagesSerialization(t *testing.T) {
	testCases := []struct {
		name           string
		commandMessage CommandMessage
	}{
		{"PollVentilatorNowCommand", NewCommandMessage(11, &PollVentilatorNowCommand{})},
		{"SetRemoteCommanderCommand: true", NewCommandMessage(12, &SetRemoteCommanderCommand{RemoteCommander: true})},
		{"SetRemoteCommanderCommand: false", NewCommandMessage(13, &SetRemoteCommanderCommand{RemoteCommander: false})},
		{"SetVentilationModeCommand: true", NewCommandMessage(23, &SetVentilationModeCommand{VentilationMode: true})},
		{"SetVentilationModeCommand: false", NewCommandMessage(24, &SetVentilationModeCommand{VentilationMode: false})},
		{"SetIntakeAirLevelCommand", NewCommandMessage(34, &SetIntakeAirLevelCommand{IntakeAirLevel: 5})},
		{"SetExhaustAirLevelCommand", NewCommandMessage(45, &SetExhaustAirLevelCommand{ExhaustAirLevel: 5})},
		{"SetBothAirLevelCommand", NewCommandMessage(56, &SetBothAirLevelCommand{IntakeAirLevel: 7, ExhaustAirLevel: 8})},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bytes, err := json.Marshal(tc.commandMessage)
			must.NoError(t, err)

			decodedMessage := EmptyCommandMessage()
			err = json.Unmarshal(bytes, &decodedMessage)
			must.NoError(t, err)

			test.Eq(t, tc.commandMessage.Address(), decodedMessage.Address())
			test.Eq(t, tc.commandMessage.Command().getCommandType(), decodedMessage.Command().getCommandType())
			test.Eq(t, tc.commandMessage.Command(), decodedMessage.Command())
		})
	}
}
