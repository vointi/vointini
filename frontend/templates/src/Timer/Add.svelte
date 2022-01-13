<script lang="ts">
    import {Table} from 'sveltestrap'
    import {createEventDispatcher, onMount} from "svelte";
    import {createForm} from "felte";
    import {handlePageSubmit} from "../common";
    import {DTOTimerAdd} from "../dto";

    const dispatch = createEventDispatcher();

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    export let title: string = ""

    let s: number = 0
    let m: number = 20
    let h: number = 0

    $: seconds = (h * 60 * 60) + (m * 60) + s

    const success = (values: object) => {
        dispatch('submit')
        console.log(values)
    }

    const {form} = createForm({
        onSubmit: async (values, event) => {
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            delete values.s
            delete values.m
            delete values.h

            let vals: DTOTimerAdd = <DTOTimerAdd>{
                seconds: seconds,
                title: values.title,
            }

            console.log(vals)

            handlePageSubmit(endpointUrl, method, vals, success)
        },
    })

    // First page load
    onMount(async () => {
    })


</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/timer"
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
                       placeholder="Clean room">
            </td>
        </tr>

        <tr> <!-- Title -->
            <td class="tlabel">
                Time
            </td>
            <td>
                Total {seconds} Second(s)

                <label for="s">Seconds</label>
                <input class="form-control numval" type="number" min="0" id="s" name="s" bind:value={s}
                       placeholder="Seconds">

                <label for="m">Minutes</label>
                <input class="form-control numval" type="number" min="0" id="m" name="m" bind:value={m}
                       placeholder="Minutes">

                <label for="h">Hours</label>
                <input class="form-control numval" type="number" min="0" id="h" name="h" bind:value={h}
                       placeholder="Hours">

            </td>
        </tr>


        <tr> <!-- Submit -->
            <td class="tlabel">Save</td>
            <td><input type="submit" value="Add"></td>
        </tr>

        </tbody>

    </Table>

</form>

<style>
    input.numval {
        width: 4em;
    }
</style>