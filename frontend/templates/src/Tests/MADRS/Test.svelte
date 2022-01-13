<script lang="ts">
    import type {question} from "./i18n";
    import {Table} from "sveltestrap";
    import {createForm} from 'felte'
    import {createEventDispatcher} from "svelte";
    import type {DTOTestMADRSAnswers} from "../../dto";
    import {handlePageSubmit} from "../../common";

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

    // List of fields which failed (a1, a2, ..., a10) (aN)
    let errors: Array<string> = []

    const {form} = createForm({
        onSubmit: async (values, event) => {
            const method = event.form.method.toUpperCase()
            const endpointUrl = event.form.action

            errors = []

            let a: DTOTestMADRSAnswers = <DTOTestMADRSAnswers>{}

            for (let i = 1; i < 11; i++) {
                if (values["a" + i] === undefined) {
                    errors.push("a" + i)
                    continue
                }

                a["a" + i] = parseInt(values["a" + i], 10)
            }

            errors = errors

            if (errors.length !== 0) {
                // Errors in form
                return
            }

            handlePageSubmit(endpointUrl, method, a, success)
        },
    })


    const q: Array<question> = [

        // Question 1
        {
            Question: {
                Title: "Apparent sadness",
                Description: "Representing despondency, gloom and despair (more than just ordinary transient low spirits), " +
                    "reflected in speech, facial expression, and posture. Rate by depth and inability to brighten up."
            },
            Answer: [
                {
                    Description: "No sadness."
                },
                {
                    Description: "Looks dispirited but does brighten up without difficulty."
                },
                {
                    Description: "Appears sad and unhappy most of the time."
                },
                {
                    Description: "Looks miserable all the time. Extremely despondent"
                },
            ]
        },

        // Question 2
        {
            Question: {
                Title: "Reported sadness",
                Description: "Representing reports of depressed mood, regardless of whether it is reflected in appearance or not. Includes low spirits, despondency or the feeling of being beyond help and without hope."
            },
            Answer: [
                {
                    Description: "Occasional sadness in keeping with the circumstances."
                },
                {
                    Description: "Sad or low but brightens up without difficulty."
                },
                {
                    Description: "Pervasive feelings of sadness or gloominess. The mood is still influenced by external" +
                        "circumstances."
                },
                {
                    Description: "Continuous or unvarying sadness, misery or despondency."
                },
            ]
        },

        // Question 3
        {
            Question: {
                Title: "Inner tension",
                Description: "Representing feelings of ill-defined discomfort, edginess, inner turmoil, mental tension mounting to " +
                    "either panic, dread or anguish. Rate according to intensity, frequency, duration and the extent of " +
                    "reassurance called for.",
            },
            Answer: [
                {
                    Description: "Placid. Only fleeting inner tension."
                },
                {
                    Description: "Occasional feelings of edginess and ill-defined discomfort."
                },
                {
                    Description: "Continuous feelings of inner tension or intermittent panic which the patient can only master with some difficulty."
                },
                {
                    Description: "Unrelenting dread or anguish. Overwhelming panic."
                },
            ]
        },

        // Question 4
        {
            Question: {
                Title: "Reduced sleep",
                Description: "Representing the experience of reduced duration or depth of sleep compared to the subject's own normal pattern when well."
            },
            Answer: [
                {
                    Description: "Sleeps as normal."
                },
                {
                    Description: "Slight difficulty dropping off to sleep or slightly reduced, light or fitful sleep."
                },
                {
                    Description: "Moderate stiffness and resistance"
                },
                {
                    Description: "Sleep reduced or broken by at least 2 hours."
                },
            ]
        },

        // Question 5
        {
            Question: {
                Title: "Reduced appetite",
                Description: "Representing the feeling of a loss of appetite compared with when-well. Rate by loss of desire for food or the need to force oneself to eat."
            },
            Answer: [
                {
                    Description: "Normal or increased appetite."
                },
                {
                    Description: "Slightly reduced appetite."
                },
                {
                    Description: "No appetite. Food is tasteless."
                },
                {
                    Description: "Needs persuasion to eat at all."
                },
            ]
        },

        // Question 6
        {
            Question: {
                Title: "Concentration difficulties",
                Description: "Representing difficulties in collecting one's thoughts mounting to an incapacitating lack of " +
                    "concentration. Rate according to intensity, frequency, and degree of incapacity produced."
            },
            Answer: [
                {
                    Description: "No difficulties in concentrating."
                },
                {
                    Description: "Occasional difficulties in collecting one's thoughts."
                },
                {
                    Description: "Difficulties in concentrating and sustaining thought which reduced ability to read or hold a conversation."
                },
                {
                    Description: "Unable to read or converse without great difficulty."
                },
            ]
        },

        // Question 7
        {
            Question: {
                Title: "Lassitude",
                Description: "Representing difficulty in getting started or slowness in initiating and performing everyday activities."
            },
            Answer: [
                {
                    Description: "Hardly any difficulty in getting started. No sluggishness."
                },
                {
                    Description: "Difficulties in starting activities."
                },
                {
                    Description: "Difficulties in starting simple routine activities which are carried out with effort."
                },
                {
                    Description: "Complete lassitude. Unable to do anything without help."
                },
            ]
        },

        // Question 8
        {
            Question: {
                Title: "Inability to feel",
                Description: "Representing the subjective experience of reduced interest in the surroundings, or activities that normally give pleasure. The ability to react with adequate emotion to circumstances or people is reduced."
            },
            Answer: [
                {
                    Description: "Normal interest in the surroundings and in other people."
                },
                {
                    Description: "Reduced ability to enjoy usual interests."
                },
                {
                    Description: "Loss of interest in the surroundings. Loss of feelings for friends and acquaintances."
                },
                {
                    Description: "The experience of being emotionally paralysed, inability to feel anger, grief or pleasure and a complete or even painful failure to feel for close relatives and friends."
                },
            ]
        },

        // Question 9
        {
            Question: {
                Title: "Pessimistic thoughts",
                Description: "Representing thoughts of guilt, inferiority, self-reproach, sinfulness, remorse and ruin."
            },
            Answer: [
                {
                    Description: "No pessimistic thoughts."
                },
                {
                    Description: "Fluctuating ideas of failure, self-reproach or self-depreciation."
                },
                {
                    Description: "Persistent self-accusations, or definite but still rational ideas of guilt or sin. Increasingly pessimistic about the future."
                },
                {
                    Description: "Delusions of ruin, remorse or irredeemable sin. Self- accusations which are absurd and unshakable."
                },
            ]
        },

        // Question 10
        {
            Question: {
                Title: "Suicidal thoughts",
                Description: "Representing the feeling that life is not worth living, that a natural death would be welcome, suicidal thoughts, and preparations for suicide. Suicide attempts should not in themselves influence the rating."
            },
            Answer: [
                {
                    Description: "Enjoys life or takes it as it comes."
                },
                {
                    Description: "Weary of life. Only fleeting suicidal thoughts."
                },
                {
                    Description: "Probably better off dead. Suicidal thoughts are common, and suicide is considered as a possible solution, but without specific plans or intenstion."
                },
                {
                    Description: "Explicit plans for suicide when there is an opportunity. Active preparations for suicide."
                },
            ]
        },

    ]

</script>

<h2>Test: MADRS</h2>

<form
        use:form
        on:keydown={handleEnter}
        method="post"
        action="/api/v1/tests/madrs"
>

    {#each q as question, i}
        <h4>{question.Question.Title}</h4>
        <p>{question.Question.Description}</p>

        <Table class="table-striped table">
            <thead>

            <tr>
                <th class="qselect">

                </th>
                <th class="qanswer">
                    Answer
                </th>
            </tr>

            </thead>

            <tbody>

            {#each question.Answer as answer, ai }
                <tr>
                    <td class="qselect" class:required={errors.indexOf("a" + (i+1)) !== -1}>
                        <input id="a{i+1}q{ai}" name="a{i+1}" type="radio" value="{ai}"/>
                    </td>
                    <td class="qanswer">
                        <label for="a{i+1}q{ai}"
                               class:required={errors.indexOf("a" + (i+1)) !== -1}>{answer.Description}</label>
                    </td>
                </tr>
            {/each}

            </tbody>
        </Table>
    {/each}

    <input type="submit" value="Save">

</form>

<style>
    td.qselect {
        width: 2em;
    }

    label.required {
        color: #8f3937;
    }

</style>