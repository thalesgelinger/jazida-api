<script lang="ts">
    import * as Accordion from "$lib/components/ui/accordion/index.js";
    import Button from "./components/ui/button/button.svelte";

    const clients = [
        { name: "Client 1", plates: ["ABC-1234", "DEF-1234", "GHI-1234"] },
        { name: "Client 2", plates: ["ABC-1234", "DEF-1234", "GHI-1234"] },
        { name: "Client 3", plates: ["ABC-1234", "DEF-1234", "GHI-1234"] },
        { name: "Client 4", plates: ["ABC-1234", "DEF-1234", "GHI-1234"] },
        { name: "Client 5", plates: ["ABC-1234", "DEF-1234", "GHI-1234"] },
        { name: "Client 6", plates: ["ABC-1234", "DEF-1234", "GHI-1234"] },
    ];
</script>

<div
    class="w-[30vw] max-h-[50vh] rounded-md border flex flex-col items-center overflow-auto"
>
    <h1 class="w-full text-center p-4 text-lg">Clientes</h1>
    <Accordion.Root class="w-full sm:max-w-[70%]">
        {#each clients as client}
            <Accordion.Item value={client.name}>
                <Accordion.Trigger>{client.name}</Accordion.Trigger>
                <Accordion.Content>
                    <ul class="flex flex-col gap-2">
                        {#each client.plates as plate}
                            <li class="rounded-md border p-4">{plate}</li>
                        {/each}
                        <Button>Adicionar Placa</Button>
                    </ul>
                </Accordion.Content>
            </Accordion.Item>
        {/each}
    </Accordion.Root>
</div>
