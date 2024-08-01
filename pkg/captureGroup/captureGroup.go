package captureGroup

import (
	"github.com/HurdyGutty/go_OCR/pkg/reloadPoint"
	"github.com/HurdyGutty/go_OCR/pkg/screenBox"
)

type ScreenBoxText struct {
	ScreenBox *screenBox.ScreenBox
	Text      string
}

type CaptureGroup struct {
	Reload    *reloadPoint.MousePoint
	TextBoxes []ScreenBoxText
}

func NewCaptureGroup() *CaptureGroup {
	return new(CaptureGroup)
}
