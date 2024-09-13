<script lang="ts">
    import { instruct } from "../wailsjs/go/models";
    import store from "./instructStore"

    let data: instruct.Instruct

    store.subscribe((value) => data = value)

</script>

<div id="board">
    {#key [data.state, data.message]}
        <h3 class={data.state}>{data.message}</h3>
    {/key}
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