<script lang="ts">
    import {getDate, getTime} from "./common";

    export let hours:number = 8

    let now: Date = new Date()
    let later: Date = now

    function updateLater() {
        now = new Date()

        later = new Date(
            now.getFullYear(), now.getMonth(), now.getDate(),
            now.getHours() + hours, now.getMinutes(), now.getSeconds(), 0
        )
    }

    updateLater()

    const timer = setInterval(() => {
        updateLater()
    }, 1000 * 60)


</script>

<p>
    {hours} hours later is {getDate(later)} {getTime(later)}
</p>
