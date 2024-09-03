<script lang="ts">
    import { StartOverwatch, StopOverwatch, NewCaptureGroup, DeleteGroup } from "../wailsjs/go/main/App.js";
    import { InstructionAlert, InstructionError, InstructionInfo } from "../wailsjs/go/instruct/instructionBoard.js";
    import { type Group, groupsMap } from "./stores.js"
    import GroupMenu from "./GroupMenu.svelte";
    let startId = 0;
    let numberQueue = [1,2,3,4,5,6]
    let groups: Map<number, Group>;
    export let drawFn: (group: Group) => void;

    groupsMap.subscribe((value) => {
        groups = value;
    })
    function startOverWatch(e: Event, id: number) {
        e.preventDefault();
        StartOverwatch(id);
        startId = id;
    }

    function stopOverwatch(e: Event, id: number) {
        e.preventDefault();
        StopOverwatch(id);
        startId = 0;
    }
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
    }

    async function deleteGroup(e: Event, id: number) {
        e.preventDefault();
        if (groups.size === 0) {
            InstructionError("There is no more group to delete")
            return;
        }
        numberQueue.unshift(id)
        numberQueue = numberQueue
        groups.delete(id)
        groups = groups
        let deletedId = await DeleteGroup(id, InstructionError)
        InstructionInfo(`Delete Group ${deletedId}`)
    }

</script>

<div id="group-container">
    {#each Array.from(groups.values()) as group}
        <div class={"group " + group.id.toString()} 
            on:click={() => drawFn(group)}
            on:keyup={() => drawFn(group)}
            role="button"
            tabindex="0">
            <span>Group {group.id}</span>
            {#if startId !== group.id}
            <button  on:click={(e) => startOverWatch(e, group.id)}>
                <img src="./src/assets/images/play-button.svg" alt="Start"/>
            </button>
            {:else}
            <button on:click={(e) => stopOverwatch(e, group.id)}>
                <img src="./src/assets/images/stop-button.svg" alt="Stop"/>
            </button>
            {/if}
            <button on:click={(e) => deleteGroup(e, group.id)}>
                <img src="./src/assets/images/remove.png" alt ="Delete" />
            </button>
            <GroupMenu group={group} drawCanvas={drawFn}></GroupMenu>
        </div>
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
        width: 20vh;
        display: flex;
        align-items: center;
        position: relative;
        gap: 3px;
        text-align: center;
    }
    .\30 {
        justify-content: center;
    }
    .group span {
        flex-grow: 0.8;
    }
    img {
        max-height: 1.5rem;
    }
    #group-container {
        display: flex;
        flex-direction: row;
        align-items: stretch;
        justify-content: space-around ;
        gap: 8px;
    }
</style>