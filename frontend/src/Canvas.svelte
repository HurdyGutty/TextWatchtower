<script lang="ts">
    import { onMount } from "svelte";
    import { type Group } from "./stores";

    let canvas: HTMLCanvasElement
    let c: CanvasRenderingContext2D;
    let elementRect: DOMRect
    const radius = 5
    onMount(() => {
        c = canvas.getContext('2d')
        elementRect = canvas.getBoundingClientRect()
    })

    function clearCanvas() {
        c.clearRect(0, 0, canvas.width, canvas.height)
    }

    export function drawCanvas(group: Group) {
        clearCanvas()
        console.log(elementRect)
        if ("watchBox" in group){
            c.strokeRect(group.watchBox.x - elementRect.left, group.watchBox.y - elementRect.top, group.watchBox.w, group.watchBox.h)
        }
        if ("reload" in group) {
            c.beginPath();
            c.arc(group.reload.X, group.reload.Y, radius, 0, 2 * Math.PI, false);
            c.fillStyle = "green";
            c.fill();
        }
    }

</script>


<canvas bind:this={canvas}></canvas>

<style>
    canvas {
        border: 1px solid black;
        margin-top: 3.5rem;
        width: 80vw;
    }
</style>