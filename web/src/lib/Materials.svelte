<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "../service/api";
    import { Button } from "./components/ui/button";

    type Material = {
        id: number;
        name: string;
    };

    let materials: Material[] = [];

    const getClients = () => {
        api.get<Material[]>("/materials", {
            headers: {
                Authorization: "loader",
            },
        }).then((response) => (materials = response.data));
    };
    onMount(getClients);
</script>

<div
    class="w-[30vw] max-h-[50vh] overflow-auto rounded-md border flex flex-col items-center"
>
    <h1 class="w-full text-center p-4 text-lg">Materiais</h1>
    <div class="w-full">
        <ul class="flex flex-col w-full p-4 gap-2">
            {#each materials as material}
                <li class="rounded-md border p-4">{material.name}</li>
            {/each}
            <Button>Novo Material</Button>
        </ul>
    </div>
</div>
