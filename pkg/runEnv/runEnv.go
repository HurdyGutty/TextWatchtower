package runEnv

import (
	"os"
	"os/exec"

	"github.com/HurdyGutty/go_OCR/pkg/instruct"
)

var InstructionBoard *instruct.InstructionBoard

func ChangeEnv() {
	path, _ := os.Getwd()
	cmd := exec.Command(path+"/Env changer.exe", "-project=Text Watchtower")
	if err := cmd.Start(); err != nil {
		InstructionBoard.InstructionError(err.Error())
	}
	err := cmd.Wait()
	if err != nil {
		InstructionBoard.InstructionError(err.Error())
	}
}

func ChangeEnvAsync() {
	go ChangeEnv()
}
