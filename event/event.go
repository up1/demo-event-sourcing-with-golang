package event

import (
	uuid "github.com/satori/go.uuid"
)

type Event struct {
	AccId string
	Type  string
}

type CreateEvent struct {
	Event
	AccName string
}

type BalanceEvent struct {
	Event
}

type DepositEvent struct {
	Event
	Amount int
}

type WithdrawEvent struct {
	Event
	Amount int
}

type TransferEvent struct {
	Event
	TargetId string
	Amount   int
}

/* helper to create events */

func NewCreateAccountEvent(name string) CreateEvent {
	event := new(CreateEvent)
	event.Type = "CreateEvent"
	event.AccId = uuid.NewV4().String()
	event.AccName = name
	return *event
}

func NewBalanceEvent(id string) BalanceEvent {
	event := new(BalanceEvent)
	event.Type = "BalanceEvent"
	event.AccId = id
	return *event
}

func NewDepositEvent(id string, amt int) DepositEvent {
	event := new(DepositEvent)
	event.Type = "DepositEvent"
	event.AccId = id
	event.Amount = amt
	return *event
}

func NewWithdrawEvent(id string, amt int) WithdrawEvent {
	event := new(WithdrawEvent)
	event.Type = "WithdrawEvent"
	event.AccId = id
	event.Amount = amt
	return *event
}

func NewTransferEvent(id string, targetId string, amt int) TransferEvent {
	event := new(TransferEvent)
	event.Type = "TransferEvent"
	event.AccId = id
	event.Amount = amt
	event.TargetId = targetId
	return *event
}
