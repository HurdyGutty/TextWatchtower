import { writable } from 'svelte/store';
import { ReceiveInstruct } from "../wailsjs/go/instruct/instructionBoard"
import { instruct } from '../wailsjs/go/models';

let instructStruct: Promise<instruct.Instruct>

function fetchInstruct() {
    instructStruct = ReceiveInstruct()
}

fetchInstruct()

const messageStore = writable(instructStruct, () => {
    let interval = setInterval(() => {
        messageStore.update((value) => (value = ReceiveInstruct()))
    }, 2000)

    return () => {
        clearInterval(interval)
    }
});

export default {
	subscribe: messageStore.subscribe,
}

