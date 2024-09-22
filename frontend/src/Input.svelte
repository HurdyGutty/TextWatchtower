<script lang="ts">
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();
    export let id = '';
	export let type = 'text';
	export let placeholder = '';
	export let value = '';

    const handleInput = ({ target }) => {
		value = target.value;
	};

	const handleKeypress = (e) => {
		if (e.charCode === 13) {
            e.target.blur();
        };
	};
    const handleBlur = (e) => {
        dispatch('submit')
    }


</script>

<input
		{id}
		{type}
		{placeholder}
		{value}
		on:keypress={handleKeypress}
		on:input={handleInput}
        on:blur={handleBlur}
/>

<style>
    /* Chrome, Safari, Edge, Opera */
    input::-webkit-outer-spin-button,
    input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
    }

    /* Firefox */
    input[type=number] {
    -moz-appearance: textfield;
    }

    #name{
        max-width: 4rem;
        font-weight: bold;
    }

    #min{
        max-width: 20px;
        border-radius: 50%;
        text-align: center;
        border: none;
    }
    #max{
        max-width: 20px;
        border-radius: 50%;
        text-align: center;
        border: none;
    }

</style>

