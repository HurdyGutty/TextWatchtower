<script lang="ts">
    import { AddNewScreenBox, AssignReloadButton } from "../wailsjs/go/main/App";
    import { type Group } from "./stores";
    import { InstructionInfo } from "../wailsjs/go/instruct/instructionBoard"

    export let group: Group;
    export let drawCanvas: (group: Group) => void


    async function changeReloadButton(e: Event) {
        e.preventDefault()
        if ("id" in group){
            InstructionInfo("Please click at the button then press Esc when done")
            group.reload = await AssignReloadButton(group.id)
            drawCanvas(group)
            InstructionInfo("Reload button has changed")
        }
    }

    async function changeNewWatchBox(e: Event) {
        e.preventDefault()
        if ("id" in group) {
            InstructionInfo("Draw a rectangle in the box below then press Esc when done")
            group.watchBox = await AddNewScreenBox(group.id)
            drawCanvas(group)
            InstructionInfo("Changed watch window for this group")
        }
    }
</script>

<div class="group-menu">
    <button on:click={(e) => changeNewWatchBox(e)}>
        {#if "watchBox" in group}
        Change watch zone
        {:else}
        Add new watch zone
        {/if}
    </button>

    <button on:click={(e) => changeReloadButton(e)}>
        {#if "reload" in group}
        Change reload button
        {:else}
        Add new reload button
        {/if}
    </button>
</div>

<style>
    div.group-menu{
        display: flex;
        flex-direction: column;
        gap: 2px;
        position: absolute;
        top: 100%;
        background-color: grey;
        align-items: start;
        border-radius: 6px;
        padding: 3px;
    }
    button {
        all: unset;
        cursor: pointer;
    }
</style>