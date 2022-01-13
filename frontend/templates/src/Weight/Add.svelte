<script lang="ts">
    import {Table} from "sveltestrap"
    import {createForm} from "felte";
    import {handlePageSubmit} from "../common";
    import {createEventDispatcher} from "svelte";
    import type {DTOWeightAdd} from "../dto";

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

    let weight:number = 0.0

    const {form} = createForm({
        onSubmit: async (values, event) => {
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            const wval: DTOWeightAdd = <DTOWeightAdd>{
                weight: values.weight,
            }

            handlePageSubmit(endpointUrl, method, wval, success)
        },
    })

</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/weight"
>

    <Table class="table-striped table table-dark">

        <thead>
        </thead>

        <tbody>

        <tr> <!-- Weight-->
            <td class="tlabel">
                Time
            </td>
            <td>
                <label for="weight">Weight (kg)</label>
                <input class="form-control numval" type="number" min="0" id="weight" name="weight" bind:value={weight}
                       placeholder="Weight">

            </td>
        </tr>

        <tr> <!-- Submit -->
            <td class="tlabel">Save</td>
            <td><input type="submit" value="Add"></td>
        </tr>

        </tbody>

    </Table>

</form>
