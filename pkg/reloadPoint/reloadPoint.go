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

	fmt.Println("--- Please press Esc to stop hook ---")
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("Escaped")
		hook.End()
	})
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		point.X, point.Y = robotgo.Location()
	})
	s := hook.Start()
	<-hook.Process(s)
	return point
}
