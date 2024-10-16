<script lang="ts">
    import * as Table from "$lib/components/ui/table";
    import { onMount } from "svelte";
    import { api } from "../service/api";

    type Load = {
        [x: string]: string;
        id: string;
        client: string;
        plate: string;
        material: string;
        quantity: string;
        paymentMethod: string;
        signature: string;
    };

    const getLoads = async (): Promise<Load[]> => {
        const response = await api.get("/loads", {
            headers: {
                Authorization: "admin",
            },
        });
        return response.data.map((d: { paymentmethod: any }) => ({
            ...d,
            paymentMethod: d.paymentmethod,
        })) as Load[];
    };

    let loads: Load[] = [];

    onMount(() => {
        getLoads().then((data) => {
            loads = data.reverse();
        });
    });

    const ws = new WebSocket("ws://localhost:8080/new-load-added");
    ws.onmessage = ({ data }) => {
        const newLoad: Load = JSON.parse(data);
        newLoad.paymentMethod = newLoad.paymentmethod;
        newLoad.id = `${loads.length + 1}`;
        loads = [newLoad, ...loads];
    };
</script>

<div class="rounded-md border">
    <h1 class="w-full text-center p-4 text-lg">Carregamentos</h1>
    <Table.Root>
        <Table.Header>
            <Table.Row>
                <Table.Head>Id</Table.Head>
                <Table.Head>Cliente</Table.Head>
                <Table.Head>Placa</Table.Head>
                <Table.Head>Material</Table.Head>
                <Table.Head>Quantidade</Table.Head>
                <Table.Head>Método de pagamento</Table.Head>
                <Table.Head>Assinatura</Table.Head>
            </Table.Row>
        </Table.Header>
        <Table.Body>
            {#each loads as load}
                <Table.Row>
                    <Table.Cell>{load.id}</Table.Cell>
                    <Table.Cell>{load.client}</Table.Cell>
                    <Table.Cell>{load.plate}</Table.Cell>
                    <Table.Cell>{load.material}</Table.Cell>
                    <Table.Cell>{load.quantity}</Table.Cell>
                    <Table.Cell>{load.paymentMethod}</Table.Cell>
                    <Table.Cell>
                        <img
                            src={load.signature}
                            alt={`signature ${load.id}`}
                            class="h-10"
                        />
                    </Table.Cell>
                </Table.Row>
            {/each}
        </Table.Body>
    </Table.Root>
</div>
