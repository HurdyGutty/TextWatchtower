<script lang="ts">
    import { StartOverwatch, StopOverwatch, NewCaptureGroup, DeleteGroup } from "../wailsjs/go/main/App.js";
    import { InstructionAlert, InstructionError, InstructionInfo } from "../wailsjs/go/instruct/InstructionBoard";
    import { type Group, groupsMap } from "./stores"
    import GroupMenu from "./GroupMenu.svelte";
    let start = false;
    let numberQueue = [1,2,3,4,5,6]
    let groups: Map<number, Group>;
    groupsMap.subscribe((value) => {
        groups = value;
    })
    function startOverWatch(e: Event, id: number) {
        e.preventDefault();
        StartOverwatch(id);
        start = true;
    }

    function stopOverwatch(e: Event, id: number) {
        e.preventDefault();
        StopOverwatch(id);
        start = false;
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
        numberQueue.push(id)
        numberQueue = numberQueue
        groups.delete(id)
        groups = groups
        let deletedId = await DeleteGroup(id, InstructionError)
        InstructionInfo(`Delete Group ${deletedId}`)
    }

</script>

<div id="group-container">
    {#each Array.from(groups.values()) as group}
        <div class={group.id.toString()}>
            <span>Group {group.id}</span>
            {#if !start}
            <button  on:click={(e) => startOverWatch(e, group.id)}>
                <img src="assets/images/play-button.svg" alt="Start"/>
            </button>
            {:else}
            <button on:click={(e) => stopOverwatch(e, group.id)}>
                <img src="assets/images/stop-button.svg" alt="Start"/>
            </button>
            {/if}
            <button on:click={(e) => deleteGroup(e, group.id)}>
                <img src="assets/images/delete-button.svg" alt ="Delete" />
            </button>
            <GroupMenu group={group}></GroupMenu>
        </div>
    {/each}
    {#if groups.size < 6}
        <button class="0" on:click={(e) => addNewGroup(e)}>+</button>
    {/if}
</div>