package sse

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HurdyGutty/go_OCR/pkg/instruct"
)

var InstructBoard *instruct.InstructionBoard

func Serve() {
	http.HandleFunc("/events", eventsHandler)
	http.ListenAndServe(":12345", nil)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	fmt.Println("SSE connected")

	instruct, ok := <-InstructBoard.Channel
	if ok {
		data, err := json.Marshal(instruct)
		if err != nil {
			InstructBoard.InstructionError("Error passing JSON  of the board")
		}
		fmt.Fprintf(w, "data: %s\n\n", data)
		w.(http.Flusher).Flush()
	}
}
