package message

import (
	"honeysheet-server/web/websocket"
)

// TypeStrategy Execute different strategies according to different message Type
type TypeStrategy interface {
	execute(message *websocket.MessageStruct)
}

// TypeContext Strategy context
type TypeContext struct {
	TypeStrategy TypeStrategy
}

func (tc *TypeContext) setStrategy(ts TypeStrategy) {
	tc.TypeStrategy = ts
}

func (tc *TypeContext) executeStrategy(message *websocket.MessageStruct) {
	tc.TypeStrategy.execute(message)
}
