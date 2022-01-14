<script lang="ts">
    import {createEventDispatcher, onMount} from "svelte";
    import {Table} from 'sveltestrap'
    import {createForm} from 'felte'
    import {getTime, handlePageSubmit, Level} from "../common";
    import {DTOEntry, DTOEntryAdd, DTOEntryLevel, DTOTag} from "../dto";

    const dispatch = createEventDispatcher();

    export let d: Date = new Date()

    function handleEnter(event) {
        if (event.key === "Enter") {
            if (event.target.tagName === "INPUT" && event.target.type !== "submit") {
                event.preventDefault();
            }
        }
    }

    let levelNames: Record<string, Level> = {}
    let tagNames: Array<DTOTag> = []
    let selectedTags: Array<string> = []
    let lastValues: DTOEntry = <DTOEntry>{}

    // Prepopulated
    let levels: Record<string, number> = {}

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
                Worst: i.worst,
            }

            levels[i.key] = 0 // Populate default to zero
        })

        levelNames = levelNames// See: https://svelte.dev/tutorial/updating-arrays-and-objects
        levels = levels// See: https://svelte.dev/tutorial/updating-arrays-and-objects
    }

    const getTags = async () => {
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/tags"
        )

        if (!res.ok) {
            return
        }

        tagNames = await res.json()
    }

    let activity: string = ""
    let description: string = ""
    let achievement: number = 0

    function success(values: object) {
        dispatch('submit')
        console.log(values)
    }

    const {form} = createForm({
        onSubmit: async (values, event) => {
            delete values.ranges

            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            let req: DTOEntryAdd = <DTOEntryAdd>{
                activity: values.activity,
                description: values.description,
                achievement: values.achievement,
                levels: {}, // levels
                tags: selectedTags, // selected tags (shortnames)
            }

            Object.keys(values.levels).forEach(key => {
                if (values.levels[key] === undefined) {
                    return
                }

                req.levels[key] = values.levels[key]
            })

            handlePageSubmit(endpointUrl, method, req, success)
        },
    })

    const updateActivity = (evt) => {
        // Use tags as activity name
        evt.preventDefault()

        if (tagNames === null) {
            return
        }

        activity = ""

        tagNames.forEach((val) => {
            if (selectedTags.length === 0) {
                return
            }

            selectedTags.forEach((tval) => {
                if (val.shortname == tval) {
                    activity += val.name + ', '
                }
            })
        })

        activity = activity.trim()
        activity = activity.replace(/,$/g, '')
    }

    const fetchOld = async (old: Date) => {
        // Fetch old data if we're updating
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/entries/" + old.getFullYear() + '/' + (old.getMonth() + 1) + '/' + old.getDate() + '/' + old.getHours() + '/' + old.getMinutes()
        )

        if (!res.ok) {
            return
        }

        const data = await res.json()

        if (getTime(old) in data) {
            lastValues = data[getTime(old)]
        }
    }

    onMount(async () => {
        await getLevelNames()
        await getTags()

        // Fetch old data if we're updating
        //@See backend/restapi/router.go
        const res = await fetch(
            "/api/v1/entries/" + d.getFullYear() + '/' + (d.getMonth() + 1) + '/' + d.getDate() + '/' + d.getHours() + '/' + d.getMinutes()
        )

        if (!res.ok) {
            return
        }

        const data: object = await res.json()

        if (getTime(d) in data) {
            // has hh:mm key
            const item: DTOEntry = data[getTime(d)]
            activity = item.activity
            description = item.description
            achievement = item.achievement
            levels = item.levels
            if (item.tags !== null) {
                selectedTags = item.tags
            }
        } else {
            // Fetch previous data to auto-refill some values

            // Used in loop
            let lastUpdate: Date = new Date(
                d.getFullYear(), d.getMonth(), d.getDate(),
                d.getHours() + 1, 0, 0, 0
            )

            // Try getting entries from future
            let tries: number = 0
            do {
                await fetchOld(lastUpdate)

                if (Object.keys(lastValues).length !== 0) {
                    break
                }

                lastUpdate.setHours(lastUpdate.getHours() + 1, lastUpdate.getMinutes(), 0)

                tries++
            } while (tries < 24)

            // Used in loop
            lastUpdate = new Date(
                d.getFullYear(), d.getMonth(), d.getDate(),
                d.getHours() - 1, 0, 0, 0
            )

            // Try getting entries from past
            tries = 0
            do {
                if (Object.keys(lastValues).length !== 0) {
                    // We have old data already!
                    break
                }

                await fetchOld(lastUpdate)

                if (Object.keys(lastValues).length !== 0) {
                    break
                }

                lastUpdate.setHours(lastUpdate.getHours() - 1, lastUpdate.getMinutes(), 0)

                tries++
            } while (tries < 24)


            if (lastValues === undefined) {
                // No old values
                return
            }

            if (lastValues === null) {
                // No old values
                return
            }

            if (Object.keys(lastValues).length === 0) {
                // No old values
                return
            }

            Object.keys(lastValues.levels).forEach(key => {
                levels[key] = lastValues.levels[key]
            })

            selectedTags = lastValues.tags // See: https://svelte.dev/tutorial/updating-arrays-and-objects
        }

        levels = levels // See: https://svelte.dev/tutorial/updating-arrays-and-objects
    })
</script>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/entries/{d.getFullYear()}/{d.getMonth()+1}/{d.getDate()}/{d.getHours()}/{d.getMinutes()}">

    <Table class="table-striped table table-dark">

        <thead>
        </thead>

        <tbody>

        <tr> <!-- Tag(s) (0-N can be selected) -->
            <td class="tlabel">🏷️ Tags</td>
            <td>
                {#each tagNames as tag}
                    <input class="tag" type="checkbox" id="tag{tag.id}" name="tag{tag.id}"
                           bind:group={selectedTags} value="{tag.shortname}">
                    <label for="tag{tag.id}" class="tag">{tag.name}</label>
                {/each}
            </td>
        </tr>

        <tr> <!-- Activity -->
            <td class="tlabel">
                <label for="activity">⚙️Activity</label>
            </td>
            <td>
                <input class="form-control" type="text" id="activity" name="activity" bind:value="{activity}"
                       placeholder="Woke up">
                <button on:click={updateActivity}>Update from tags</button>
            </td>
        </tr>

        <tr> <!-- Description -->
            <td class="tlabel">
                <label for="activity">📝 Description (optional)</label>
            </td>
            <td>
                    <textarea class="form-control" id="description" name="description"
                              placeholder="I had good night sleep">{description}</textarea>
            </td>
        </tr>

        {#each Object.entries(levelNames) as [k, v]}
            <tr>
                <td class="tlabel">
                    <label for="levels.{k}">{v.Name} (0-10, 10 being worst)</label>
                </td>
                <td>
                    <input class="form-control rating" type="number" min="0" max="10" id="levels.{k}"
                           name="levels.{k}"
                           bind:value="{levels[k]}">
                    0 <input class="form-range range" name="ranges.{k}" type="range" min="0" max="10"
                             bind:value="{levels[k]}" tabindex="-1"> 10 {v.Worst}
                </td>
            </tr>
        {/each}

        <tr> <!-- Achievement -->
            <td class="tlabel">
                <label for="achievement">🏆 Achievement (-10 - 10, 10 being best)</label>
            </td>
            <td>
                <input class="form-control rating" type="number" min="-10" max="10" id="achievement"
                       name="achievement" bind:value="{achievement}"
                >
                -10 <input class="form-range range" name="achievementr" type="range" min="-10" max="10"
                           bind:value="{achievement}" tabindex="-1"
            > 10
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

    input.rating {
        width: 5em;
    }

    input.range {
        width: 20em;
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

    label.tag {
        min-width: 4em;
        display: inline-block;
        border: 2px solid #111;
        background-color: #62686d;
        padding: 0.5em;
        border-radius: 4px;
        text-align: center;
        margin-bottom: 0.2em;
    }

    input.tag {
        padding: 0 0 0 0;
        margin: 0 0 0 0;
        border: none;
        appearance: none;
    }

    input.tag:checked + label {
        background-color: #2f5e2f;
    }

</style>
