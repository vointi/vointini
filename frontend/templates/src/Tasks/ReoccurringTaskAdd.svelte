<script lang="ts">
    import {Table} from 'sveltestrap'
    import {createForm} from 'felte'
    import {handlePageSubmit} from "../common";
    import {createEventDispatcher} from "svelte";
    import {DTOReoccurringTaskAdd} from "../dto";

    const dispatch = createEventDispatcher();

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    export let id:number = -1
    let title: string = ""

    let m: number = 0 // Minutes
    let h: number = 0 // Hours
    let d: number = 0 // Days
    let w: number = 0 // Weeks

    $: seconds = (w * 7 * 24 * 60 * 60) + (d * 24 * 60 * 60) + (h * 60 * 60) + (m * 60)

    function success(values: object) {
        dispatch('submit')
        console.log(values)
    }

    const {form} = createForm({
        onSubmit: async (values, event) => {
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            let vals: DTOReoccurringTaskAdd = <DTOReoccurringTaskAdd>{
                s: seconds,
                title: values.title,
            }

            handlePageSubmit(endpointUrl, method, vals, success)
        },
    })

</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/reoccurring-task"
>

    <input type="hidden" name="seconds" bind:value="{seconds}">

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
                       placeholder="Take out trash">
            </td>
        </tr>

        <tr> <!-- Title -->
            <td class="tlabel">
                Time
            </td>
            <td>

                <label for="m">Minutes</label>
                <input class="form-control numval" type="number" min="0" id="m" name="m" bind:value={m}
                       placeholder="Minutes">

                <label for="h">Hours</label>
                <input class="form-control numval" type="number" min="0" id="h" name="h" bind:value={h}
                       placeholder="Hours">

                <label for="d">Days</label>
                <input class="form-control numval" type="number" min="0" id="d" name="d" bind:value={d}
                       placeholder="Days">

                <label for="w">Weeks</label>
                <input class="form-control numval" type="number" min="0" id="w" name="w" bind:value={w}
                       placeholder="Weeks">

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
