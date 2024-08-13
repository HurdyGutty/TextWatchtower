package screenshot

import (
	"github.com/go-vgo/robotgo"
)

func CreateImagePath(imgName string) string {
	return "images/" + imgName + ".png"
}

func TakeScreenShot(x1, y1, w, h int, imgPath string) {
	bit := robotgo.CaptureScreen(x1, y1, w, h)
	defer robotgo.FreeBitmap(bit)

	img := robotgo.ToImage(bit)
	robotgo.Save(img, imgPath)
}
