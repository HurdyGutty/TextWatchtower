<script lang="ts">
  import Canvas from "./Canvas.svelte"
  import InstructionBoard from "./InstructionBoard.svelte"
  import GroupContainer from "./GroupContainer.svelte"
  import Title from "./Title.svelte"
  // import Test from "./Test.svelte";
  import { type Group } from "./stores";
  import { onMount } from "svelte";
  import { drawFunction } from "./drawFn"

  let canvas: Canvas
  let board: InstructionBoard
  let drawCanvasFn: (group: Group) => void
  let updateBoardFn: () => void

  onMount(() => {
    drawCanvasFn = canvas.drawCanvas
    drawFunction.set(drawCanvasFn)
    updateBoardFn = board.updateInstructionBoard
  })


</script>

<div>
  <Title/>
  <InstructionBoard bind:this={board}/>
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