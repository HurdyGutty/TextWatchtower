<script lang="ts">
    import { StartOverwatch, StopOverwatch, NewCaptureGroup, DeleteGroup } from "../wailsjs/go/main/App.js";
    import { InstructionAlert, InstructionError, InstructionInfo } from "../wailsjs/go/instruct/InstructionBoard.js";
    type Group = {
        id : number
    }
    const groups: Group[] = [];
    let start = false;
    const numberQueue = [1,2,3,4,5,6]
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
        NewCaptureGroup(id);
    }

    async function deleteGroup(e: Event, id: number) {
        e.preventDefault();
        if (groups.length === 0) {
            InstructionError("There is no more group to delete")
            return;
        }
        numberQueue.push(id)
        let deletedId = await DeleteGroup(id, InstructionError)
        InstructionInfo(`Delete Group ${deletedId}`)
    }

</script>

<div id="group-container">
    {#each groups as group}
        <div class={group.id.toString()}>
            <p>Group {group.id}</p>
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
        </div>
    {/each}
    {#if groups.length < 6}
        <button class="0" on:click={(e) => addNewGroup(e)}>+</button>
    {/if}
</div>