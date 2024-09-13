<script lang="ts">
    import { NewCaptureGroup, DeleteGroup } from "../wailsjs/go/main/App";
    import { InstructionAlert, InstructionError, InstructionInfo } from "../wailsjs/go/instruct/instructionBoard.js";
    import { type Group, groupsMap } from "./stores.js"
    import { clearFunction, type ClearFn } from "./drawFn"
    import WatchGroup from "./WatchGroup.svelte";
    let numberQueue = [1,2,3,4,5,6]
    let groups: Map<number, Group>;

    groupsMap.subscribe((value) => {
        groups = value;
    })

    let clearCanvas: ClearFn

    clearFunction.subscribe((func) => {
        clearCanvas = func
    })

    function addNewGroup(e: Event) {
        e.preventDefault();
        if (numberQueue.length === 0) {
            InstructionAlert("Maximum of 6 groups only")
            return;
        }
        let id = numberQueue.shift()
        numberQueue = numberQueue
        NewCaptureGroup(id);
        groups = groups.set(id, {id})
        clearCanvas()
        InstructionInfo("Added a new group")
    }

    async function deleteGroup(e: Event, id: number) {
        e.preventDefault();
        if (groups.size === 0) {
            InstructionError("There is no more group to delete")
            return;
        }
        let deletedId = await DeleteGroup(id)
        if (deletedId === 0) {
            return 
        }
        numberQueue.unshift(id)
        numberQueue = numberQueue
        groups.delete(id)
        groups = groups
        clearCanvas()
        InstructionInfo(`Delete Group ${deletedId}`)
    }

</script>

<div id="group-container">
    {#each Array.from(groups.values()) as group}
        <WatchGroup group={group} deleteGroup={deleteGroup}></WatchGroup>
    {/each}
    {#if groups.size < 6}
        <button class="group 0" on:click={(e) => addNewGroup(e)}>+</button>
    {/if}
</div>

<style>
    button {
        border: none;
        background: none;
        cursor: pointer;
        margin: 0;
        padding: 0;
        font-size: medium;
        color: black;
    }
    .group {
        background-color: aliceblue;
        border-radius: 16px;
        padding: 8px;
        width: calc(20vh + 6px);
        min-height: 44px;
        display: flex;
        align-items: center;
        position: relative;
        gap: 3px;
        text-align: center;
        justify-content: center;
        font-weight: bold;
    }
    #group-container {
        display: flex;
        flex-direction: row;
        align-items: stretch;
        justify-content: space-around ;
        gap: 8px;
    }
</style>