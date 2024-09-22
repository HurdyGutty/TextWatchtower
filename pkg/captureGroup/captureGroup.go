package captureGroup

import (
	"embed"
	"fmt"
	"math/rand"
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
	Name      string
	Reload    *reloadPoint.MousePoint
	TextBoxes []ScreenBoxText
	Min       int
	Max       int
}

func NewCaptureGroup(id int) *CaptureGroup {
	return &CaptureGroup{Id: id, Name: fmt.Sprintf("Group %d", id), Min: 15, Max: 15}
}

func (group *CaptureGroup) ChangeMin(min int) {
	group.Min = min
}

func (group *CaptureGroup) ChangeMax(max int) {
	group.Max = max
}

func (group *CaptureGroup) ChangeName(name string) {
	group.Name = name
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

func (group *CaptureGroup) getRandomTime() int64 {
	return int64((group.Min + rand.Intn((group.Max - group.Min))) * 1000)
}

func (group *CaptureGroup) Overwatch() chan bool {
	ticker := time.NewTicker(time.Duration(group.getRandomTime()) * time.Millisecond)
	done := make(chan bool, 1)
	groupPath := "group" + strconv.Itoa(group.Id) + "/"
	err := os.MkdirAll("images/"+groupPath, 0750)
	if err != nil {
		InstructionBoard.InstructionError("Can't make new capture group")
	}
	InstructionBoard.InstructionInfo(fmt.Sprintf("%s has started watching", group.Name))
	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				InstructionBoard.InstructionAlert(fmt.Sprintf("%s has stop", group.Name))
				return
			case t := <-ticker.C:
				for i, textBox := range group.TextBoxes {
					imageName := groupPath + "Capture" + strconv.Itoa(i)
					newText := textBox.getNewText(imageName)
					if textBox.isNewText(newText) {
						sent, err := telegram.SendMessage(t, fmt.Sprintf("%s: %s", group.Name, newText))
						if err != nil {
							InstructionBoard.InstructionError("Can't sent telegram messages")
						}
						if sent {
							group.TextBoxes[i].changeText(newText)
							InstructionBoard.InstructionAlert(fmt.Sprintf("%s changes have been sent", group.Name))
						}
					}
				}
				group.Reload.ReloadPage()
			}
		}
	}()
	return done
}
