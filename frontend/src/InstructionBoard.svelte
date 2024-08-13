<script>
    import { GetNewInstruct } from "../wailsjs/go/instruct/instructionBoard.js"
    import { onMount, tick } from 'svelte'
    let message = "Welcome to Text Watchtower! Let's start by creating a group"
    let state = "info"
    async function showMessage() {
        await tick();
        GetNewInstruct().then((result) => {
            message = result.message;
            state = result.state;
        })
        .catch((err) => {
            message = "Error in the board. Please wait"
        })
    }
    function StopMessaging() {
        message = "Board stop"
    }

    onMount(async () => {
        await showMessage()
        return StopMessaging
    })
</script>

<div id="board">
    <h3 class={state}>{message}</h3>
</div>

<style>
    #board {
        background-color: aliceblue;
        width: 80%;
        height: 40px;
    }

    .info {
        color: green
    }

    .error {
        color: red
    }

    .alert {
        color: yellow
    }

</style>