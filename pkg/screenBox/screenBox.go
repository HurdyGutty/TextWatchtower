package screenBox

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func DrawBox() (int, int, int, int) {
	var x1, y1, x2, y2, w, h int
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please draw rectangle---")
	if hook.AddMouse("left") {
		fmt.Println("Left")
		x1, y1 = robotgo.Location()
	}

	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		fmt.Printf("Right")
		x2 = int(e.X)
		y2 = int(e.Y)
	})
	w = x2 - x1
	h = y2 - y1

	s := hook.Start()
	<-hook.Process(s)
	return x1, y1, w, h
}
