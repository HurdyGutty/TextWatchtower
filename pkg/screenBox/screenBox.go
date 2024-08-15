package screenBox

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

type ScreenBox struct {
	X1 int `json:"x"`
	Y1 int `json:"y"`
	W  int `json:"w"`
	H  int `json:"h"`
}

func DrawBox() *ScreenBox {
	newScreenBox := &ScreenBox{X1: 0, Y1: 0}
	var x2, y2 int
	fmt.Println("--- Please press Esc to stop hook ---")
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("Escaped")
		hook.End()
	})

	fmt.Println("--- Please draw rectangle---")
	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if newScreenBox.X1 == 0 && newScreenBox.Y1 == 0 {
			fmt.Println("Start")
			newScreenBox.X1 = int(e.X)
			newScreenBox.Y1 = int(e.Y)
		}
	})

	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		fmt.Println("End")
		x2 = int(e.X)
		y2 = int(e.Y)
	})

	s := hook.Start()
	<-hook.Process(s)
	if newScreenBox.X1 > x2 {
		newScreenBox.W = newScreenBox.X1 - x2
		newScreenBox.X1 = x2
	} else {
		newScreenBox.W = x2 - newScreenBox.X1
	}

	if newScreenBox.Y1 > y2 {
		newScreenBox.H = newScreenBox.Y1 - y2
		newScreenBox.Y1 = y2
	} else {
		newScreenBox.H = y2 - newScreenBox.Y1
	}

	return newScreenBox
}
