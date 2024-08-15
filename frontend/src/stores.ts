import { writable } from 'svelte/store';

export type Group = {
    id: number
    reload?: {
        X: number,
        Y: number,
    }
    watchBox?: {
        x: number,
        y: number,
        w: number,
        h: number,
    }
}

export const groupsMap = writable(new Map<number, Group>());
