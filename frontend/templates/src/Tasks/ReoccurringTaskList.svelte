<script lang="ts">
    import {Button, Table} from 'sveltestrap'
    import {Modal, ModalHeader} from 'sveltestrap';
    import {onMount} from "svelte";
    import {DTOTask} from "../dto";
    import {isNumber} from "../common";
    import {default as ReoccurringTaskAdd} from "./ReoccurringTaskAdd.svelte";

    let selectedId: number = -1
    let items: Array<DTOTask> = []

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
        const res = await fetch(
            "/api/v1/reoccurring-tasks"
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

<!-- Modal for adding/updating entry form -->
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
    {#if selectedId === -1}
        <ModalHeader {toggle}>Add task</ModalHeader>
        <ReoccurringTaskAdd on:submit={toggle}/>
    {:else }
        <ModalHeader {toggle}>Edit task</ModalHeader>
        <ReoccurringTaskUpdate on:submit={toggle} id={selectedId}/>
    {/if}
</Modal>

<h1>Re-occurring Tasks</h1>

<Button on:click={() => toggle(-1)}>Add reoccurring task</Button>

<Table class="table-striped">
    <thead>
    <tr>
        <th>Id</th>
        <th>Added</th>
        <th>Title</th>
    </tr>
    </thead>

    <tbody>

    {#each items as item}
        <tr>
            <td on:click={() => toggle(item.id)}># {item.id}</td>
            <td class="date">{item.added_at}</td>
            <td on:click={() => toggle(item.id)}>{item.title}</td>
        </tr>
    {/each}

    </tbody>

</Table>

<style>
    td.date {
        width: 4em;
    }
</style>
