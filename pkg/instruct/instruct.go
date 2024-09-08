package instruct

import (
	"context"
)

type Instruct struct {
	Message string `json:"message"`
	State   string `json:"state"`
}

type InstructionBoard struct {
	ctx     context.Context
	Channel chan *Instruct
}

func (i *InstructionBoard) Startup(ctx context.Context) {
	i.ctx = ctx
}

func NewInstructBoard() *InstructionBoard {
	instructWelcome := &Instruct{Message: "Welcome to Text Watchtower! Let's start by creating a group", State: "info"}
	InstructChannel := make(chan *Instruct, 1)
	InstructChannel <- instructWelcome
	return &InstructionBoard{Channel: InstructChannel}
}

func (i *InstructionBoard) InstructionInfo(message string) {
	newInstruct := &Instruct{Message: message, State: "info"}
	i.Channel <- newInstruct
}

func (i *InstructionBoard) InstructionAlert(message string) {
	newInstruct := &Instruct{Message: message, State: "alert"}
	i.Channel <- newInstruct
}

func (i *InstructionBoard) InstructionError(message string) {
	newInstruct := &Instruct{Message: message, State: "error"}
	i.Channel <- newInstruct
}

func (i *InstructionBoard) GetNewInstruct() *Instruct {
	instruct := <-i.Channel
	return instruct
}
