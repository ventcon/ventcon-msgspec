# ventcon-msgspec

[![Go Reference](https://pkg.go.dev/badge/github.com/ventcon/ventcon-msgspec.svg)](https://pkg.go.dev/github.com/ventcon/ventcon-msgspec)
[![CI](https://github.com/ventcon/ventcon-msgspec/actions/workflows/ci.yml/badge.svg)](https://github.com/ventcon/ventcon-msgspec/actions/workflows/ci.yml)

The specification of the messaging format for my ventcon project.

## Specification

- All messages are sent via the [NATS](https://nats.io/) messaging protocol using JSON encoded payloads.
- [`heartbeat.go`](heartbeat.go) defines the message format and subject for the broadcasted heartbeat message.
  - Heartbeat messages are repeatedly published by the hardware interface.
  - They contain information on the state of the ventilation system
  - The frequency and variance are implementation specific.
  - No message durability is used.
- [`commands.go`](commands.go) defines the message formats and subjects for the possible commands.
  - Commands instruct one component to do some action.
  - They are sent using a [NATS JetStream](https://docs.nats.io/nats-concepts/jetstream) and durable consumers.
  - The receiver only acknowledges a message once it can guarantee its completion
    - (either by it being completed or by storing it persistently).
  - The topic for each command is produced by joining the `CommandTopicPrefix` and the respective `Command` string with a dot.
