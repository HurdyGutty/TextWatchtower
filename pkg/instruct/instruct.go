package instruct

type Instruct struct {
	message string
	state   string
}

type InstructionBoard struct {
	Channel chan *Instruct
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

func (i *InstructionBoard) GetNewInstruct() (string, string) {
	instruct := <-i.Channel
	return instruct.message, instruct.state
}
