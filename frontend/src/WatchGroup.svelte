<script lang="ts">
    import { type Group } from "./stores";
    import GroupMenu from "./GroupMenu.svelte";
    import { type DrawFn, drawFunction } from "./drawFn";
    import { StartOverwatch, StopOverwatch } from "../wailsjs/go/main/App";

    export let group: Group
    export let deleteGroup: (e: Event, id: number) => void;
    let drawFn: DrawFn;

    drawFunction.subscribe((value) => {
        drawFn = value;
    })
    let play = true;
    let showMenu = true;

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
        <span>Group {group.id}</span>
        <button on:click={toggleMenu}>
            <img src="./src/assets/images/menu.svg" alt="Menu"/>
        </button>

        {#if play}
        <button  on:click={(e) => {
            startOverWatch(e, group.id);
            togglePlay();
            }}>
            <img src="./src/assets/images/play-button.svg" alt="Start"/>
        </button>
        {:else}
        <button on:click={(e) => {
            stopOverwatch(e, group.id);
            togglePlay();
        }}>
            <img src="./src/assets/images/stop-button.svg" alt="Stop"/>
        </button>
        {/if}
        <button on:click={(e) => deleteGroup(e, group.id)}>
            <img src="./src/assets/images/remove.png" alt ="Delete" />
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
    .group span {
        flex-grow: 0.6;
        font-weight: bold;
    }
    img {
        max-height: 1.5rem;
    }
</style>