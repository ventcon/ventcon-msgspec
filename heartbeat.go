package msgspec

const HeartbeatTopic string = "ventcon.heartbeat"

// VentilatorData contains data about the state of a single ventilator
type VentilatorData struct {
	// IntakeAirLevel is a number from 0 to 10 indicating on which level the air intake fan is running
	IntakeAirLevel int `json:"intakeAirLevel"`
	// ExhaustAirLevel is a number from 0 to 10 indicating on which level the air exhaust fan is running
	ExhaustAirLevel int `json:"exhaustAirLevel"`
	// ExternalSwitchPosition is a number from 1 to 3 indicating on which position the external switch of the fan is
	ExternalSwitchPosition int `json:"externalSwitchPosition"`
	// FilterInstalled is a flag indictaing whether the ventilator has a filter installed
	FilterInstalled bool `json:"filterInstalled"`
	// FilterDirty is a flag indictaing whether the filter of the ventilator is dirty
	FilterDirty bool `json:"filterDirty"`
	// FrostRisk is a flag indictaing whether the ventilator detected the risk of freezing
	FrostRisk bool `json:"frostRisk"`
	// ExternalAirTemp is the temperature of the intake air before the heat exchanger
	ExternalAirTemp float32 `json:"externalAirTemp"`
	// InletAirTemp is the temperature of the intake air after the heat exchanger
	InletAirTemp float32 `json:"inletAirTemp"`
	// RoomAirTemp is the temperature of the exhausted air before the heat exchanger
	RoomAirTemp float32 `json:"roomAirTemp"`
	// ExhaustAirTemp is the temperature of the exhausted air after the heat exchanger
	ExhaustAirTemp float32 `json:"exhaustAirTemp"`
	// InletAirHumidity is the air humidity in percent of the intake air before the heat exchanger
	InletAirHumidity int `json:"inletAirHumidity"`
	// ExhaustAirHumidity is the air humidity in percent of the exhausted air before the heat exchanger
	ExhaustAirHumidity int `json:"exhaustAirHumidity"`
	// MixedGasConcentration is the concentration of mixed gas in the room in parts per million.
	// If the ventilator does not support this, it is omitted
	MixedGasConcentration int `json:"mixedGasConcentration,omitempty"`
}

type VentilatorState struct {
	// Address is the numeric address of the ventilator in question
	Address int `json:"address"`
	// Online is a flag indicating if that ventilator can currently be reached
	Online bool `json:"online"`
	// Data is the data of the ventilator. If the ventilator is offline, this is omitted.
	// The age of the actual data is implementation specific and may vary between different properties.
	Data VentilatorData `json:"data,omitempty"`
}

// The message payload of a heartbeet message
type HeartbeatMessage struct {
	// CurrentState is mapping the ventilators address to its state.
	CurrentState map[int]VentilatorState
}
