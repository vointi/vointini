<script lang="ts">
    import {onMount} from "svelte";
    import {Button, Modal, ModalHeader, Table} from 'sveltestrap'
    import {genHours, getDate, getLastDay, getNextDay, getTime, getWeekNumber, Level, onInterval} from "../common"
    import EntryUpdate from "./Update.svelte"
    import {DTOEntryLevel, DTOTag} from "../dto"


    const dayNames: Array<string> = ["su", "ma", "ti", "ke", "to", "pe", "la"]

    let levelNames: Record<string, Level> = {}
    let tagNames: Array<DTOTag> = []

    const getLevelNames = async () => {
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/entries/levels"
        )

        if (!res.ok) {
            return
        }

        const data: Array<DTOEntryLevel> = await res.json()
        data.forEach((i) => {
            levelNames[i.key] = <Level>{
                Name: i.name,
                Show: i.show,
                Worst: "",
            }
        })

        levelNames = levelNames// See: https://svelte.dev/tutorial/updating-arrays-and-objects
    }

    // List of entries
    let items: object = {}

    let now: Date = new Date()

    const today: Date = new Date(
        now.getFullYear(), now.getMonth(), now.getDate(),
        0, 0, 0, 0
    )

    let selectedDate: Date = new Date(
        now.getFullYear(), now.getMonth(), now.getDate(),
        0, 0, 0, 0
    )

    let nextDate: Date = getNextDay(selectedDate)
    let lastDate: Date = getLastDay(selectedDate)

    let hours: Array<Date> = genHours(selectedDate)

    async function getEntries(d: Date) {

        // Get entries for given day
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/entries/" + d.getFullYear() + '/' + (d.getMonth() + 1) + '/' + d.getDate()
        )

        if (!res.ok) {
            return
        }

        items = await res.json()
    }

    const getTags = async () => {
        const res = await fetch(
            "/api/v1/tags"
        )

        if (!res.ok) {
            return
        }

        tagNames = await res.json()
    }


    function updateSelectedDate(d: Date) {
        now = new Date()
        selectedDate = d
        nextDate = getNextDay(d)
        lastDate = getLastDay(d)
        hours = genHours(d)

        getEntries(d)
    }

    updateSelectedDate(selectedDate)

    // Time selected
    let selectedTime: Date = new Date(
        selectedDate.getFullYear(), selectedDate.getMonth(), selectedDate.getDate(),
        0, 0, 0, 0
    )

    let modalIsOpen = false;
    let status = 'Closed';

    const toggle = (d: Date) => {
        if (d instanceof Date) {
            selectedTime = d
        }

        modalIsOpen = !modalIsOpen
    }

    onMount(async () => {
        await getLevelNames()
        await getTags()
        await getEntries(selectedDate) // Load entries
    });

    function isNow(d: Date): boolean {
        return now.getFullYear() === d.getFullYear() &&
            now.getMonth() === d.getMonth() &&
            now.getDate() === d.getDate() &&
            now.getHours() === d.getHours()
    }

    onInterval(() => {
        now = new Date()
    }, 1000)

</script>

<!-- Modal for adding/updating entry form -->
<Modal
        body
        isOpen={modalIsOpen}
        size="xl"
        {toggle}
        on:opening={() => (status = 'Opening...')}
        on:open={() => (status = 'Opened')}
        on:closing={() => (status = 'Closing...')}
        on:close={() => (updateSelectedDate(selectedDate))}
>
    <ModalHeader {toggle}>Update entry {getDate(selectedTime)} {getTime(selectedTime)}</ModalHeader>
    <EntryUpdate d={selectedTime} on:submit={toggle}/>
</Modal>

<h2>🖊️ Entries</h2>

<a href="tags.html">Edit tags</a>

<div>
    <Button on:click={() => updateSelectedDate(lastDate)}>⬅️</Button>
    {getDate(selectedDate)} {dayNames[selectedDate.getDay()]} week {getWeekNumber(selectedDate)}
    <Button on:click={() => updateSelectedDate(nextDate)}>➡️</Button>
    <Button on:click={() => updateSelectedDate(today)}>Today</Button>
</div>

<Table class="table-striped">
    <thead>
    <tr> <!-- Columns -->

        <th class="time">Time</th>

        <th>Activity</th>
        <th>Description</th>
        <th>Tags</th>

        {#each Object.entries(levelNames) as [k, v]}
            <th class="rating">{v.Name}</th>
        {/each}

        <th>Achievement</th>

    </tr>
    </thead>

    <tbody> <!-- Rows -->

    {#each hours as item, idx}
        <tr class="{isNow(item) ? 'now' : ''}">

            <td class="time" on:click={() => {toggle(item)}}>{getTime(item)}</td>

            {#if items[getTime(item)] !== undefined}
                <td class="activity" on:click={() => {toggle(item)}}>{items[getTime(item)]["activity"]}</td>
                <td class="description" on:click={() => {toggle(item)}}>{items[getTime(item)]["description"]}</td>

                <!-- Tags -->
                <td class="tags">
                    <dl>
                    {#each tagNames as tag}
                        {#if items[getTime(item)]["tags"] !== null && items[getTime(item)]["tags"].includes(tag.shortname)}
                            <dd class="tag">{tag.name}</dd>
                        {/if}
                    {/each}
                    </dl>
                </td>

                <!-- Levels -->
                {#each Object.entries(levelNames) as [k, v]}
                    <td class="rating" on:click={() => {toggle(item)}}>
                        {#if items[getTime(item)]["levels"][k] !== undefined}
                            {items[getTime(item)]["levels"][k]}
                        {/if}
                    </td>
                {/each}

                <td on:click={() => {toggle(item)}}>{items[getTime(item)]["achievement"]}</td>

            {:else }
                <!-- empty -->
                <td on:click={() => {toggle(item)}}></td>
                <td></td> <!-- Description -->
                <td class="tags"></td> <!-- Tags -->

                <!-- Levels -->
                {#each Object.entries(levelNames) as [k, v]}
                    <td class="rating" on:click={() => {toggle(item)}}>-</td>
                {/each}

                <!-- Achievement -->
                <td on:click={() => {toggle(item)}}></td>
            {/if}

        </tr>
    {/each}

    </tbody>

</Table>

<style>
    td.rating {
        width: 2em;
    }

    td.time {
        font-family: monospace, monospace;
        width: 2em;
        font-weight: bold;
    }

    td.activity {
        min-width: 8em;
    }

    td.description {
        width: 8em;
        font-size: smaller;
        overflow: no-display;
    }

    dd.tag {
        font-size: small;
        display: inline;
        margin: 0;
    }

    dd.tag:not(:last-of-type)::after {
        content: ', ';
    }
</style>