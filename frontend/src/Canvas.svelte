<script lang="ts">
    import { onMount } from "svelte";
    import { type Group } from "./stores";

    let canvas: HTMLCanvasElement
    let c: CanvasRenderingContext2D;
    let elementRect: DOMRect
    const radius = 5
    let vw = Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0)
    let vh = Math.max(document.documentElement.clientHeight || 0, window.innerHeight || 0)
    let canvasWidth = 0.92 * vw;
    let canvasHeight = 0.75 * vh;

    
    onMount(() => {
        c = canvas.getContext('2d')
        elementRect = canvas.getBoundingClientRect()
    })

    function clearCanvas() {
        c.clearRect(0, 0, canvasWidth, canvasHeight)
    }

    export function drawCanvas(group: Group) {
        clearCanvas()
        if ("watchBox" in group){
            c.strokeRect(group.watchBox.x - elementRect.left - 1, group.watchBox.y - elementRect.top - 1, group.watchBox.w, group.watchBox.h)
        }
        if ("reload" in group) {
            c.beginPath();
            c.arc(group.reload.X - elementRect.left - 1, group.reload.Y - elementRect.top - 1, radius, 0, 2 * Math.PI, false);
            c.fillStyle = "green";
            c.fill();
        }
    }

</script>


<canvas bind:this={canvas} width={canvasWidth} height={canvasHeight}></canvas>

<style>
    canvas {
        border: 1px solid black;
        margin-top: 3.5rem;
        padding: 0;
    }
</style>