package screenshot

import (
	"github.com/go-vgo/robotgo"
)

func TakeScreenShot(x1, x2, y1, y2 int, imgName string) {
	bit := robotgo.CaptureScreen(x1, y1, x2, y2)
	defer robotgo.FreeBitmap(bit)

	img := robotgo.ToImage(bit)
	robotgo.Save(img, "images/"+imgName+".png")
}
