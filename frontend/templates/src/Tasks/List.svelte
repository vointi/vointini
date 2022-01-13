<script lang="ts">
    import {isNumber} from "../common";
    import {Button, Table} from 'sveltestrap'
    import {Modal, ModalHeader} from 'sveltestrap';
    import {onMount} from "svelte";
    import {DTOTask} from "../dto";
    import {default as TaskAdd} from "./Add.svelte";
    import {default as TaskActions} from "./Actions.svelte";
    import {default as TaskUpdate} from "./Update.svelte";

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
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/tasks"
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
        <TaskAdd on:submit={toggle}/>
    {:else }
        <ModalHeader {toggle}>Edit task</ModalHeader>
        <TaskUpdate on:submit={toggle} id={selectedId}/>
    {/if}
</Modal>

<h1>Tasks</h1>

<Button on:click={() => toggle(-1)}>Add task</Button>

<!--
<Button on:click={() => toggle(-1)}>Add recurring task</Button>
-->
<Table class="table-striped">
    <thead>
    <tr>
        <th>Id</th>
        <th>Added</th>
        <th>Completed</th>
        <th>Title</th>
        <th>Description</th>
        <th>Action</th>
    </tr>
    </thead>

    <tbody>

    {#each items as item}
        <tr>
            <td on:click={() => toggle(item.id)}># {item.id}</td>
            <td class="date">{item.added_at}</td>
            <td class="date">
                {#if item.completed_at !== null}
                    {item.completed_at}
                {/if}
            </td>
            <td on:click={() => toggle(item.id)}>{item.title}</td>
            <td>{item.description}</td>
            <td class="actions">
                {#if item.completed_at === null}
                    <TaskActions id={item.id}/>
                {/if}
            </td>
        </tr>
    {/each}

    </tbody>

</Table>

<style>
    td.date {
        width: 4em;
    }
</style>
