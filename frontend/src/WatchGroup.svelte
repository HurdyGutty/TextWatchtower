<script lang="ts">
    import { type Group } from "./stores";
    import GroupMenu from "./GroupMenu.svelte";
    import { type DrawFn, drawFunction } from "./drawFn";
    import { ChangeName, StartOverwatch, StopOverwatch } from "../wailsjs/go/main/App";
    import { createEventDispatcher } from 'svelte';
    import Input from "./Input.svelte";
    import { InstructionAlert } from "../wailsjs/go/instruct/instructionBoard";
    import GroupName from "./GroupName.svelte";

    const dispatch = createEventDispatcher();

    export let group: Group
    export let deleteGroup: (e: Event, id: number) => void;
    let drawFn: DrawFn;
    let rename = false;
    let name = group.name;

    drawFunction.subscribe((value) => {
        drawFn = value;
    })
    let play = true;
    let showMenu = true;

    function toggleRename() {
        rename = !rename
    }

    function handleNameSubmit(e: Event) {
        e.preventDefault();
        group.name = name;
        ChangeName(group.id, name)
        toggleRename()
        InstructionAlert(`Changed name to ${group.name}`)
    }

    function togglePlay() {
        play = !play
    }

    function startOverWatch(e: Event, id: number) {
        e.preventDefault();
        StartOverwatch(id);
    }

    function stopOverwatch(e: Event, id: number) {
        e.preventDefault();
        StopOverwatch(id);
    }

    function toggleMenu() {
        showMenu = !showMenu
    }

</script>

<div class={"group " + group.id.toString()} 
    on:click={() => drawFn(group)}
    on:keyup={() => drawFn(group)}
    role="button"
    tabindex="0">
        {#if !rename}
        <GroupName {group} {toggleRename}/>
        {:else}
        <Input 
            id="name"
            type="text" 
            placeholder={name} 
            bind:value={name} 
            on:submit={handleNameSubmit}
        />
        {/if}

        

        <button on:click={toggleMenu}>
            <img src="/assets/images/menu.svg" alt="Menu"/>
        </button>

        {#if play}
        <button  on:click={(e) => {
            startOverWatch(e, group.id);
            togglePlay();
            }}>
            <img src="/assets/images/play-button.svg" alt="Start"/>
        </button>
        {:else}
        <button on:click={(e) => {
            stopOverwatch(e, group.id);
            togglePlay();
        }}>
            <img src="/assets/images/stop-button.svg" alt="Stop"/>
        </button>
        {/if}
        <button on:click={(e) => deleteGroup(e, group.id)}>
            <img src="/assets/images/remove.png" alt ="Delete" />
        </button>
        {#if showMenu}
        <GroupMenu group={group} drawCanvas={drawFn}></GroupMenu>
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
        display: flex;
        align-items: center;
        position: relative;
        text-align: center;
        justify-content: space-evenly;
    }
    img {
        max-height: 1.5rem;
    }

</style>