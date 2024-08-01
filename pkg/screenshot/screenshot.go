package screenshot

import (
	"github.com/go-vgo/robotgo"
)

func TakeScreenShot(x1, y1, w, h int, imgName string) {
	bit := robotgo.CaptureScreen(x1, y1, w, h)
	defer robotgo.FreeBitmap(bit)

	img := robotgo.ToImage(bit)
	robotgo.Save(img, "images/"+imgName+".png")
}
