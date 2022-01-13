<script lang="ts">
    import {default as Tasks} from "./Tasks/List.svelte"
    import {default as ReoccurringTaskList} from "./Tasks/ReoccurringTaskList.svelte"
    import Header from "./Header.svelte";
</script>

<main>
    <Header />

    <Tasks/>
    <ReoccurringTaskList/>
</main>

<style>
</style>