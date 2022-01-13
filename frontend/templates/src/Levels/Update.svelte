<script lang="ts">
    import {Table} from "sveltestrap"
    import {createForm} from "felte";
    import {handlePageSubmit} from "../common";
    import {createEventDispatcher} from "svelte";
    import type {DTOEntryLevelAdd} from "../dto";

    const dispatch = createEventDispatcher();

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    function success(values: object) {
        dispatch('submit')
        console.log(values)
    }

    export let id: number = -1
    let name: string = ""
    let key: string = ""
    let worst: string = ""

    const {form} = createForm({
        onSubmit: async (values, event) => {
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            const wval: DTOEntryLevelAdd = <DTOEntryLevelAdd>{
                name: "",
                key: "",
                show: true,
                worst: "",
            }

            handlePageSubmit(endpointUrl, method, wval, success)
        },
    })

</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/height"
>

    <Table class="table-striped table table-dark">

        <thead>
        </thead>

        <tbody>

        <tr> <!-- Level-->
            <td class="tlabel">
                Level
            </td>
            <td>
                <label for="name">Name</label>
                <input class="form-control" type="text" id="name" name="name" bind:value={name}
                       placeholder="Name">
                <label for="key">Key</label>
                <input class="form-control" type="text" id="key" name="key" bind:value={key}
                       placeholder="Name">
                <label for="worst">Description for worst situation (10/10)</label>
                <textarea id="worst" name="worst"></textarea>
            </td>
        </tr>

        <tr> <!-- Submit -->
            <td class="tlabel">Save</td>
            <td><input type="submit" value="Add"></td>
        </tr>

        </tbody>

    </Table>

</form>
