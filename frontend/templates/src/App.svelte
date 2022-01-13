<script lang="ts">
    import {default as TimerList} from "./Timer/List.svelte"
    import {default as WeightList} from "./Weight/List.svelte"
    import {default as HeightList} from "./Height/List.svelte"
    import {default as LevelList} from "./Levels/List.svelte"
    import Header from "./Header.svelte";
</script>

<main>
    <Header />

    <TimerList/>
    <WeightList/>
    <HeightList/>
    <LevelList/>
</main>

<style>
</style>