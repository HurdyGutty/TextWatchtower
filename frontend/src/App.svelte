<script lang="ts">
  import Canvas from "./Canvas.svelte"
  import InstructionBoard from "./InstructionBoard.svelte"
  import GroupContainer from "./GroupContainer.svelte"
  import Title from "./Title.svelte"
  // import Test from "./Test.svelte";
  import { type Group } from "./stores";
  import { onMount } from "svelte";

  let canvas: Canvas
  let drawCanvasFn: (group: Group) => void

  onMount(() => {
    drawCanvasFn = canvas.drawCanvas
  })

</script>

<div>
  <Title/>
  <InstructionBoard/>
  <GroupContainer drawFn={drawCanvasFn}/>
  <Canvas bind:this={canvas}/>
  <!-- <Test drawCanvasFn={drawCanvasFn}/> -->
</div>
<style>
  div {
    display: flex;
    background-color: rgba(170, 12, 10, 0.4);
    height: 100vh;
    flex-direction: column;
    align-items: center;
    gap: 5px;
  }
</style>