<script lang="ts">
    import {Button, Modal, ModalHeader, Table} from "sveltestrap";
    import {createEventDispatcher, onMount} from "svelte";
    import {isNumber, Level} from "../common";
    import type {DTOResolutionEntity, DTOResolutions} from "../dto";
    import {default as Update} from "./Update.svelte"
    import UploadForm from "./UploadForm.svelte";
    import {DTOEntryLevel} from "../dto";

    let selectedId: number = -1
    let items: Array<DTOResolutions> = []
    let entities: Record<number, string> = {}

    let modalIsOpen: boolean = false;
    let modalStatus: string = 'Closed';

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    const toggle = (evt) => {
        if (isNumber(evt)) {
            selectedId = evt
        }

        modalIsOpen = !modalIsOpen
    }

    // First page load
    onMount(async () => {
        // Get resolution entity names
        //@See backend/restapi/router.go
        const entitylist = await fetch(
            "/api/v1/resolution-entities"
        )

        const data: Array<DTOResolutionEntity> = await entitylist.json()
        data.forEach((i) => {
            entities[i.id] = i.name
        })

        entities = entities

        // Get resolution files
        //@See backend/restapi/router.go
        const flist = await fetch(
            "/api/v1/resolution-entities"
        )

        const fdata: Array<DTOResolutionEntity> = await flist.json()
        fdata.forEach((i) => {
            entities[i.id] = i.name
        })

        entities = entities


        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/resolutions"
        )

        if (!res.ok) {
            return
        }

        items = await res.json()
    })

</script>

<!-- Modal for adding resolution form -->
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
    <ModalHeader {toggle}>Add resolution</ModalHeader>
    <Update on:submit={toggle} id={selectedId}/>
</Modal>

<Button on:click={toggle}>Add</Button>

<Table class="table table-striped">
    <thead>
    <tr>
        <th>Id</th>
        <th>Added</th>
        <th>Entity</th>
        <th>Name</th>
        <th>Files</th>
    </tr>
    </thead>

    <tbody>
    {#each items as item}
        <tr>
            <td class="id" on:click={() => toggle(item.id)}>#{item.id}</td>
            <td on:click={() => toggle(item.id)}>{item.added_at}</td>
            <td on:click={() => toggle(item.id)}>{entities[item.entityid]}</td>
            <td on:click={() => toggle(item.id)}>{item.name}</td>
            <!-- File upload -->
            <td class="uploadfile">
                <div class="uploadfile">
                    <UploadForm resolutionid={item.id}/>
                </div>
            </td>
        </tr>
    {/each}
    </tbody>
</Table>

<style>
    td.id {
        width: 2em;
    }

</style>