<script lang="ts">
    import {Button, Modal, ModalHeader, Table} from "sveltestrap"
    import {onDestroy, onMount} from "svelte";
    import {DTOTimer} from "../dto";
    import {default as TimerAdd} from "./Add.svelte";

    let modalIsOpen: boolean = false;
    let modalStatus:string = 'Closed';

    const toggle = (evt) => {
        modalIsOpen = !modalIsOpen
    }

    let timers: Array<DTOTimer> = []

    function addToTimers(data: DTOTimer) {

        /*
        if (!(data instanceof DTOTimer)) {
            return
        }*/

        let idx: number = -1
        timers.forEach((item, i) => {
            if (idx >= 0) {
                return
            }

            if (item.id === data.id) {
                idx = i
                return
            }
        })

        if (idx === -1) {
            idx = timers.push(data)
            idx--
        }

        timers[idx] = data

        timers = timers // See: https://svelte.dev/tutorial/updating-arrays-and-objects
    }

    let listener:EventSource

    async function addTimer() {
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/timer",
        )

        if (!res.ok) {
            return
        }

        //addToTimers(new DTOTimer(await res.json()))
    }

    async function stopTimer(id: number) {
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/timer/" + id + "/stop",
        )

        if (!res.ok) {
            return
        }
    }

    async function removeTimer(id: number) {
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/timer/" + id + "/remove",
        )

        if (!res.ok) {
            return
        }
    }

    const connectEventSource = () => {
        //@See backend/restapi/router.go
        listener = new EventSource("/api/v1/timers/events")
        listener.onerror = (evt) => {
            console.log(evt)
        }
        listener.onmessage = (evt) => {
            const data: DTOTimer = JSON.parse(evt.data)
            addToTimers(data)
        }

    }

    // First page load
    onMount(async () => {

        await connectEventSource()

        return

        // Fetch old data if we're updating
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/timers"
        )

        if (!res.ok) {
            return
        }

        const data = await res.json()
        console.log(data)
    })

    onDestroy(() => {
        if(listener.readyState && listener.readyState === 1) {
            listener.close();
        }
    })


</script>

<!-- Modal for adding timer form -->
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
        <ModalHeader {toggle}>Add timer</ModalHeader>
        <TimerAdd on:submit={toggle}/>
</Modal>

<Button on:click={toggle}>Add timer</Button>

<Table class="table-striped">
    <thead>
    <tr>
        <th>Id</th>
        <th>Remaining</th>
        <th>Seconds</th>
        <th>Actions</th>
    </tr>
    </thead>

    <tbody>
    {#each timers as t}
        <tr>
            <td class="id">#{t.id}</td>
            <td>{t.f}</td>
            <td>{t.s} s</td>
            <td>
                <Button on:click={() => stopTimer(t.id)}>Stop</Button>
                <Button on:click={() => removeTimer(t.id)}>Remove</Button>
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