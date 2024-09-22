<script lang="ts">
    import { type Group } from "./stores";

    export let toggleRename: svelte.JSX.EventHandler<Event>
    export let group: Group

</script>

<p>
    <span on:click={toggleRename}
    on:keyup={toggleRename}
    role="button"
    tabindex="0">
    {group.name}
    </span>
</p>

<style>
    p {
        flex-grow: 0.6;
        overflow-x: hidden;
        white-space: nowrap;
        box-sizing: border-box;
        position: relative;
        max-width: 4rem;
    }
    span{
        display: inline-block;
        position: relative;
        font-weight: bold;
    }
    span:hover{
        animation: move 5s ease-out infinite;
    }
    @keyframes move {
        25% { transform: translate(0, 0); left: 0%; }
        75%, 100% { transform: translate(-100%, 0); left: 100%; }
    }
</style>