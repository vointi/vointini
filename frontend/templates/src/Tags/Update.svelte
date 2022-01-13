<script lang="ts">
    import {Table} from 'sveltestrap'
    import {createForm} from 'felte'
    import {handlePageSubmit} from "../common";
    import {createEventDispatcher, onMount} from "svelte";
    import type {DTOTag} from "../dto";

    const dispatch = createEventDispatcher();

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    export let id: number = -1

    let name: string = ""
    let shortname: string = ""

    const success = (values: object) => {
        dispatch('submit')
        console.log(values)
    }

    const {form} = createForm({
        onSubmit: async (values, event) => {
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            handlePageSubmit(endpointUrl, method, values, success)
        },
    })

    onMount(async () => {
        // Fetch old data if we're updating
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/tag/" + id
        )

        if (!res.ok) {
            return
        }

        const data: DTOTag = await res.json()

        name = data.name
        shortname = data.shortname

    })
</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/tag/{id}"
>

    <Table class="table-striped table table-dark">

        <thead>
        </thead>

        <tbody>

        <tr> <!-- Name -->
            <td class="tlabel">
                <label for="name">Name</label>
            </td>
            <td>
                <input class="form-control" type="text" id="name" name="name" bind:value="{name}"
                       placeholder="Exercise">
            </td>
        </tr>

        <tr> <!-- Short name -->
            <td class="tlabel">
                <label for="shortname">Short name (machine)</label>
            </td>
            <td>
                <input class="form-control" type="text" id="shortname" name="shortname" bind:value="{shortname}"
                       placeholder="exercise">
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
