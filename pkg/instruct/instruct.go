package instruct

import (
	"context"
)

type Instruct struct {
	message string `json:"message"`
	state   string `json:"state"`
}

type InstructionBoard struct {
	ctx     context.Context
	Channel chan *Instruct
}

func (instruct *InstructionBoard) Startup(ctx context.Context) {
	instruct.ctx = ctx
}

func NewInstructBoard() *InstructionBoard {
	var InstructChannel chan *Instruct
	return &InstructionBoard{Channel: InstructChannel}
}

func (i *InstructionBoard) InstructionInfo(message string) {
	newInstruct := &Instruct{message: message, state: "info"}
	i.Channel <- newInstruct
}

func (i *InstructionBoard) InstructionAlert(message string) {
	newInstruct := &Instruct{message: message, state: "alert"}
	i.Channel <- newInstruct
}

func (i *InstructionBoard) InstructionError(message string) {
	newInstruct := &Instruct{message: message, state: "error"}
	i.Channel <- newInstruct
}

func (i *InstructionBoard) GetNewInstruct() *Instruct {
	instruct := <-i.Channel
	return instruct
}
