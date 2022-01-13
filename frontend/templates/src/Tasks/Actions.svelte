<script lang="ts">
    import {Button} from "sveltestrap";
    import {isNumber} from "../common";

    export let id: number = -1

    let completeIsOpen: boolean = false;
    const toggleComplete = (id: number) => {
        if (!isNumber(id)) {
            return
        }

        completeIsOpen = !completeIsOpen
    }

    const setAsCompleted = (id: number) => {
        if (!isNumber(id)) {
            return
        }

        console.log(`completing id:` + id)

        toggleComplete(id)
    }

</script>

{#if !completeIsOpen}
    <Button on:click={() => toggleComplete(id)}>Mark as completed</Button>
{:else }
    <Button on:click={() => setAsCompleted(id)}>Complete task</Button>
{/if}

<style>

</style>