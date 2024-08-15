package main

import (
	"context"
	"fmt"
	"log"

	"io"
	"net/http"

	"encoding/json"
	"sort"

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

var Groups GroupMap

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
	Groups[id].group = newGroup
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

func (a *App) GetBreedList() []string {
	var breeds []string

	response, err := http.Get("https://dog.ceo/api/breeds/list/all")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data AllBreeds
	json.Unmarshal(responseData, &data)

	for k := range data.Message {
		breeds = append(breeds, k)
	}

	sort.Strings(breeds)

	return breeds
}

func (a *App) GetRandomImageUrl() string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data RandomImage
	json.Unmarshal(responseData, &data)

	return data.Message
}

func (a *App) GetImageUrlsByBreed(breed string) []string {

	url := fmt.Sprintf("%s%s%s%s", "https://dog.ceo/api/", "breed/", breed, "/images")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data ImagesByBreed
	json.Unmarshal(responseData, &data)

	return data.Message
}
