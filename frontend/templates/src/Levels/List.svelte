<script lang="ts">
    import {Button, Modal, ModalHeader, Table} from "sveltestrap";
    import type {DTOEntryLevel} from "../dto";
    import {onMount} from "svelte";
    import {isNumber} from "../common";
    import {default as Update} from "./Update.svelte"

    let items: Array<DTOEntryLevel> = []

    let modalIsOpen: boolean = false;
    let modalStatus: string = 'Closed';
    let selectedId: number = -1

    const toggle = (evt) => {
        if (isNumber(evt)) {
            selectedId = evt
        }

        modalIsOpen = !modalIsOpen
    }


    // First page load
    onMount(async () => {

        // Fetch old data if we're updating
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/entries/levels"
        )

        if (!res.ok) {
            return
        }

        items = await res.json()
    })

</script>

<!-- Modal for level -->
<Modal
        body
        isOpen={modalIsOpen}
        size="xl"
        {toggle}
        on:opening={() => (modalStatus = 'Opening...')}
        on:open={() => (modalStatus = 'Opened')}
        on:closing={() => (modalStatus = 'Closing...')}
        on:close={() => (modalStatus = 'Closed')}
>
    <ModalHeader {toggle}>Add/update level</ModalHeader>
    <Update on:submit={toggle} id={selectedId}/>
</Modal>

<Button on:click={toggle}>Add</Button>

<Table class="table-striped">
    <thead>
    <tr>
        <th>Id</th>
        <th>Added</th>
        <th>Name</th>
    </tr>
    </thead>

    <tbody>
    {#each items as item}
        <tr>
            <td class="id" on:click={() => {toggle(item.id)}}>#{item.id}</td>
            <td on:click={() => {toggle(item.id)}}>{item.added_at}</td>
            <td on:click={() => {toggle(item.id)}}>{item.name} ({item.key})</td>
        </tr>
    {/each}
    </tbody>
</Table>

<style>
    td.id {
        width: 2em;
    }
</style>