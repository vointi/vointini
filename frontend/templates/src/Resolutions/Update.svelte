<script lang="ts">
    import {createEventDispatcher, onMount} from "svelte";
    import {Table} from 'sveltestrap'
    import {createForm} from 'felte'
    import {handlePageSubmit} from "../common";
    import type {DTOResolutionEntity, DTOResolutionsUpdate} from "../dto";

    const dispatch = createEventDispatcher();

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    let entities: Array<DTOResolutionEntity> = []

    export let id: number = -1

    let name: string = ""
    let selectedEntity: number = -1
    let sentdate: string = ""
    let decisiondate: string = ""
    let startdate: string = ""
    let enddate: string = ""

    const success = (values: object) => {
        dispatch('submit')
        console.log(values)
    }

    const {form} = createForm({
        onSubmit: async (values, event) => {
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            const req: DTOResolutionsUpdate = <DTOResolutionsUpdate>{
                name: values.name,
                entityid: selectedEntity,
                decisiondate: values.decisiondate,
                sentdate: values.sentdate,
                startdate: values.startdate,
                enddate: values.enddate,
            }

            handlePageSubmit(endpointUrl, method, req, success)
        },
    })

    onMount(async () => {
        // Fetch list of entities
        //@See backend/restapi/router.go
        const fe = await fetch(
            "/api/v1/resolution-entities"
        )

        entities = await fe.json()

        selectedEntity = entities[0].id

        if (id === -1) {
            // Do not fetch old data
            return
        }

        // Fetch old data if we're updating
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/resolution/" + id
        )

        if (!res.ok) {
            return
        }

        const data: DTOResolutionsUpdate = await res.json()

        name = data.name
        selectedEntity = data.entityid
        sentdate = data.sentdate
        decisiondate = data.decisiondate
        startdate = data.startdate
        enddate = data.enddate
    })
</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/resolution/{id}"
>

    <Table class="table table-striped">

        <thead>
        </thead>

        <tbody>

        <tr> <!-- Entity -->
            <td class="tlabel">
                <label for="entityid">Entity name</label>
            </td>
            <td>
                <select name="entityid" id="entityid" bind:value={selectedEntity}>
                    {#each entities as entity}
                        <option value="{entity.id}">{entity.name}</option>
                    {/each}
                </select>
            </td>
        </tr>

        <tr> <!-- Name -->
            <td class="tlabel">
                <label for="name">Name</label>
            </td>
            <td>
                <input class="form-control" type="text" id="name" name="name" bind:value="{name}"
                       placeholder="Unemployment benefit">
            </td>
        </tr>

        <tr> <!-- Sent -->
            <td class="tlabel">
                <label for="sentdate">Sent date (optional)</label>
            </td>
            <td>
                <input class="form-control" type="date" id="sentdate" name="sentdate" bind:value="{sentdate}"
                       placeholder="2020-01-01">
            </td>
        </tr>

        <tr> <!-- Decision -->
            <td class="tlabel">
                <label for="decisiondate">Decision date (optional)</label>
            </td>
            <td>
                <input class="form-control" type="date" id="decisiondate" name="decisiondate"
                       bind:value="{decisiondate}"
                       placeholder="2020-01-15">
            </td>
        </tr>

        <tr> <!-- Start -->
            <td class="tlabel">
                <label for="startdate">Start date</label>
            </td>
            <td>
                <input class="form-control" type="date" id="startdate" name="startdate"
                       bind:value="{startdate}"
                       placeholder="2020-02-01">
            </td>
        </tr>

        <tr> <!-- End -->
            <td class="tlabel">
                <label for="enddate">End date</label>
            </td>
            <td>
                <input class="form-control" type="date" id="enddate" name="enddate" bind:value="{enddate}"
                       placeholder="2020-03-01">
            </td>
        </tr>

        <tr> <!-- Submit -->
            <td class="tlabel">Save</td>
            <td><input type="submit" value="Save"></td>
        </tr>

        </tbody>

    </Table>

</form>

<style>
    input {
        display: inline;
    }

    label {
        display: inline;
    }

    td.tlabel {
        min-width: 10em;
        text-align: right;
        padding-right: 1em;
    }

</style>
