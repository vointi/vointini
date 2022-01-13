<script lang="ts">
    import {Table} from 'sveltestrap'
    import {createForm} from 'felte'
    import {handlePageSubmit} from "../common";
    import {onMount, createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    export let id: number = -1

    let title: string = ""
    let description: string = ""

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
            "/api/v1/task/" + id
        )

        if (!res.ok) {
            return
        }

        const data = await res.json()

        title = data.title

    })


</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/task/{id}"
>

    <Table class="table-striped table table-dark">

        <thead>
        </thead>

        <tbody>

        <tr> <!-- Title -->
            <td class="tlabel">
                <label for="title">Title</label>
            </td>
            <td>
                <input class="form-control" type="text" id="title" name="title" bind:value="{title}"
                       placeholder="Woke up">
            </td>
        </tr>

        <tr> <!-- Description -->
            <td class="tlabel">
                <label for="description">📝 Description (optional)</label>
            </td>
            <td>
                    <textarea class="form-control" id="description" name="description"
                              placeholder="...">{description}</textarea>
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

    textarea {
        width: 100%;
        height: 15em;
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
