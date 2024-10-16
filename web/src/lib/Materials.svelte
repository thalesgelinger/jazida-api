<script>
    import { Button } from "./components/ui/button";

    const materials = ["Saibro", "Areia"];
</script>

<div
    class="w-[30vw] max-h-[50vh] overflow-auto rounded-md border flex flex-col items-center"
>
    <h1 class="w-full text-center p-4 text-lg">Materiais</h1>
    <div class="w-full">
        <ul class="flex flex-col w-full p-4 gap-2">
            {#each materials as material}
                <li class="rounded-md border p-4">{material}</li>
            {/each}
            <Button>Adicionar</Button>
        </ul>
    </div>
</div>
