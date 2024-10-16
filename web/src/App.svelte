<script lang="ts">
    import Clients from "$lib/Clients.svelte";
    import Materials from "$lib/Materials.svelte";
    import Loads from "./lib/Loads.svelte";
</script>

<div class="flex flex-row w-full justify-center gap-4 p-4">
    <div class="flex flex-col gap-4">
        <Clients />
        <Materials />
    </div>
    <div class="max-w-[60vw] w-full">
        <Loads />
    </div>
</div>
