package event

import "time"

type BalanceCreated struct {
	Name    string
	Payload interface{}
}

func NewBalanceCreated() *BalanceUpdated {
	return &BalanceUpdated{
		Name: "BalanceCreated",
	}
}

func (e *BalanceCreated) GetName() string {
	return e.Name
}

func (e *BalanceCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *BalanceCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *BalanceCreated) GetDateTime() time.Time {
	return time.Now()
}
