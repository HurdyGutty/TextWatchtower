<script lang="ts">
    import { AddNewScreenBox, AssignReloadButton, ChangeMax, ChangeMin } from "../wailsjs/go/main/App";
    import { type Group } from "./stores";
    import { InstructionError, InstructionInfo } from "../wailsjs/go/instruct/instructionBoard"
    import Input from "./Input.svelte";

    export let group: Group;
    export let drawCanvas: (group: Group) => void
    let min = group.min.toString();
    let max = group.max.toString();

    function handleMinSubmit(e: Event) {
        e.preventDefault();
        let value = Number(min);
        if (!Number.isInteger(value)) {
            InstructionError("Minimum time must be an integer")
            return
        }
        if (value < 0) {
            InstructionError("Maximum time must be larger than 0")
            return
        }
        if (value > group.max){
            InstructionError("Minimum time must be smaller than maximum time")
            return
        }
        group.min = value;
        ChangeMin(group.id, group.min)
        InstructionInfo(`Changed minimum time for ${group.name}`)
    }
    function handleMaxSubmit(e: Event) {
        e.preventDefault();
        let value = Number(max);
        if (!Number.isInteger(value)) {
            InstructionError("Maximum time must be an integer")
            return
        }
        if (value < 0) {
            InstructionError("Maximum time must be larger than 0")
            return
        }
        if (value < group.min){
            InstructionError("Maximum time must be larger than minimum time")
            return
        }
        group.max = value;
        ChangeMax(group.id, group.max)
        InstructionInfo(`Changed maximum time for ${group.name}`)
    }


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
    <p>
        <span>Minimum time: </span>
        <Input
            id="min"
            type="number"
            bind:value={min}
            on:submit={handleMinSubmit}
    />
        <span>s</span>
    </p>

    <p>
        <span>Maximum time: </span>
        <Input
        id="max"
        type="number"
        bind:value={max}
        on:submit={handleMaxSubmit}
    />
    <span>s</span>
    </p>
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