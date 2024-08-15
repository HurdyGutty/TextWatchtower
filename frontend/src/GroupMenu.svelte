<script lang="ts">
    import { AddNewScreenBox, AssignReloadButton } from "wailsjs/go/main/App";
    import { Group } from "./stores";
    import { InstructionInfo } from "../wailsjs/go/instruct/instructionBoard.js"

    export let group: Group;

    async function changeReloadButton(e: Event) {
        e.preventDefault()
        InstructionInfo("Please click at the button then press Esc when done")
        group.reload = await AssignReloadButton(group.id)
    }

    async function changeNewWatchBox(e: Event) {
        e.preventDefault()
        InstructionInfo("Draw a rectangle in the box below then press Esc when done")
        group.watchBox = await AddNewScreenBox(group.id)
    }
</script>

<div class="group-menu">
    <button on:click={(e) => changeNewWatchBox(e)}>
        {#if "watchBox" in group}
        Change this group's watch zone
        {:else}
        Add new watch zone
        {/if}
    </button>

    <button on:click={(e) => changeReloadButton(e)}>
        {#if "reload" in group}
        Change this groups's reload button
        {:else}
        Add new reload button
        {/if}
    </button>
</div>

<style>
    div {
        display: flex;
        flex-direction: column;
        gap: 2px;
    }
    button {
        all: unset;
    }
</style>