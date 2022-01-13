<script lang="ts">
    import {Button, Modal, ModalHeader, Table} from "sveltestrap";
    import type {DTOHeight} from "../dto";
    import {onMount} from "svelte";
    import {default as HeightAdd} from "./Add.svelte"

    let items: Array<DTOHeight> = []

    let modalIsOpen: boolean = false;
    let modalStatus:string = 'Closed';

    const toggle = (evt) => {
        modalIsOpen = !modalIsOpen
    }


    // First page load
    onMount(async () => {

        // Fetch old data if we're updating
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/height"
        )

        if (!res.ok) {
            return
        }

        items = await res.json()
    })

</script>

<!-- Modal for adding height form -->
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
    <ModalHeader {toggle}>Add weight</ModalHeader>
    <HeightAdd on:submit={toggle}/>
</Modal>

<Button on:click={toggle}>Add</Button>

<Table class="table-striped">
    <thead>
    <tr>
        <th>Id</th>
        <th>Added</th>
        <th>Height</th>
    </tr>
    </thead>

    <tbody>
    {#each items as item}
        <tr>
            <td class="id">#{item.id}</td>
            <td>{item.added}</td>
            <td>{item.height} cm</td>
        </tr>
    {/each}
    </tbody>
</Table>

<style>
    td.id {
        width: 2em;
    }
</style>