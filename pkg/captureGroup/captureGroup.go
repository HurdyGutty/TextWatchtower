package captureGroup

import (
	"embed"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/HurdyGutty/go_OCR/pkg/OCR"
	"github.com/HurdyGutty/go_OCR/pkg/reloadPoint"
	"github.com/HurdyGutty/go_OCR/pkg/screenBox"
	"github.com/HurdyGutty/go_OCR/pkg/screenshot"
	"github.com/HurdyGutty/go_OCR/pkg/telegram"
)

var TrainingData embed.FS

type ScreenBoxText struct {
	ScreenBox *screenBox.ScreenBox
	Text      string
}

func (textBox *ScreenBoxText) getNewText(imageName string) string {
	newImagePath := screenshot.CreateImagePath(imageName)
	screenshot.TakeScreenShot(textBox.ScreenBox.X1, textBox.ScreenBox.Y1, textBox.ScreenBox.W, textBox.ScreenBox.H, newImagePath)
	fmt.Printf("%+v\n", TrainingData)
	newText := OCR.ProcessOCR(TrainingData, newImagePath)
	return newText
}

func (textBox *ScreenBoxText) isNewText(newText string) bool {
	if runtime.GOOS == "windows" {
		newText = strings.TrimRight(newText, "\r\n")
	} else {
		newText = strings.TrimRight(newText, "\n")
	}
	fmt.Printf("New text: %s Old text: %s", newText, textBox.Text)

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
	stop := make(chan bool)
	ticker := time.NewTicker(15000 * time.Millisecond)
	done := make(chan bool)
	groupPath := "group" + strconv.Itoa(group.Id) + "/"
	err := os.MkdirAll("images/"+groupPath, 0750)
	if err != nil {
		log.Fatalf("Can't make capture group %d", group.Id)
	}

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				for i, textBox := range group.TextBoxes {
					imageName := groupPath + "Capture" + strconv.Itoa(i)
					newText := textBox.getNewText(imageName)
					if textBox.isNewText(newText) {
						sent, err := telegram.SendMessage(t, newText)
						if err != nil {
							log.Fatal(err)
						}
						if sent {
							textBox.changeText(newText)
						}
					}
				}
			}
		}
	}()

	if <-stop {
		ticker.Stop()
		done <- true
		fmt.Printf("Group %d has stop/n", group.Id)
	}
	return stop
}
