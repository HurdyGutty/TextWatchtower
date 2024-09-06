import { writable } from 'svelte/store';
import { type Group } from "./stores"

export type DrawFn = (group: Group) => void;

export const drawFunction = writable((group: Group) => {});
