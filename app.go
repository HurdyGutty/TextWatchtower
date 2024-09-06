package main

import (
	"context"

	"github.com/HurdyGutty/go_OCR/pkg/captureGroup"
	"github.com/HurdyGutty/go_OCR/pkg/reloadPoint"
	"github.com/HurdyGutty/go_OCR/pkg/screenBox"
)

// App struct
type App struct {
	ctx context.Context
}

type RandomImage struct {
	Message string
	Status  string
}

type AllBreeds struct {
	Message map[string]map[string][]string
	Status  string
}

type ImagesByBreed struct {
	Message []string
	Status  string
}

type CaptureGroupState struct {
	group    *captureGroup.CaptureGroup
	stopChan chan bool
}

type GroupMap map[int]*CaptureGroupState

type InstructionError func(string)

var Groups = make(GroupMap)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) NewCaptureGroup(id int) {
	newGroup := captureGroup.NewCaptureGroup(id)
	Groups[id] = &CaptureGroupState{
		group:    newGroup,
		stopChan: make(chan bool),
	}
}

func (a *App) AssignReloadButton(id int) *reloadPoint.MousePoint {
	group := Groups[id].group
	reload := reloadPoint.ReloadPoint()
	group.AddReloadPoint(reload)
	return reload
}

func (a *App) AddNewScreenBox(id int) *screenBox.ScreenBox {
	group := Groups[id].group
	box := screenBox.DrawBox()
	group.AddNewTextBox(box)
	return box
}

func (a *App) StartOverwatch(id int) {
	groupState := Groups[id]
	groupState.stopChan = groupState.group.Overwatch()
}

func (a *App) StopOverwatch(id int) {
	groupState := Groups[id]
	groupState.stopChan <- true
}

func (a *App) DeleteGroup(id int, i InstructionError) int {
	deletedId, err := Groups[id].group.DeleteCaptureGroup()
	if err != nil {
		i(err.Error())
	}
	return deletedId
}
