<script lang="ts">
    import * as Accordion from "$lib/components/ui/accordion/index.js";
    import { onMount } from "svelte";
    import { api } from "../service/api";
    import Button from "./components/ui/button/button.svelte";
    import * as AlertDialog from "./components/ui/alert-dialog";
    import { Input } from "./components/ui/input";

    type Client = {
        id: number;
        name: string;
        plates: Array<{
            id: number;
            plate: string;
        }>;
    };

    let clients: Client[] = [];

    const getClients = async () => {
        const response = await api.get<Client[]>("/clients", {
            headers: {
                Authorization: "loader",
            },
        });

        clients = response.data;
    };

    onMount(getClients);

    let newClient = "";
    let newPlate = "";

    const addClient = async () => {
        await api.post(
            "/clients",
            {
                name: newClient,
            },
            {
                headers: {
                    Authorization: "admin",
                },
            },
        );

        clients = [
            ...clients,
            {
                id: clients.length + 1,
                name: newClient,
                plates: [],
            },
        ];

        newClient = "";
    };

    const addPlate = (clientId: number) => async () => {
        await api.post(
            `/clients/${clientId}/plates`,
            {
                plate: newPlate,
            },
            {
                headers: {
                    Authorization: "admin",
                },
            },
        );

        const clientsCopy = [...clients];
        const clientIdx = clientsCopy.findIndex((c) => c.id === clientId);
        const client = clientsCopy[clientIdx];
        client?.plates.push({
            id: client.plates.length + 1,
            plate: newPlate,
        });

        clientsCopy[clientIdx] = client;
        clients = [...clientsCopy];
    };
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
                            <AlertDialog.Root>
                                <AlertDialog.Trigger>
                                    <Button>Nova Placa</Button>
                                </AlertDialog.Trigger>
                                <AlertDialog.Content>
                                    <AlertDialog.Header>
                                        <AlertDialog.Title
                                            >Nova Placa</AlertDialog.Title
                                        >
                                        <Input
                                            type="text"
                                            placeholder="Digite o nome do cliente"
                                            bind:value={newPlate}
                                        />
                                    </AlertDialog.Header>
                                    <AlertDialog.Footer>
                                        <AlertDialog.Cancel
                                            >Cancel</AlertDialog.Cancel
                                        >
                                        <AlertDialog.Action
                                            on:click={addPlate(client.id)}
                                            >Salvar</AlertDialog.Action
                                        >
                                    </AlertDialog.Footer>
                                </AlertDialog.Content>
                            </AlertDialog.Root>
                        </ul>
                    </Accordion.Content>
                </Accordion.Item>
            {/each}
        </Accordion.Root>

        <AlertDialog.Root>
            <AlertDialog.Trigger>
                <Button class="mt-2">Novo Cliente</Button>
            </AlertDialog.Trigger>
            <AlertDialog.Content>
                <AlertDialog.Header>
                    <AlertDialog.Title>Novo Cliente</AlertDialog.Title>
                    <Input
                        type="text"
                        placeholder="Digite o nome do cliente"
                        bind:value={newClient}
                    />
                </AlertDialog.Header>
                <AlertDialog.Footer>
                    <AlertDialog.Cancel>Cancel</AlertDialog.Cancel>
                    <AlertDialog.Action on:click={addClient}
                        >Salvar</AlertDialog.Action
                    >
                </AlertDialog.Footer>
            </AlertDialog.Content>
        </AlertDialog.Root>
    </div>
</div>
