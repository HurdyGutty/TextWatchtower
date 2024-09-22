import { writable } from 'svelte/store';

export type Group = {
    id: number
    name: string
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
    min: number
    max: number
}

export const groupsMap = writable(new Map<number, Group>());
