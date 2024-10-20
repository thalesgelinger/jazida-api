<script lang="ts">
    import * as Accordion from "$lib/components/ui/accordion/index.js";
    import { onMount } from "svelte";
    import { api } from "../service/api";
    import Button from "./components/ui/button/button.svelte";

    type Client = {
        id: number;
        name: string;
        plates: Array<{
            id: number;
            plate: string;
        }>;
    };

    let clients: Client[] = [];

    const getClients = () => {
        api.get<Client[]>("/clients", {
            headers: {
                Authorization: "loader",
            },
        }).then((response) => (clients = response.data));
    };
    onMount(getClients);
</script>

<div
    class="w-[30vw] max-h-[50vh] rounded-md border flex flex-col items-center overflow-auto"
>
    <h1 class="w-full text-center p-4 text-lg">Clientes</h1>

    <div class="w-full flex flex-col items-center m-4">
        <Accordion.Root class="w-full sm:max-w-[70%]">
            {#each clients as client}
                <Accordion.Item value={client.name}>
                    <Accordion.Trigger>{client.name}</Accordion.Trigger>
                    <Accordion.Content>
                        <ul class="flex flex-col gap-2">
                            {#each client.plates as plate}
                                <li class="rounded-md border p-4">
                                    {plate.plate}
                                </li>
                            {/each}
                            <Button>Nova Placa</Button>
                        </ul>
                    </Accordion.Content>
                </Accordion.Item>
            {/each}
        </Accordion.Root>
        <Button class="mt-2">Novo Cliente</Button>
    </div>
</div>
