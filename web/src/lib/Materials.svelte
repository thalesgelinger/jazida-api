<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "../service/api";
    import { Button } from "./components/ui/button";
    import * as AlertDialog from "./components/ui/alert-dialog";
    import { Input } from "./components/ui/input";

    type Material = {
        id: number;
        name: string;
    };

    let materials: Material[] = [];

    const getMaterials = async () => {
        const response = await api.get<Material[]>("/materials", {
            headers: {
                Authorization: "loader",
            },
        });
        materials = response.data;
    };
    onMount(getMaterials);

    let newMaterial = "";

    const addMaterial = async () => {
        await api.post(
            "/materials",
            { material: newMaterial },
            {
                headers: {
                    Authorization: "admin",
                },
            },
        );

        materials = [
            ...materials,
            {
                id: materials.length + 1,
                name: newMaterial,
            },
        ];

        newMaterial = "";
    };
</script>

<div
    class="w-[30vw] max-h-[50vh] overflow-auto rounded-md border flex flex-col items-center"
>
    <AlertDialog.Root>
        <h1 class="w-full text-center p-4 text-lg">Materiais</h1>
        <div class="w-full">
            <ul class="flex flex-col w-full p-4 gap-2">
                {#each materials as material}
                    <li class="rounded-md border p-4">{material.name}</li>
                {/each}
                <AlertDialog.Trigger>
                    <Button>Novo Material</Button>
                </AlertDialog.Trigger>
            </ul>
        </div>
        <AlertDialog.Content>
            <AlertDialog.Header>
                <AlertDialog.Title>Novo Material</AlertDialog.Title>
                <Input
                    type="text"
                    placeholder="Digite o nome do material"
                    bind:value={newMaterial}
                />
            </AlertDialog.Header>
            <AlertDialog.Footer>
                <AlertDialog.Cancel>Cancel</AlertDialog.Cancel>
                <AlertDialog.Action on:click={addMaterial}
                    >Salvar</AlertDialog.Action
                >
            </AlertDialog.Footer>
        </AlertDialog.Content>
    </AlertDialog.Root>
</div>
