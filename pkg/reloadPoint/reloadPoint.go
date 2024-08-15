package reloadPoint

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type MousePoint struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

func ReloadPoint() *MousePoint {
	point := new(MousePoint)
	s := hook.Start()
	fmt.Println("--- Please press Esc to stop hook ---")
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("Escaped")
		hook.End()
	})
	if hook.AddMouse("left") {
		point.X, point.Y = robotgo.Location()
	}
	<-hook.Process(s)
	return point
}
