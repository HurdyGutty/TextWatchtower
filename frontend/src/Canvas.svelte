<script lang="ts">
    import { type Group } from "./stores";

    let canvas: HTMLCanvasElement
    let c = canvas.getContext("2d")
    const elementRect = canvas.getBoundingClientRect()

    export function showWatchBox (group: Group) {
        if ("watchBox" in group){
            c.strokeRect(group.watchBox.x - elementRect.left, group.watchBox.y - elementRect.top, group.watchBox.w, group.watchBox.h)
        }
    }

    export function showReloadPoint (group: Group) {
        const radius = 5;
        if ("reload" in group) {
            c.beginPath();
            c.arc(group.reload.X, group.reload.Y, radius, 0, 2 * Math.PI, false);
            c.fillStyle = "green";
            c.fill();
        }
        
    }

    export function clearCanvas() {
        c.clearRect(0, 0, canvas.width, canvas.height)
    }


</script>

<canvas bind:this={canvas}></canvas>

<style>
    canvas {
        border: 1px solid black;

    }
</style>