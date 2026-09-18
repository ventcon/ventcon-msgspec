# ventcon-msgspec

[![Go Reference](https://pkg.go.dev/badge/github.com/ventcon/ventcon-msgspec.svg)](https://pkg.go.dev/github.com/ventcon/ventcon-msgspec)
[![CI](https://github.com/ventcon/ventcon-msgspec/actions/workflows/ci.yml/badge.svg)](https://github.com/ventcon/ventcon-msgspec/actions/workflows/ci.yml)

The specification of the messaging format for my ventcon project.

## Specification

- All messages are sent via the [NATS](https://nats.io/) messaging protocol using JSON encoded payloads.
- [`heartbeat.go`](heartbeat.go) defines the message format and subject for the broadcasted heartbeat message.
  - Heartbeat messages are repeatedly published by the hardware interface on the topic `ventcon.heartbeat` (defined in variable `HeartbeatTopic`).
  - They contain information on the state of the ventilation system
  - The frequency and variance are implementation specific.
  - No message durability is used.
- [`commands.go`](commands.go) defines the message formats and subjects for the possible commands.
  - Commands instruct one component to do some action.
  - They are sent using a [NATS JetStream](https://docs.nats.io/nats-concepts/jetstream) and durable consumers.
  - The receiver only acknowledges a message once it can guarantee its completion
    - (either by it being completed or by storing it persistently).
  - The topic for each command is produced by joining the `CommandTopicPrefix` which is `ventcon.command` followed by a dot, followed by the respective ventilator id as a decimal string.

## Examples

The following examples show how to interact with the ventcon protocol using the natscli command line tool [^1].

[^1]: https://github.com/nats-io/natscli

### Receiving Heartbeats

Subscribe to the heartbeat topic by calling

``` bash
nats sub 'ventcon.heartbeat' -r
```

You will receive messages like this:

``` json
{
  "address":1,
  "online":false,
  "data":{
    "remoteCommander":false,
    "ventilationMode":false,
    "requestedIntakeAirLevel":0,
    "requestedExhaustAirLevel":0,
    "actualIntakeAirLevel":0,
    "actualExhaustAirLevel":0,
    "externalSwitchPosition":0,
    "filterInstalled":false,
    "filterDirty":false,
    "frostRisk":false,
    "exhaustAirTemp":0,
    "roomAirTemp":0,
    "externalAirTemp":0,
    "inletAirTemp":0
  }
}
```

### Sending Commands

To set the ventilator with id 50 into remote command mode call

```bash
nats pub ventcon.command.50 '{"address":50,"commandType":"setRemoteCommander","command":{"remoteCommander":true}}'
```

Those are the different command payloads currently supported:

```json
'{"address":50,"commandType":"pollVentilatorNow", "command": {}}'
```

```json
'{"address":50,"commandType":"setRemoteCommander", "command": {"remoteCommander": true}}'
```

```json
'{"address":50,"commandType":"setVentilationMode", "command": {"ventilationMode": true}}'
```

```json
'{"address":50,"commandType":"setIntakeAirLevel", "command": {"intakeAirLevel": 1}}'
```

```json
'{"address":50,"commandType":"setExhaustAirLevel", "command": {"exhaustAirLevel": 1}}'
```

```json
'{"address":50,"commandType":"setBothAirLevel", "command": {"intakeAirLevel": 1, "exhaustAirLevel": 1}}'
```
