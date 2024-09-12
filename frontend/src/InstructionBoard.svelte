<script lang="ts">
    import { instruct } from "../wailsjs/go/models";
    import store from "./instructStore"

    let data: Promise<instruct.Instruct>

    store.subscribe((value) => data = value)

</script>

<div id="board">
    {#await data}
        <h3 class="info">Waiting</h3>
    {:then data}
        {#key [data.state, data.message]}
            <h3 class={data.state}>{data.message}</h3>
        {/key}
    {:catch error}
        <h3 class="error">Error in the board. Please wait</h3>
    {/await}
</div>

<style>
    #board {
        background-color: aliceblue;
        height: 40px;
        text-align: center;
        padding: 8px;
        border-radius: 4px;
    }

    h3 {
        line-height: 40px
    }

    .info {
        color: green
    }

    .error {
        color: red
    }

    .alert {
        color: orange
    }

</style>