<script lang="ts">
    import { onMount } from "svelte";
    import { type Group } from "./stores";
    import { GetAppHeight, GetAppWidth } from "../wailsjs/go/main/App";

    let canvas: HTMLCanvasElement
    let c: CanvasRenderingContext2D;
    let elementRect: DOMRect
    const radius = 10

    const scale = window.devicePixelRatio;

    onMount(async () => {
        c = canvas.getContext('2d')
        elementRect = canvas.getBoundingClientRect()
        let appWidth = await GetAppWidth()
        let appHeight = await GetAppHeight()
        let vw = Math.floor(appWidth / scale)
        let vh = Math.floor(appHeight / scale)

        canvas.style.width = `${vw}px`;
        canvas.style.height = `${vh}px`;
        
        canvas.width = appWidth
        canvas.height = appHeight
    })

    export function clearCanvas() {
        c.clearRect(0, 0, canvas.width, canvas.height)
    }

    export function drawCanvas(group: Group) {
        clearCanvas()
        if ("watchBox" in group){
            c.shadowColor = "#d53";
            c.shadowBlur = 8;
            c.lineWidth = 6;
            c.lineJoin = "bevel";
            c.strokeStyle = "#38f";
            c.strokeRect(group.watchBox.x - elementRect.left - 1, group.watchBox.y - elementRect.top - 1, group.watchBox.w, group.watchBox.h)

        }
        if ("reload" in group) {
            c.beginPath();
            c.arc(group.reload.X - elementRect.left - 1, group.reload.Y - elementRect.top - 1, radius, 0, 2 * Math.PI, false);
            c.fillStyle = "rgba(173, 3, 222, 1)";
            c.fill();
        }
    }

</script>

<canvas bind:this={canvas}></canvas>

<style>
    canvas {
        position: absolute;
        top: 0;
        left: 0;
        padding: 0;
        z-index: -1;
    }
</style>