import { writable } from 'svelte/store';
import { ReceiveInstruct } from "../wailsjs/go/instruct/instructionBoard"
import { instruct } from '../wailsjs/go/models';

let instructStruct: instruct.Instruct = {
    message: "Welcome to Text Watchtower! Let's start by creating a group",
    state: "info"
}

const messages = writable(instructStruct);
	const evtSource = new EventSource("http://localhost:12345/events");
	evtSource.onmessage = function(event) {
		console.log(event);
		var dataobj = JSON.parse(event.data) as instruct.Instruct;
		messages.set(dataobj);
	}

export default {
	subscribe: messages.subscribe,
}

