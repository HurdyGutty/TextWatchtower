<script lang="ts">
    import GroupMenu from "./GroupMenu.svelte";
    import { groupsMap, type Group } from "./stores"
    export let drawCanvasFn: (group: Group) => void

    let groups: Map<number, Group>

    groupsMap.subscribe((value) => {
        groups = value
    })

    let group = groups.get(1)
    
</script>

{#if typeof group === "undefined"}
<div>1</div>
{:else}
<div on:click={() => drawCanvasFn(group)}
    on:keyup={() => drawCanvasFn(group)}
    role="button"
    tabindex="0">{group.id}</div>
<GroupMenu group={group} drawCanvas={drawCanvasFn}/>
{/if}