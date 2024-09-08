import { writable } from 'svelte/store';

export type UpdateBoardFn = () => void;

export const updateBoard = writable(() => {});