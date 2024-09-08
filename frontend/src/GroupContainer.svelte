<script lang="ts">
    import { NewCaptureGroup, DeleteGroup } from "../wailsjs/go/main/App.js";
    import { InstructionAlert, InstructionError, InstructionInfo } from "../wailsjs/go/instruct/instructionBoard.js";
    import { type Group, groupsMap } from "./stores.js"
    import WatchGroup from "./WatchGroup.svelte";
    import { updateBoard, type UpdateBoardFn } from "./instructStore.js";
    let numberQueue = [1,2,3,4,5,6]
    let groups: Map<number, Group>;

    groupsMap.subscribe((value) => {
        groups = value;
    })

    let updateInstructBoard: UpdateBoardFn

    updateBoard.subscribe((value) => {
        updateInstructBoard = value
    })

    function addNewGroup(e: Event) {
        e.preventDefault();
        if (numberQueue.length === 0) {
            InstructionAlert("Maximum of 6 groups only")
            updateInstructBoard()
            return;
        }
        let id = numberQueue.shift()
        numberQueue = numberQueue
        NewCaptureGroup(id);
        groups = groups.set(id, {id})
        InstructionInfo("Added a new group")
        updateInstructBoard()
    }

    async function deleteGroup(e: Event, id: number) {
        e.preventDefault();
        if (groups.size === 0) {
            InstructionError("There is no more group to delete")
            updateInstructBoard()
            return;
        }
        numberQueue.unshift(id)
        numberQueue = numberQueue
        groups.delete(id)
        groups = groups
        let deletedId = await DeleteGroup(id, InstructionError)
        InstructionInfo(`Delete Group ${deletedId}`)
        updateInstructBoard()
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