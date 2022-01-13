<script lang="ts">
    import {Button, Modal, ModalHeader, Table} from 'sveltestrap'
    import {onMount} from "svelte";
    import {isNumber} from "../common";
    import {DTOTag} from "../dto";
    import {default as Update} from "./Update.svelte";

    let selectedId: number = -1
    let items: Array<DTOTag> = []

    let modalIsOpen: boolean = false;
    let modalStatus = 'Closed';

    const toggle = (id: number) => {
        if (isNumber(id)) {
            selectedId = id
        }

        modalIsOpen = !modalIsOpen
    }

    async function fetchTasks() {
        // Load tasks
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/tags"
        )

        if (!res.ok) {
            return
        }

        items = await res.json()

    }

    async function update() {
        await fetchTasks()
    }

    // Load on first visit
    onMount(async () => {
        await update()
    })

</script>

<!-- Modal for adding/updating tag form -->
<Modal
        body
        isOpen={modalIsOpen}
        size="xl"
        {toggle}
        on:opening={() => (modalStatus = 'Opening...')}
        on:open={() => (modalStatus = 'Opened')}
        on:closing={() => (modalStatus = 'Closing...')}
        on:close={() => (update())}
>
    <ModalHeader {toggle}>Edit tag</ModalHeader>
    <Update on:submit={toggle} id={selectedId}/>
</Modal>

<h1>Tags</h1>

<Button on:click={() => toggle(-1)}>Add tag</Button>

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
            <td on:click={() => toggle(item.id)}># {item.id}</td>
            <td class="date">{item.added_at}</td>
            <td on:click={() => toggle(item.id)}>{item.name} ({item.shortname})</td>
        </tr>
    {/each}

    </tbody>

</Table>

<style>
    td.date {
        width: 4em;
    }
</style>
