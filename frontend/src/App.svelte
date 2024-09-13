<script lang="ts">
  import Canvas from "./Canvas.svelte"
  import InstructionBoard from "./InstructionBoard.svelte"
  import GroupContainer from "./GroupContainer.svelte"
  import Title from "./Title.svelte"
  // import Test from "./Test.svelte";
  import { type Group } from "./stores";
  import { onMount } from "svelte";
  import { type DrawFn, type ClearFn ,drawFunction ,clearFunction } from "./drawFn"

  let canvas: Canvas
  let drawCanvasFn: DrawFn
  let clearFn: ClearFn

  onMount(() => {
    drawCanvasFn = canvas.drawCanvas
    clearFn = canvas.clearCanvas
    drawFunction.set(drawCanvasFn)
    clearFunction.set(clearFn)
  })


</script>

<div>
  <Title/>
  <InstructionBoard/>
  <GroupContainer/>
  <Canvas bind:this={canvas}/>
  <!-- <Test drawCanvasFn={drawCanvasFn}/> -->
</div>
<style>
  div {
    display: flex;
    height: 100vh;
    flex-direction: column;
    align-items: center;
    gap: 5px;
  }
</style>