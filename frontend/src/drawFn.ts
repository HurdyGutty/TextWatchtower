import { writable } from 'svelte/store';
import { type Group } from "./stores"

export type DrawFn = (group: Group) => void;
export type ClearFn = () => void;

export const drawFunction = writable((group: Group) => {});

export const clearFunction = writable(() => {});
