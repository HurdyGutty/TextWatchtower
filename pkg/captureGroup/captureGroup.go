package captureGroup

import (
	"embed"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/HurdyGutty/go_OCR/pkg/OCR"
	"github.com/HurdyGutty/go_OCR/pkg/instruct"
	"github.com/HurdyGutty/go_OCR/pkg/reloadPoint"
	"github.com/HurdyGutty/go_OCR/pkg/screenBox"
	"github.com/HurdyGutty/go_OCR/pkg/screenshot"
	"github.com/HurdyGutty/go_OCR/pkg/telegram"
)

var TrainingData embed.FS
var InstructionBoard *instruct.InstructionBoard

func trimText(text string) string {
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	} else {
		text = strings.TrimRight(text, "\n")
	}
	return text
}

type ScreenBoxText struct {
	ScreenBox *screenBox.ScreenBox
	Text      string
}

func (textBox *ScreenBoxText) getNewText(imageName string) string {
	newImagePath := screenshot.CreateImagePath(imageName)
	screenshot.TakeScreenShot(textBox.ScreenBox.X1, textBox.ScreenBox.Y1, textBox.ScreenBox.W, textBox.ScreenBox.H, newImagePath)
	fmt.Printf("%+v\n", TrainingData)
	newText := OCR.ProcessOCR(TrainingData, newImagePath)
	newText = trimText(newText)
	return newText
}

func (textBox *ScreenBoxText) isNewText(newText string) bool {
	newText = trimText(newText)
	fmt.Printf("### New text ### \n %s \n ### Old text ### \n %s \n", newText, textBox.Text)

	return strings.Compare(newText, textBox.Text) != 0

}

func (textBox *ScreenBoxText) changeText(newText string) {
	textBox.Text = newText
}

type CaptureGroup struct {
	Id        int
	Reload    *reloadPoint.MousePoint
	TextBoxes []ScreenBoxText
}

func NewCaptureGroup(id int) *CaptureGroup {
	return &CaptureGroup{Id: id}
}

func (group *CaptureGroup) AddReloadPoint(mousePoint *reloadPoint.MousePoint) {
	group.Reload = mousePoint
}

func (group *CaptureGroup) AddNewTextBox(screenBox *screenBox.ScreenBox) {
	newTextBox := ScreenBoxText{
		ScreenBox: screenBox,
		Text:      "",
	}
	group.TextBoxes = append(group.TextBoxes, newTextBox)
}

func (group *CaptureGroup) DeleteCaptureGroup() (int, error) {
	groupPath := "images/group" + strconv.Itoa(group.Id) + "/"
	err := os.RemoveAll(groupPath)
	if err != nil {
		return 0, err
	}
	return group.Id, nil
}

func (group *CaptureGroup) Overwatch() chan bool {
	ticker := time.NewTicker(15000 * time.Millisecond)
	done := make(chan bool, 1)
	groupPath := "group" + strconv.Itoa(group.Id) + "/"
	err := os.MkdirAll("images/"+groupPath, 0750)
	if err != nil {
		InstructionBoard.InstructionError("Can't make new capture group")
	}

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				fmt.Printf("Group %d has stop\n", group.Id)
				return
			case t := <-ticker.C:
				for i, textBox := range group.TextBoxes {
					imageName := groupPath + "Capture" + strconv.Itoa(i)
					newText := textBox.getNewText(imageName)
					if textBox.isNewText(newText) {
						sent, err := telegram.SendMessage(t, newText)
						if err != nil {
							InstructionBoard.InstructionError("Can't sent telegram messages")
						}
						if sent {
							group.TextBoxes[i].changeText(newText)
							fmt.Println("Sent")
						}
					}
				}
				group.Reload.ReloadPage()
			}
		}
	}()
	return done
}
