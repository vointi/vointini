<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import {createForm} from "felte";
    import {handlePageSubmit} from "../common";

    const dispatch = createEventDispatcher();

    export let resolutionid: number = -1

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    const success = (values: object) => {
        dispatch('submit')
        console.log(values)
    }

    const {form} = createForm({
        onSubmit: async (values, event) => {
            //event.preventDefault()
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            let fd = new FormData()

            Object.keys(values).forEach((k) => {
                fd.append(k, values[k][0])
            })

            handlePageSubmit(endpointUrl, method, fd, success)
        },
    })
</script>

<form
        use:form
        method="post"
        action="/api/v1/resolution-file/{resolutionid}"
        enctype="multipart/form-data"
>

    <!-- File chooser -->
    <label for="file">Upload file</label>
    <input type="file" name="file" id="file" multiple>

    <!-- Submit button -->
    <label for="upload"></label>
    <input type="submit" id="upload" value="Upload">
</form>
