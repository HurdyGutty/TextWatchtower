<script lang="ts">
    import { instruct } from "../wailsjs/go/models";
    import { GetNewInstruct } from "../wailsjs/go/instruct/instructionBoard"


    async function fetchUpdateBoard() {
        const res = await GetNewInstruct();
        console.log(res)
        if (typeof res !== "undefined") {
            return res
        } else {
            return {
                message: "Error in the board. Please wait",
                state: "alert",
            } as instruct.Instruct
        }
    }

    let data = fetchUpdateBoard()

    export function updateInstructionBoard() {
        data = fetchUpdateBoard()
    }

</script>

<div id="board">
    {#await data}
        <h3 class="info">Waiting</h3>
    {:then data}
        <h3 class={data.state}>{data.message}</h3>
    {:catch error}
        <h3 class="error">Error in the board. Please wait</h3>
    {/await}
</div>

<style>
    #board {
        background-color: aliceblue;
        height: 40px;
        text-align: center;
        padding: 8px;
        border-radius: 4px;
    }

    h3 {
        line-height: 40px
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